#!/bin/bash

set -e -o pipefail

if ! command -v yarn &>/dev/null; then
	echo "yarn not installed or available in the PATH" >&2
	exit 1
fi

output="$(yarn prettier . --write)"
echo "$output"
[[ -z "$output" ]]
