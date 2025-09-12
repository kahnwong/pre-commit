#!/bin/bash

if ! command -v kingfisher &>/dev/null; then
	echo "kingfisher not installed or available in the PATH" >&2
	exit 1
fi

# shellcheck disable=SC2016
# shellcheck disable=SC2038
kingfisher scan "$@"
