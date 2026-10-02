#!/bin/bash

# Usage: bash run-in-changed-projects.sh MARKER COMMAND [ARG...] -- [FILE...]
# Example: bash run-in-changed-projects.sh go.mod go vet ./... -- "$@"
# Runs COMMAND once in each changed file's nearest project directory.
# Files without a matching marker are skipped; no files means no commands.

if [[ $# -lt 3 ]]; then
	echo "Usage: $0 MARKER COMMAND [ARG...] -- [FILE...]" >&2
	exit 2
fi

marker=$1
shift
command=()
while [[ $# -gt 0 && "$1" != -- ]]; do
	command+=("$1")
	shift
done
if [[ ${#command[@]} -eq 0 || $# -eq 0 ]]; then
	echo "Expected a command followed by -- and optional changed files" >&2
	exit 2
fi
shift

projects=()
for file in "$@"; do
	dir=$(dirname -- "$file")
	while [[ ! -f "$dir/$marker" && "$dir" != . && "$dir" != / ]]; do
		dir=$(dirname -- "$dir")
	done
	[[ -f "$dir/$marker" ]] || continue
	dir=$(cd -- "$dir" && pwd -P) || exit 1
	for project in "${projects[@]}"; do
		[[ "$project" == "$dir" ]] && continue 2
	done
	projects+=("$dir")
done

status=0
for project in "${projects[@]}"; do
	(cd -- "$project" && "${command[@]}") || status=$?
done
exit "$status"
