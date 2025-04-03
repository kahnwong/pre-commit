#!/bin/bash

if ! command -v tofu &>/dev/null; then
	echo "tofu not installed or available in the PATH" >&2
	exit 1
fi

# shellcheck disable=SC2016
# shellcheck disable=SC2038
find . -name '*.tf' -not -path '*.terraform*' | xargs dirname | sort -u | xargs -L 1 bash -c 'cd "$0" && output=`tofu validate 2>&1` || echo -n "$output"'
