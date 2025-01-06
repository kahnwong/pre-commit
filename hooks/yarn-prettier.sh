#!/bin/bash

if ! command -v yarn &>/dev/null; then
	echo "yarn not installed or available in the PATH" >&2
	exit 1
fi

# shellcheck disable=SC2016
# shellcheck disable=SC2038
yarn prettier --write "**/*.{js,ts,vue,scss,html,md,json}" --ignore-path .gitignore
