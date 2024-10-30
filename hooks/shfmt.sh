#!/bin/bash

if ! command -v shfmt &>/dev/null; then
	echo "shfmt not installed or available in the PATH" >&2
	exit 1
fi

shfmt -w "$@"
