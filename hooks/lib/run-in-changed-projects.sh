#!/bin/bash

# Usage: bash run-in-changed-projects.sh MARKER COMMAND [ARG...] -- [FILE...]
# Example: bash run-in-changed-projects.sh go.mod go vet ./... -- "$@"
# Runs COMMAND once in each changed file's nearest project directory.
# MARKER is a filename or glob (quote globs, e.g. '*.tf').
# For commands containing -- or shell syntax, use: MARKER bash -c 'COMMAND' -- FILE...
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

has_marker() {
	local candidate
	for candidate in "$1"/* "$1"/.[!.]* "$1"/..?*; do
		# The unquoted right-hand side intentionally matches a filename glob.
		if [[ -f "$candidate" && "${candidate##*/}" == $marker ]]; then
			return 0
		fi
	done
	return 1
}

projects=()
for file in "$@"; do
	# Ignore downloaded Terraform modules, but not .terraform.lock.hcl.
	if [[ "$marker" == '*.tf' && "/$file" == */.terraform/* ]]; then
		continue
	fi
	dir=$(dirname -- "$file")
	while ! has_marker "$dir" && [[ "$dir" != . && "$dir" != / ]]; do
		dir=$(dirname -- "$dir")
	done
	has_marker "$dir" || continue
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
