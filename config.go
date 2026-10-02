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
			Command:    runInChangedProjects("go vet ./...", "go.mod"),
		},
		"golangci-lint": {
			Type:       "go",
			Executable: "golangci-lint",
			Command:    runInChangedProjects("golangci-lint run --fix", "go.mod"),
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
			Command:    "ruff check --extend-select I,F --fix",
		},
		"ruff-format": {
			Type:       "python",
			Executable: "ruff",
			Command:    "ruff format",
		},
		"uv-lock": {
			Files:      "pyproject\\.toml",
			Executable: "uv",
			Command:    runInChangedProjects("uv lock", "pyproject.toml"),
		},
		"uv-export": {
			Files:      "uv\\.lock$",
			Executable: "uv",
			Command:    runInChangedProjects("uv export --no-hashes --no-dev --no-emit-project --output-file=requirements.txt", "uv.lock"),
		},
		"ty": {
			Type:       "python",
			Executable: "uvx",
			Command:    "uvx ty check .",
		},
		"nbstripout": {
			Files:      "\\.ipynb",
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
			Command:    runInChangedProjects("cargo fmt", "Cargo.toml"),
		},
		"cargo-check": {
			Type:       "rust",
			Executable: "cargo",
			Command:    runInChangedProjects("cargo check", "Cargo.toml"),
		},
		"cargo-clippy": {
			Type:       "rust",
			Executable: "cargo",
			Command:    runInChangedProjects("cargo clippy -- -D warnings", "Cargo.toml"),
		},
	}
}

func terraformHooksManifest() map[string]hook {
	return map[string]hook{
		"terraform-fmt": {
			Files:      "(\\.tf|\\.tfvars)$",
			Executable: "terraform",
			Command:    runInChangedProjects("terraform fmt", "*.tf"),
		},
		"tofu-fmt": {
			Files:      "(\\.tf|\\.tfvars)$",
			Executable: "tofu",
			Command:    runInChangedProjects("tofu fmt", "*.tf"),
		},
		"terraform-validate": {
			Files:      "\\.(tf(vars)?|terraform\\.lock\\.hcl)$",
			Executable: "terraform",
			Command:    runInChangedProjects(`output=$(terraform validate 2>&1) || printf '%s' "$output"`, "*.tf"),
		},
		"tofu-validate": {
			Files:      "\\.(tf(vars)?|terraform\\.lock\\.hcl)$",
			Executable: "tofu",
			Command:    runInChangedProjects(`output=$(tofu validate 2>&1) || printf '%s' "$output"`, "*.tf"),
		},
		"terraform-docs": {
			Files:      "(\\.tf|\\.terraform\\.lock\\.hcl)$",
			Executable: "terraform-docs",
			Command:    runInChangedProjects(`output=$(terraform-docs markdown table --html=false --anchor=false --output-file README.md --output-mode inject . 2>&1) || printf '%s' "$output"`, "*.tf"),
		},
		"checkov": {
			Files:      "(\\.tf|\\.tfvars)$",
			Executable: "checkov",
			Command:    "checkov -f \"$@\" --quiet",
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
		//"yamllint": {
		//	Files:      "(\\.yaml|\\.yml)$",
		//	Executable: "yamllint",
		//	Command:    "yamllint \"$@\"",
		//},
		"hadolint": {
			Type:       "dockerfile",
			Executable: "hadolint",
			Command:    "hadolint \"$@\"",
		},
		"kingfisher": {
			Files:      ".*",
			Executable: "kingfisher",
			Command:    "kingfisher scan \"$@\"",
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
		"trufflehog": {
			Files:      ".*",
			Executable: "trufflehog",
			Command:    "trufflehog filesystem \"$@\" --no-update --fail",
		},
	}
}

func nixHooksManifest() map[string]hook {
	return map[string]hook{
		"nix-fmt": {
			Files:      "\\.nix$",
			Executable: "nixfmt",
			Command:    "nixfmt \"$@\"",
		},
		"statix": {
			Files:      "\\.nix$",
			Executable: "statix",
			Command:    "statix check .",
		},
		"deadnix": {
			Files:      "\\.nix$",
			Executable: "deadnix",
			Command:    "deadnix -eq \"$@\"",
		},
	}
}

func markdownHooksManifest() map[string]hook {
	return map[string]hook{
		"markdownlint-cli2": {
			Type:       "markdown",
			Executable: "markdownlint-cli2",
			Command:    "markdownlint-cli2 --ignores node_modules src --fix \"$@\"",
		},
		"mdsf": {
			Type:       "markdown",
			Executable: "mdsf",
			Command:    "mdsf format \"$@\"",
		},
		"rumdl": {
			Type:       "markdown",
			Executable: "rumdl",
			Command:    "rumdl fmt \"$@\"",
		},
	}
}

func sqlHooksManifest() map[string]hook {
	return map[string]hook{
		"sqruff-lint": {
			Type:       "sql",
			Executable: "sqruff",
			Command:    "sqruff lint \"$@\"",
		},
		"sqruff-fix": {
			Type:       "sql",
			Executable: "sqruff",
			Command:    "sqruff fix --force \"$@\"",
		},
	}
}

func miscHooksManifest() map[string]hook {
	return map[string]hook{
		"typos": {
			Type:       "markdown",
			Executable: "typos",
			Command:    "typos \"$@\"",
		},
		"oxfmt": {
			Files:      "\\.(js|jsx|ts|tsx|vue|css|scss|html|json|yaml|yml|md|toml|graphql|gql)$",
			Executable: "oxfmt",
			Command:    "oxfmt \"$@\"",
		},
	}
}
