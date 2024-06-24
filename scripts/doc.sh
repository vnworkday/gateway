#!/usr/bin/env bash

set -eufo pipefail

curr_dir=$(pwd)

echo "Current directory: ${curr_dir}"

gen() {
  docker run --rm -v "$(pwd)":/code ghcr.io/swaggo/swag:latest init
}

fmt() {
  docker run --rm -v "$(pwd)":/code ghcr.io/swaggo/swag:latest fmt
}

gen
fmt