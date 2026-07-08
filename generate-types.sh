#!/bin/sh
# Regenerates types.go from the DocuSeal OpenAPI spec.
# Usage: ./scripts/generate-types.sh [path-or-url-to-openapi-json]
set -e

cd "$(dirname "$0")"

SPEC="${1:-https://console.docuseal.com/openapi.yml?format=json}"
TMP_DIR="$(mktemp -d)"
trap 'rm -rf "$TMP_DIR"' EXIT

case "$SPEC" in
  http*) curl -sf "$SPEC" -o "$TMP_DIR/openapi.json" ;;
  *) cp "$SPEC" "$TMP_DIR/openapi.json" ;;
esac

# The SDK exposes only the REST API client: drop webhook payload schemas.
ruby -rjson -e '
  path = ARGV[0]
  spec = JSON.parse(File.read(path))
  spec.delete("webhooks")
  File.write(path, JSON.generate(spec))
' "$TMP_DIR/openapi.json"

# OpenAPI 3.1 support is merged into oapi-codegen master but not released yet
# (latest release v2.7.2 fails on `type: [T, null]`), hence the pseudo-version
# pin. Bump to the next tagged release when it ships.
go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.7.1-0.20260708002751-cdf0907a107e \
  -generate types -package docuseal -o types.go "$TMP_DIR/openapi.json"

gofmt -w types.go
