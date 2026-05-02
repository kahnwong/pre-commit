#!/bin/bash

if ! command -v statix &>/dev/null; then
	echo "statix not installed or available in the PATH" >&2
	exit 1
fi

# shellcheck disable=SC2016
# shellcheck disable=SC2038
statix check .
