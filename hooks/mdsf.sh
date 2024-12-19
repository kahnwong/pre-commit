#!/bin/bash

if ! command -v mdsf &>/dev/null; then
	echo "mdsf not installed or available in the PATH" >&2
	exit 1
fi

# shellcheck disable=SC2016
# shellcheck disable=SC2038
mdsf format "$@"