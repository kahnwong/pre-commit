#!/bin/bash

if ! command -v go &>/dev/null; then
	echo "go not installed or available in the PATH" >&2
	exit 1
fi

# shellcheck disable=SC2016
go mod tidy
