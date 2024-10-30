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
    # -- node -- #
    - id: yarn-prettier
    # -- rust -- #
    - id: rust-fmt
    - id: cargo-check
    - id: cargo-clippy
```

## Refs

- <https://github.com/dnephin/pre-commit-golang>
