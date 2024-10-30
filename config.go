package main

func goHooksManifest() map[string]hook {
	return map[string]hook{
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
}

func rustHooksManifest() map[string]hook {
	return map[string]hook{
		"rust-fmt": {
			Files:      "\\.rs$",
			Executable: "cargo",
			Command:    "cargo fmt",
		},
		"cargo-check": {
			Files:      "^Cargo",
			Executable: "cargo",
			Command:    "cargo check",
		},
		"cargo-clippy": {
			Files:      "\\.rs$",
			Executable: "cargo",
			Command:    "cargo clippy -- -D warnings",
		},
	}
}

func miscHooksManifest() map[string]hook {
	return map[string]hook{
		"yarn-prettier": {
			Executable: "yarn",
			Command:    "yarn prettier . --write",
		},
		"nix-fmt": {
			Executable: "nixfmt",
			Command:    "nixfmt \"$@\"",
		},
	}
}
