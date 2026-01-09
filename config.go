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
			Command:    "find . -name 'go.mod' | xargs dirname | sort -u | xargs -L 1 bash -c 'cd \"$0\" && go vet'",
		},
		"golangci-lint": {
			Type:       "go",
			Executable: "golangci-lint",
			Command:    "find . -name 'go.mod' | xargs dirname | sort -u | xargs -L 1 bash -c 'cd \"$0\" && golangci-lint run'",
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
			Command:    "find . -name 'pyproject.toml' | xargs dirname | sort -u | xargs -L 1 bash -c 'cd \"$0\" && uv lock'",
		},
		"uv-export": {
			Files:      "uv\\.lock$",
			Executable: "uv",
			Command:    "find . -name 'uv.lock' | xargs dirname | sort -u | xargs -L 1 bash -c 'cd \"$0\" && uv export --no-hashes --no-dev --no-emit-project --output-file=requirements.txt'",
		},
		"mypy": {
			Type:       "python",
			Executable: "uv",
			Command:    "uv run mypy .",
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
			Command:    "find . -name 'Cargo.toml' | xargs dirname | sort -u | xargs -L 1 bash -c 'cd \"$0\" && cargo fmt'",
		},
		"cargo-check": {
			Type:       "rust",
			Executable: "cargo",
			Command:    "find . -name 'Cargo.toml' | xargs dirname | sort -u | xargs -L 1 bash -c 'cd \"$0\" && cargo check'",
		},
		"cargo-clippy": {
			Type:       "rust",
			Executable: "cargo",
			Command:    "find . -name 'Cargo.toml' | xargs dirname | sort -u | xargs -L 1 bash -c 'cd \"$0\" && cargo clippy -- -D warnings'",
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
		"tofu-fmt": {
			Files:      "(\\.tf|\\.tfvars)$",
			Executable: "tofu",
			Command:    "find . -name '*.tf' -not -path '*.terraform*' | xargs dirname | sort -u | xargs tofu fmt",
		},
		"terraform-validate": {
			Files:      "\\.(tf(vars)?|terraform\\.lock\\.hcl)$",
			Executable: "terraform",
			Command:    "find . -name '*.tf' -not -path '*.terraform*' | xargs dirname | sort -u | xargs -L 1 bash -c 'cd \"$0\" && output=`terraform validate 2>&1` || echo -n \"$output\"'",
		},
		"tofu-validate": {
			Files:      "\\.(tf(vars)?|terraform\\.lock\\.hcl)$",
			Executable: "tofu",
			Command:    "find . -name '*.tf' -not -path '*.terraform*' | xargs dirname | sort -u | xargs -L 1 bash -c 'cd \"$0\" && output=`tofu validate 2>&1` || echo -n \"$output\"'",
		},
		"terraform-docs": {
			Files:      "(\\.tf|\\.terraform\\.lock\\.hcl)$",
			Executable: "terraform-docs",
			Command:    "find . -name '*.tf' -not -path '*.terraform*' | xargs dirname | sort -u | xargs -L 1 bash -c 'cd \"$0\" && output=`terraform-docs markdown table --html=false --anchor=false --output-file README.md --output-mode inject . 2>&1` || echo -n \"$output\"'",
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

func miscHooksManifest() map[string]hook {
	return map[string]hook{
		"yarn-prettier": {
			Files:      "\\.([jt]sx?|vue)$",
			Executable: "yarn",
			Command:    "yarn prettier --write \"**/*.{js,ts,vue,scss,html,md,json}\" --ignore-path .gitignore",
		},
		"graphql": {
			Files:      "\\.graphql",
			Executable: "prettier",
			Command:    "prettier -w \"$@\"",
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
