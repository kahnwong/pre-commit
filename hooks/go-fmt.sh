#!/bin/bash

if ! command -v gofmt &>/dev/null; then
	echo "gofmt not installed or available in the PATH" >&2
	exit 1
fi

# shellcheck disable=SC2016
# shellcheck disable=SC2038
gofmt -l -w "$@"
