# Contribution Guide

Welcome and thank you for considering contributing to gorse-go!

Reading and following these guidelines will help make the contribution process easy and effective for everyone. It also communicates that you agree to respect the time of the developers managing and developing this open-source project. In return, we will address your issues, assess changes, and help you finalize your pull requests.

## Code of Conduct

Please be kind and respectful to others. By participating, you agree to uphold a welcoming, inclusive environment. If you experience or witness unacceptable behavior, please open an issue or contact the maintainers.

## Your First Contribution

We welcome all contributions, including bug reports, documentation improvements, tests, and new features. If you are unsure where to start:

- Look for issues labeled `good first issue` or `help wanted`.
- Open a discussion if you want feedback on an idea before implementing.
- For non-trivial changes, propose your approach in an issue first to align early.

### Contribution Workflow

To contribute to the gorse-go code base, please follow this workflow:

- Fork the repository to your own GitHub account.
- Create a topic branch from `main` in your fork.
- Make commits; add tests if the change fixes a bug or adds functionality.
- Run tests locally and ensure they pass.
- Push your branch to your fork.
- Submit a pull request (PR) to `gorse-io/gorse-go:main`.

We appreciate your contributions!

## Getting Started

### Setup Development Environment

The following installations are required:

- **Go 1.20+**: Minimum Go version to build and test gorse-go.
- **Docker Compose**: For starting a local Gorse cluster when you need to test against a live server.

After installing required software, run the following commands to verify the setup:

```bash
# Verify Go is installed
go version

# Initialize or tidy modules
go mod tidy

# Start a local Gorse cluster for integration testing
curl -sL https://github.com/gorse-io/gorse/raw/refs/heads/master/client/setup-test.sh | bash
```

### Run Unit Tests

```bash
go test ./...
```

If tests are flaky or depend on a running Gorse service, please mark them clearly and prefer deterministic tests using mocks when possible.

## Project Structure

This repository is a small Go module. Key files include:

- `client.go`: Client implementation for interacting with Gorse.
- `client_test.go`: Unit tests for the client.
- `model.go`: Data models and types used by the client.
- `README.md`: Usage and documentation.
- `go.mod`: Module definition and dependencies.

## Development Guidelines

- **Code Style**: Follow standard Go conventions (`gofmt`/`go fmt`) and idiomatic Go patterns.
- **Linting**: Prefer `go vet` and consider `golangci-lint` if available locally. Keep changes minimal and focused.
- **Testing**:
  - Add unit tests for new features or bug fixes.
  - Aim for clear, deterministic tests. Avoid external dependencies unless necessary.
- **Documentation**:
  - Update `README.md` if public behavior changes or new features are added.
  - Add inline comments only when they add clarity beyond self-explanatory code.
- **Dependencies**: Keep dependencies minimal. Update `go.mod` and `go.sum` as needed via `go mod tidy`.

## Pull Request Checklist

Before opening a PR:

- Tests pass locally: `go test ./...`.
- Code is formatted: `go fmt ./...`.
- Static checks pass: `go vet ./...`.
- Public API changes are documented in `README.md` and commit message.
- Commit messages are clear and descriptive.

## Issues and Bug Reports

When reporting a bug, please include:

- Steps to reproduce and the expected vs. actual behavior.
- Go version (`go version`) and OS information.
- Relevant logs, stack traces, or minimal code samples.

## Release and Versioning

This project follows semantic versioning where possible. Breaking changes should be clearly communicated and discussed before merging.

## Getting Help

- Join us on Discord: https://discord.gg/x6gAtNNkAE (`#developers` channel).
- Open an issue for questions, proposals, or support.

Thank you for helping improve gorse-go!