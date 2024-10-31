#!/bin/bash

if ! command -v markdownlint-cli2 &>/dev/null; then
	echo "markdownlint-cli2 not installed or available in the PATH" >&2
	exit 1
fi

# shellcheck disable=SC2016
markdownlint-cli2 --ignores node_modules src --fix "$@"
