#!/usr/bin/env bash

set -eufo pipefail

swag_version="v1.16.3"
redoc_version="1.18.0"

init() {
  docker run --platform linux/amd64 --rm -v "$(pwd)":/code ghcr.io/swaggo/swag:${swag_version} init
}

fmt() {
  docker run --platform linux/amd64 --rm -v "$(pwd)":/code ghcr.io/swaggo/swag:${swag_version} fmt
}

lint() {
  docker run --rm -v "$(pwd)/docs":/spec redocly/cli:${redoc_version} lint --extends recommended swagger.yaml
}

build() {
  docker run --rm -v "$(pwd)/docs":/spec redocly/cli:${redoc_version} build-docs swagger.yaml -o index.html
}

init
fmt
lint
build