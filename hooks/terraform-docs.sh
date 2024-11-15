#!/bin/bash

if ! command -v terraform-docs &>/dev/null; then
	echo "terraform-docs not installed or available in the PATH" >&2
	exit 1
fi

# shellcheck disable=SC2016
# shellcheck disable=SC2038
find . -name '*.tf' -not -path '*.terraform*' | xargs dirname | sort -u | xargs -L 1 bash -c 'cd "$0" && output=`terraform-docs markdown table --html=false --anchor=false --output-file README.md --output-mode inject . 2>&1` || echo -n "$output"'
