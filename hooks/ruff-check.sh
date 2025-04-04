#!/bin/bash

if ! command -v ruff &>/dev/null; then
	echo "ruff not installed or available in the PATH" >&2
	exit 1
fi

# shellcheck disable=SC2016
# shellcheck disable=SC2038
ruff check --extend-select I,F --fix
