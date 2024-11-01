#!/bin/bash

if ! command -v cargo &>/dev/null; then
	echo "cargo not installed or available in the PATH" >&2
	exit 1
fi

# shellcheck disable=SC2016
# shellcheck disable=SC2038
cargo check
