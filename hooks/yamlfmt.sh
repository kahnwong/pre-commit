#!/bin/bash

if ! command -v yamlfmt &>/dev/null; then
	echo "yamlfmt not installed or available in the PATH" >&2
	exit 1
fi

# shellcheck disable=SC2016
yamlfmt "$@"
