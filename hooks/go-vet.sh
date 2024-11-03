#!/bin/bash

if ! command -v go &>/dev/null; then
	echo "go not installed or available in the PATH" >&2
	exit 1
fi

# shellcheck disable=SC2016
# shellcheck disable=SC2038
find . -name 'go.mod' | xargs dirname | sort -u | xargs -L 1 bash -c 'cd "$0" && go vet'
