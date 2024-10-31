#!/bin/bash

if ! command -v terraform &>/dev/null; then
	echo "terraform not installed or available in the PATH" >&2
	exit 1
fi

# shellcheck disable=SC2016
find . -name '*.tf' -not -path '*.terraform*' -printf '%h\n' | sort -u | xargs terraform fmt
