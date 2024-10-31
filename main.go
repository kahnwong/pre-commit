package main

import (
	"fmt"
	"maps"
	"os"
	"sort"
	"strings"
)

type hook struct {
	Type       string
	Files      string
	Executable string
	Command    string
}

func main() {
	hooks := map[string]hook{}
	hooksManifest := []map[string]hook{
		goHooksManifest(),
		rustHooksManifest(),
		terraformHooksManifest(),
		opsHooksManifest(),
		miscHooksManifest(),
	}
	for _, i := range hooksManifest {
		maps.Insert(hooks, maps.All(i))
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
	config := ""
	keys := make([]string, 0, len(services))
	for k := range services {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("- %s\n", k)

		// set vars
		id := k
		name := strings.ReplaceAll(k, "-", " ")
		executable := services[k].Executable

		// create hook
		hookScript := fmt.Sprintf(`#!/bin/bash

if ! command -v %s &>/dev/null; then
	echo "%s not installed or available in the PATH" >&2
	exit 1
fi

# shellcheck disable=SC2016
%s
`, executable, executable, services[k].Command)
		err := os.WriteFile(fmt.Sprintf("./hooks/%s.sh", k), []byte(hookScript), 0775)
		if err != nil {
			panic(err)
		}

		// append hooks config
		config += fmt.Sprintf(`- id: %s
  name: '%s'
  entry: ./hooks/%s.sh
  types: [%s]
  files: '%s'
  language: 'script'
`, id, name, id, services[k].Type, services[k].Files)
	}

	return config
}
