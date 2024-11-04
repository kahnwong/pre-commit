#!/bin/bash

if ! command -v terraform-docs &>/dev/null; then
	echo "terraform-docs not installed or available in the PATH" >&2
	exit 1
fi

# shellcheck disable=SC2016
# shellcheck disable=SC2038
terraform-docs .
