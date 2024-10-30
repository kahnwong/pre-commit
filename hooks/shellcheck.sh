#!/bin/bash

if ! command -v shellcheck &>/dev/null; then
	echo "shellcheck not installed or available in the PATH" >&2
	exit 1
fi

shellcheck "$@"