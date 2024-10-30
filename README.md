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
    # -- rust -- #
    - id: rust-fmt
    - id: cargo-check
    - id: cargo-clippy
    # -- ops -- #
    - id: shellcheck
    - id: shfmt
    # -- misc -- #
    - id: nix-fmt
    - id: yarn-prettier
```

## Refs

- <https://github.com/dnephin/pre-commit-golang>
