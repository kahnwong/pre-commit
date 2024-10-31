#!/bin/bash

if ! command -v goimports &>/dev/null; then
	echo "goimports not installed or available in the PATH" >&2
	exit 1
fi

# shellcheck disable=SC2016
goimports -l -w "$@"
