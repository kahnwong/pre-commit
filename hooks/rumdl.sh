#!/bin/bash

if ! command -v rumdl &>/dev/null; then
	echo "rumdl not installed or available in the PATH" >&2
	exit 1
fi

# shellcheck disable=SC2016
# shellcheck disable=SC2038
rumdl fmt "$@"
