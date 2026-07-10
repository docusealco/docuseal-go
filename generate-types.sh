#!/bin/sh
# Regenerates the SDK from the DocuSeal OpenAPI spec using Fern
# (runs the generator locally in Docker).
# Usage: ./generate-types.sh [path-or-url-to-openapi-json]
set -e

cd "$(dirname "$0")"

# The generator image is built from the docusealco/fern fork, which adds the
# `optionalsAsValues` config option (optional string properties as plain Go
# values instead of pointers). The image is local-only and carries a marker
# label: fern requires a plain X.Y.Z version, so an upstream 1.47.0 pulled
# from Docker Hub would otherwise shadow the fork build. Build it on first
# run; requires pnpm in addition to Node.js and Docker.
GENERATOR_IMAGE="fernapi/fern-go-sdk:1.47.0"

if [ "$(docker image inspect -f '{{index .Config.Labels "com.docuseal.fern-fork"}}' "$GENERATOR_IMAGE" 2>/dev/null)" != "true" ]; then
  echo "Building $GENERATOR_IMAGE from the docusealco/fern fork..."
  FERN_FORK_DIR="${FERN_FORK_DIR:-.fern-fork}"
  if [ ! -d "$FERN_FORK_DIR" ]; then
    git clone https://github.com/docusealco/fern "$FERN_FORK_DIR"
  fi
  (
    cd "$FERN_FORK_DIR"
    pnpm install
    pnpm exec turbo run dist:cli --filter=@fern-api/go-sdk
    docker build -f generators/go/sdk/Dockerfile \
      --label com.docuseal.fern-fork=true -t "$GENERATOR_IMAGE" .
  )
fi

SPEC="${1:-https://console.docuseal.com/openapi.yml?format=json}"

case "$SPEC" in
  http*) curl -sf "$SPEC" -o openapi.tmp.json ;;
  *) cp "$SPEC" openapi.tmp.json ;;
esac

# Drop webhook payload schemas (the SDK exposes only the REST client), swap
# the legacy POST /submissions for the newer /submissions/init (same request
# body, envelope response), strip tags so the client surface is flat and
# name method-input wrappers <OperationId>Params.
ruby -rjson -e '
  path = "openapi.tmp.json"
  spec = JSON.parse(File.read(path))
  spec.delete("webhooks")
  spec.delete("tags")

  init = spec["paths"]["/submissions"].delete("post")
  init["responses"]["200"]["content"]["application/json"].delete("example")
  init["responses"]["200"]["content"]["application/json"]["schema"] = {
    "type" => "object",
    "required" => %w[id submitters expired_at created_at],
    "properties" => {
      "id" => { "type" => "integer", "description" => "Submission unique ID number." },
      "submitters" => { "$ref" => "#/components/schemas/CreateSubmissionsFromEmailsResponse" },
      "expired_at" => { "type" => %w[string null], "description" => "The date and time when the submission expires." },
      "created_at" => { "type" => "string", "description" => "The date and time when the submission was created." }
    }
  }
  spec["paths"]["/submissions/init"] = { "post" => init }

  spec["paths"].each_value do |methods|
    methods.each_value do |op|
      next unless op.is_a?(Hash)

      op.delete("tags")

      params_name = op["operationId"].sub(/\A./) { |c| c.upcase } + "Params"
      op["x-fern-request-name"] = params_name
      op["x-fern-sdk-request-name"] = params_name
    end
  end

  File.write(path, JSON.generate(spec))
'

rm -rf .fern-out
CI=true npx -y fern-api@5.67.1 generate --local

# Keep only the SDK itself: no test scaffolding or generated meta docs.
rm -rf .fern-out/wiremock .fern-out/.fern
rm -f .fern-out/CONTRIBUTING.md .fern-out/README.md .fern-out/reference.md
find .fern-out -name "*_test.go" -delete
rm -rf .fern-out/client/root_test

find . -maxdepth 1 -name "*.go" -delete
rm -rf client core option internal wiremock
cp -r .fern-out/. .

rm -rf .fern-out openapi.tmp.json
go mod tidy
gofmt -w ./*.go client core option internal
