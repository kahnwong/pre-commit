#!/bin/bash

if ! command -v cargo &>/dev/null; then
	echo "cargo not installed or available in the PATH" >&2
	exit 1
fi

# shellcheck disable=SC2016
# shellcheck disable=SC2038
find . -name 'Cargo.toml' | xargs dirname | sort -u | xargs -L 1 bash -c 'cd "$0" && cargo clippy -- -D warnings'
