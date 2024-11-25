#!/bin/bash

if ! command -v trufflehog &>/dev/null; then
	echo "trufflehog not installed or available in the PATH" >&2
	exit 1
fi

# shellcheck disable=SC2016
# shellcheck disable=SC2038
trufflehog filesystem "$@" --no-update --fail
