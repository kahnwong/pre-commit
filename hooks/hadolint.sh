#!/bin/bash

if ! command -v hadolint &>/dev/null; then
	echo "hadolint not installed or available in the PATH" >&2
	exit 1
fi

hadolint "$@"