#!/bin/bash

if ! command -v oxfmt &>/dev/null; then
	echo "oxfmt not installed or available in the PATH" >&2
	exit 1
fi

# shellcheck disable=SC2016
# shellcheck disable=SC2038
oxfmt "$@"
