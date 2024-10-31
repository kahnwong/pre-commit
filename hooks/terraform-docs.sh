#!/bin/bash

if ! command -v terraform-docs &>/dev/null; then
	echo "terraform-docs not installed or available in the PATH" >&2
	exit 1
fi

# shellcheck disable=SC2016
find . -name '*.tf' -not -path '*.terraform*' -printf '%h\n' | sort -u | xargs -L 1 bash -c 'terraform-docs "$0"'
