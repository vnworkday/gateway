#!/usr/bin/env bash

set -euo pipefail

script_dir=$(dirname "$0")
source "${script_dir}/.common.sh"

build_container
start_container