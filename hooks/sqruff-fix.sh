#!/bin/bash

if ! command -v sqruff &>/dev/null; then
	echo "sqruff not installed or available in the PATH" >&2
	exit 1
fi

# shellcheck disable=SC2016
# shellcheck disable=SC2038
sqruff fix --force "$@"
