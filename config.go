package main

func hooksManifest() map[string]hook {
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
			Files:      "\\.go$",
			Executable: "go",
			Command:    "go mod tidy",
		},
	}

	return hooks
}
