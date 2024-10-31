#!/bin/bash

if ! command -v yarn &>/dev/null; then
	echo "yarn not installed or available in the PATH" >&2
	exit 1
fi

# shellcheck disable=SC2016
yarn prettier . --write
