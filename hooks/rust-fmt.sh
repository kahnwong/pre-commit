#!/bin/bash

set -e -o pipefail

if ! command -v cargo &>/dev/null; then
	echo "cargo not installed or available in the PATH" >&2
	exit 1
fi

output="$(cargo fmt)"
echo "$output"
[[ -z "$output" ]]
