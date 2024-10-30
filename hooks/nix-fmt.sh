#!/bin/bash

if ! command -v nixfmt &>/dev/null; then
	echo "nixfmt not installed or available in the PATH" >&2
	exit 1
fi

nixfmt "$@"
