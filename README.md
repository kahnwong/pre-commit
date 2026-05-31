# pre-commit

Since I use multiple languages (Bash, Go, Python, Rust, Terraform) it's not very fun to remember all the hooks I need to use...

## Usage

```yaml
  - repo: https://github.com/kahnwong/pre-commit
    rev: "master"
    hooks:
      # -- go -- #
      - id: go-build
      - id: go-fmt
      - id: go-imports
      - id: go-mod-tidy
      - id: go-unit-tests
      - id: go-vet
      - id: golangci-lint
      # -- python -- #
      - id: ruff-check
      - id: ruff-format
      - id: uv-lock
      - id: uv-export
      - id: ty
      - id: nbstripout
      # -- rust -- #
      - id: rust-fmt
      - id: cargo-check
      - id: cargo-clippy
      # -- terraform -- #
      - id: terraform-fmt
      - id: terraform-validate
      - id: tofu-fmt
      - id: tofu-validate
      - id: terraform-docs
      - id: checkov
      # -- ops -- #
      - id: yamlfmt
      - id: hadolint
      - id: kingfisher
      - id: shellcheck
      - id: shfmt
      - id: trufflehog
      # -- nix -- #
      - id: nix-fmt
      - id: statix
      - id: deadnix
      # -- markdown -- #
      - id: markdownlint-cli2
      - id: mdsf
      - id: rumdl
      # -- sql -- #
      - id: sqruff-lint
      - id: sqruff-fix
      # -- misc -- #
      - id: typos
```

## Refs

- <https://github.com/dnephin/pre-commit-golang>
- <https://pre-commit.com/#plugins>
