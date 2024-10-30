package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type hook struct {
	Files      string
	Executable string
	Command    string
}

func main() {
	hooks := map[string]hook{
		"go-fmt": {
			Files:   "\\.go$",
			Command: "gofmt -l -w \"$@\"",
		},
		"go-imports": {
			Files:   "\\.go$",
			Command: "goimports -l -w \"$@\"",
		},
	}

	// generate config
	createDir("hooks")
	hooksConfig := generateHooksConfig(hooks)

	// write to file
	err := os.WriteFile(".pre-commit-hooks.yaml", []byte(hooksConfig), 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully generated pre-commit hooks")
}

func generateHooksConfig(services map[string]hook) string {
	config := "---\n"
	keys := make([]string, 0, len(services))
	for k := range services {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		// set vars
		id := k
		name := strings.ReplaceAll(k, "-", " ")
		executable := strings.Split(services[k].Command, " ")[0]

		// create hook
		hookScript := fmt.Sprintf(`#!/bin/bash

set -e -o pipefail

if ! command -v %s &>/dev/null; then
	echo "%s not installed or available in the PATH" >&2
	exit 1
fi

output="$(%s)"
echo "$output"
[[ -z "$output" ]]
`, executable, executable, services[k].Command)
		err := os.WriteFile(fmt.Sprintf("./hooks/%s.sh", k), []byte(hookScript), 0775)
		if err != nil {
			panic(err)
		}

		// append hooks config
		config += fmt.Sprintf(`- id: %s
  name: '%s'
  entry: ./hooks/%s.sh
  files: '%s'
  language: 'script'
`, id, name, id, services[k].Files)
	}

	return config
}
