#!/usr/bin/env bash

set -eufo pipefail

usage() {
    echo "📋 Usage: $0 -s/--service <service> -r/--resource <resource>"
    echo "✅  Example: $0 -s account -r user"
    exit 0
}

service=""
resource=""

while [ "$#" -gt 0 ]; do
    case "$1" in
        -s | --service) shift; service="$1";;
        -r | --resource) shift; resource="$1";;
        *) echo "❗️ Invalid argument: $1"; usage;;
    esac
    shift
done

if [ -z "$service" ] || [ -z "$resource" ]; then
    if [ -z "$service" ]; then
        echo "❗️ Missing required arguments: -s/--service <service>"
        missing_arguments+=("-s/--service")
    fi
    if [ -z "$resource" ]; then
        echo "❗️ Missing required arguments: -r/--resource <resource>"
        missing_arguments+=("-r/--resource")
    fi
    usage
fi

# 1. Create router file for the service in ./routes/<service>/<resource>.router.go
mkdir -p ./internal/routes/"$service"
touch ./internal/routes/"$service"/"$resource".router.go

# 2. Create handler file for the service in ./handlers/<service>/<resource>.handler.go
mkdir -p ./internal/handlers/"$service"
touch ./internal/handlers/"$service"/"$resource".handler.go

echo "🚀 Created router and handler files for $service $resource"
