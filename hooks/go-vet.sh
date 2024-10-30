#!/bin/bash

set -e -o pipefail

if ! command -v go &>/dev/null; then
	echo "go not installed or available in the PATH" >&2
	exit 1
fi

output="$(go vet "$@")"
echo "$output"
[[ -z "$output" ]]
