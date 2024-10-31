#!/bin/bash

if ! command -v nbstripout &>/dev/null; then
	echo "nbstripout not installed or available in the PATH" >&2
	exit 1
fi

# shellcheck disable=SC2016
nbstripout "$@"
