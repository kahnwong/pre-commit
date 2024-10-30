#!/bin/bash

if ! command -v golangci-lint &>/dev/null; then
	echo "golangci-lint not installed or available in the PATH" >&2
	exit 1
fi

golangci-lint run "$@"
