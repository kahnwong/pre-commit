#!/bin/bash

if ! command -v oxfmt &>/dev/null; then
	echo "oxfmt not installed or available in the PATH" >&2
	exit 1
fi

# shellcheck disable=SC2016
# shellcheck disable=SC2038
files=()
for file in "$@"; do
	if [[ "${file##*/}" == README.md ]]; then
		for tf in "$(dirname -- "$file")"/*.tf; do
			if [[ -f "$tf" ]]; then
				continue 2
			fi
		done
	fi
	files+=("$file")
done
if [[ ${#files[@]} -gt 0 ]]; then
	oxfmt "${files[@]}"
fi
