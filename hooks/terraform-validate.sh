#!/bin/bash

if ! command -v terraform &>/dev/null; then
	echo "terraform not installed or available in the PATH" >&2
	exit 1
fi

# shellcheck disable=SC2016
# shellcheck disable=SC2038
find . -name '*.tf' -not -path '*.terraform*' | xargs dirname | sort -u | xargs -L 1 bash -c 'cd "$0" && output=`terraform validate 2>&1` || echo -n "$output"'
