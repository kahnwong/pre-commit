#!/bin/bash

if ! command -v uv &>/dev/null; then
	echo "uv not installed or available in the PATH" >&2
	exit 1
fi

# shellcheck disable=SC2016
uv run mypy .
