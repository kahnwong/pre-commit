#!/bin/bash

if ! command -v uv &>/dev/null; then
	echo "uv not installed or available in the PATH" >&2
	exit 1
fi

# shellcheck disable=SC2016
# shellcheck disable=SC2038
bash "$(dirname -- "${BASH_SOURCE[0]}")/lib/run-in-changed-projects.sh" 'uv.lock' bash -c 'uv export --no-hashes --no-dev --no-emit-project --output-file=requirements.txt' -- "$@"
