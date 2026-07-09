#!/bin/sh
# Regenerates types.go from the DocuSeal OpenAPI spec.
# Usage: ./generate-types.sh [path-or-url-to-openapi-json]
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
#
# Optional fields are generated as plain values (prefer-skip-optional-pointer),
# except where a zero value and an absent field mean different things on the
# wire: booleans (`send_email: false` would be dropped by omitempty and the
# server would apply its default) and struct-typed objects (omitempty never
# omits a zero struct, so an untouched `message` would be sent as `{}`).
# Those keep pointers via x-go-type-skip-optional-pointer: false.
ruby -rjson -e '
  path = ARGV[0]
  spec = JSON.parse(File.read(path))
  spec.delete("webhooks")

  deref = ->(node) do
    return node unless node.is_a?(Hash) && node["$ref"]
    spec.dig(*node["$ref"].delete_prefix("#/").split("/"))
  end

  keep_pointer = ->(schema) do
    schema = deref.call(schema)
    return false unless schema.is_a?(Hash)
    Array(schema["type"]).include?("boolean") || schema["properties"].is_a?(Hash)
  end

  walk = ->(node) do
    case node
    when Hash
      if node["properties"].is_a?(Hash)
        node["properties"].each_value do |prop|
          prop["x-go-type-skip-optional-pointer"] = false if keep_pointer.call(prop)
        end
      end
      if node["in"] && node["schema"].is_a?(Hash)
        node["schema"]["x-go-type-skip-optional-pointer"] = false if keep_pointer.call(node["schema"])
      end
      node.each_value { |value| walk.call(value) }
    when Array
      node.each { |value| walk.call(value) }
    end
  end
  walk.call(spec)

  File.write(path, JSON.generate(spec))
' "$TMP_DIR/openapi.json"

cat > "$TMP_DIR/config.yaml" << 'EOF'
package: docuseal
output: types.go
generate:
  models: true
output-options:
  prefer-skip-optional-pointer: true
EOF

# OpenAPI 3.1 support is merged into oapi-codegen master but not released yet
# (latest release v2.7.2 fails on `type: [T, null]`), hence the pseudo-version
# pin. Bump to the next tagged release when it ships.
go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.7.1-0.20260708002751-cdf0907a107e \
  -config "$TMP_DIR/config.yaml" "$TMP_DIR/openapi.json"

gofmt -w types.go
