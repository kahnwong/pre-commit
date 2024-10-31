#!/bin/bash

if ! command -v yamllint &>/dev/null; then
	echo "yamllint not installed or available in the PATH" >&2
	exit 1
fi

# shellcheck disable=SC2016
yamllint "$@"
