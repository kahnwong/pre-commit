#!/bin/bash

set -e -o pipefail

if ! command -v golangci-lint &>/dev/null; then
	echo "golangci-lint not installed or available in the PATH" >&2
	exit 1
fi

output="$(golangci-lint run "$@")"
echo "$output"
[[ -z "$output" ]]
