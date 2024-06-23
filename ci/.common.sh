#!/usr/bin/env bash

set -euo pipefail

project="${PROJECT_NAME:-vnworkday}"

build_container() {
  echo "🐳 Building Pre-commit container..."
  docker compose --file ./ci/docker-compose.yaml --project-name "${project}" build
  echo "🐳 Pre-commit container built successfully."
}

start_container() {
  echo "🐳 Starting Pre-commit container..."
  docker compose --file ./ci/docker-compose.yaml --project-name "${project}" up --detach --quiet-pull
  echo "🐳 Pre-commit container started successfully."
}

stop_container() {
  echo "🐳 Stopping Pre-commit container..."
  docker compose --file ./ci/docker-compose.yaml --project-name "${project}" down
  echo "🐳 Pre-commit container stopped successfully."
}

get_container_id() {
  echo "🐳 Getting Pre-commit container ID..."
  CONTAINER_ID=$(docker ps --filter "label=com.vnworkday.docker.name=pre-commit" -q)
  if [ -z "$CONTAINER_ID" ]; then
      echo "⚠️ Pre-commit container not found. Please run the container first."
      exit 1
  fi
  echo "🐳 Pre-commit container ID: $CONTAINER_ID"
  export CONTAINER_ID
}

error_report() {
  echo "🚨 Error on line $(caller)" >&2
}

## When a command fails, the ERR signal is triggered and the trap command catches it
## to execute the error_report function, which prints the error message
## and the line number where the error occurred.
trap error_report ERR