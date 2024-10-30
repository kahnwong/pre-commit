package main

// toolchains
func goHooksManifest() map[string]hook {
	return map[string]hook{
		"go-fmt": {
			Type:       "go",
			Executable: "gofmt",
			Command:    "gofmt -l -w \"$@\"",
		},
		"go-imports": {
			Type:       "go",
			Executable: "goimports",
			Command:    "goimports -l -w \"$@\"",
		},
		"go-vet": {
			Type:       "go",
			Executable: "go",
			Command:    "go vet",
		},
		"golangci-lint": {
			Type:       "go",
			Executable: "golangci-lint",
			Command:    "golangci-lint run \"$@\"",
		},
		"go-unit-tests": {
			Type:       "go",
			Executable: "go",
			Command:    "go test ./...",
		},
		"go-build": {
			Type:       "go",
			Executable: "go",
			Command:    "go build",
		},
		"go-mod-tidy": {
			Type:       "go",
			Executable: "go",
			Command:    "go mod tidy",
		},
	}
}

func rustHooksManifest() map[string]hook {
	return map[string]hook{
		"rust-fmt": {
			Type:       "rust",
			Executable: "cargo",
			Command:    "cargo fmt",
		},
		"cargo-check": {
			Type:       "rust",
			Executable: "cargo",
			Command:    "cargo check",
		},
		"cargo-clippy": {
			Type:       "rust",
			Executable: "cargo",
			Command:    "cargo clippy -- -D warnings",
		},
	}
}

// misc
func opsHooksManifest() map[string]hook {
	return map[string]hook{
		"hadolint": {
			Type:       "dockerfile",
			Executable: "hadolint",
			Command:    "hadolint \"$@\"",
		},
		"shellcheck": {
			Files:      "\\.sh$",
			Executable: "shellcheck",
			Command:    "shellcheck \"$@\"",
		},
		"shfmt": {
			Files:      "\\.sh$",
			Executable: "shfmt",
			Command:    "shfmt -w \"$@\"",
		},
	}
}

func miscHooksManifest() map[string]hook {
	return map[string]hook{
		"yarn-prettier": {
			Type:       "ts",
			Executable: "yarn",
			Command:    "yarn prettier . --write",
		},
		"nix-fmt": {
			Files:      "\\.nix$",
			Executable: "nixfmt",
			Command:    "nixfmt \"$@\"",
		},
	}
}
