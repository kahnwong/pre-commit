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

func pythonHooksManifest() map[string]hook {
	return map[string]hook{
		"ruff-check": {
			Type:       "python",
			Executable: "ruff",
			Command:    "ruff check --fix",
		},
		"ruff-format": {
			Type:       "python",
			Executable: "ruff",
			Command:    "ruff format",
		},
		"uv-export": {
			Files:      "^uv\\.lock$",
			Executable: "uv",
			Command:    "uv export --no-hashes --output-file=requirements.txt",
		},
		"mypy": {
			Type:       "python",
			Executable: "uv",
			Command:    "uv run mypy .",
		},
		"nbstripout": {
			Type:       "python",
			Executable: "nbstripout",
			Command:    "nbstripout \"$@\"",
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

func terraformHooksManifest() map[string]hook {
	return map[string]hook{
		"terraform-fmt": {
			Files:      "(\\.tf|\\.tfvars)$",
			Executable: "terraform",
			Command:    "find . -name '*.tf' -not -path '*.terraform*' | xargs dirname | sort -u | xargs terraform fmt",
		},
		"terraform-validate": {
			Files:      "\\.(tf(vars)?|terraform\\.lock\\.hcl)$",
			Executable: "terraform",
			Command:    "find . -name '*.tf' -not -path '*.terraform*' | xargs dirname | sort -u | xargs -L 1 bash -c 'cd \"$0\" && output=`terraform validate 2>&1` || echo -n \"$output\"'",
		},
		"terraform-docs": {
			Files:      "(\\.tf|\\.terraform\\.lock\\.hcl)$",
			Executable: "terraform-docs",
			Command:    "find . -name '*.tf' -not -path '*.terraform*' | xargs dirname | sort -u | xargs -L 1 bash -c 'output=`terraform-docs \"$0\" 2>&1` || echo -n \"$output\"'",
		},
		"terraform-trivy": {
			Files:      "(\\.tf|\\.tfvars)$",
			Executable: "trivy",
			Command:    "trivy config --tf-exclude-downloaded-modules --exit-code 1 .",
		},
	}
}

// misc
func opsHooksManifest() map[string]hook {
	return map[string]hook{
		"yamlfmt": {
			Files:      "(\\.yaml|\\.yml)$",
			Executable: "yamlfmt",
			Command:    "yamlfmt \"$@\"",
		},
		"yamllint": {
			Files:      "(\\.yaml|\\.yml)$",
			Executable: "yamllint",
			Command:    "yamllint \"$@\"",
		},
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
		"markdownlint-cli2": {
			Type:       "markdown",
			Executable: "markdownlint-cli2",
			Command:    "markdownlint-cli2 --ignores node_modules src --fix \"$@\"",
		},
	}
}
