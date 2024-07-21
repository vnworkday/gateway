#!/usr/bin/env bash

set -eufo pipefail

swag_version="v1.16.3"
redoc_version="1.18.0"

init() {
  docker run --platform linux/amd64 --rm -v "$(pwd)":/code ghcr.io/swaggo/swag:${swag_version} init --generatedTime true --parseInternal true
  docker run --platform linux/amd64 --rm -v $(pwd):/usr/src/app mermade/swagger2openapi swagger2openapi --yaml --outfile docs/openapi.yaml docs/swagger.yaml
  docker run --platform linux/amd64 --rm -v $(pwd):/usr/src/app mermade/swagger2openapi swagger2openapi --outfile docs/openapi.json docs/swagger.json
}

fmt() {
  docker run --platform linux/amd64 --rm -v "$(pwd)":/code ghcr.io/swaggo/swag:${swag_version} fmt
}

lint() {
  docker run --rm -v "$(pwd)/docs":/spec redocly/cli:${redoc_version} lint --extends recommended openapi.yaml
}

build() {
  docker run --rm -v "$(pwd)/docs":/spec redocly/cli:${redoc_version} build-docs openapi.yaml -o index.html
}

init
fmt
lint
build