#!/bin/bash

set -e -o pipefail

if ! command -v nixfmt &>/dev/null; then
	echo "nixfmt not installed or available in the PATH" >&2
	exit 1
fi

output="$(nixfmt "$@")"
echo "$output"
[[ -z "$output" ]]
