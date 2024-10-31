#!/bin/bash

if ! command -v black &>/dev/null; then
	echo "black not installed or available in the PATH" >&2
	exit 1
fi

# shellcheck disable=SC2016
black "$@"
