package main

func goHooksManifest() map[string]hook {
	hooks := map[string]hook{
		"go-fmt": {
			Files:      "\\.go$",
			Executable: "gofmt",
			Command:    "gofmt -l -w \"$@\"",
		},
		"go-imports": {
			Files:      "\\.go$",
			Executable: "goimports",
			Command:    "goimports -l -w \"$@\"",
		},
		"go-vet": {
			Files:      "\\.go$",
			Executable: "go",
			Command:    "go vet",
		},
		"golangci-lint": {
			Files:      "\\.go$",
			Executable: "golangci-lint",
			Command:    "golangci-lint run \"$@\"",
		},
		"go-unit-tests": {
			Files:      "\\.go$",
			Executable: "go",
			Command:    "go test ./...",
		},
		"go-build": {
			Files:      "\\.go$",
			Executable: "go",
			Command:    "go build",
		},
		"go-mod-tidy": {
			Files:      "\\.mod",
			Executable: "go",
			Command:    "go mod tidy -v \"$@\"",
		},
	}

	return hooks
}

func rustHooksManifest() map[string]hook {
	hooks := map[string]hook{
		"rust-fmt": {
			Files:      "\\.rs$",
			Executable: "cargo",
			Command:    "cargo fmt",
		},
	}

	return hooks
}
