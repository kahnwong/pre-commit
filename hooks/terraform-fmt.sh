#!/bin/bash

if ! command -v terraform &>/dev/null; then
	echo "terraform not installed or available in the PATH" >&2
	exit 1
fi

# shellcheck disable=SC2016
# shellcheck disable=SC2038
bash "$(dirname -- "${BASH_SOURCE[0]}")/lib/run-in-changed-projects.sh" '*.tf' bash -c 'terraform fmt' -- "$@"
