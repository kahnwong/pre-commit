#!/bin/bash

if ! command -v uv &>/dev/null; then
	echo "uv not installed or available in the PATH" >&2
	exit 1
fi

# shellcheck disable=SC2016
# shellcheck disable=SC2038
find . -name 'pyproject.toml' | xargs dirname | sort -u | xargs -L 1 bash -c 'cd "$0" && uv lock'
