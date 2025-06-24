# Guidelines for AI Contributors

This repository hosts the Go backend for the OpenCloud server. When submitting
changes via an AI agent, please follow these rules:

## Development workflow

1. **Format Go code**
   - Run `gofmt -w` on any modified Go files.
2. **Lint the project**
   - Execute `make golangci-lint` from the repository root.
3. **Run tests**
   - Execute `make test` to run the Go unit tests.
   - For acceptance tests see `tests/README.md`.
4. **Add a changelog entry**
   - If your change is user facing, create a file in `changelog/unreleased/`
     based on `changelog/TEMPLATE`.

## Pull request message

Summarize what was changed and mention whether `make golangci-lint` and
`make test` succeeded.

For additional contribution guidelines see `CONTRIBUTING.md`.
