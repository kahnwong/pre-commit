#!/bin/bash

if ! command -v trivy &>/dev/null; then
	echo "trivy not installed or available in the PATH" >&2
	exit 1
fi

# shellcheck disable=SC2016
# shellcheck disable=SC2038
trivy config --tf-exclude-downloaded-modules --exit-code 1 .
