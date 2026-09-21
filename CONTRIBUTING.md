<!-- Source: GitHub Open Source Guides (CC BY 4.0, structure only) — https://opensource.guide/starting-a-project/; GitHub Docs (official) — https://docs.github.com/en/communities/setting-up-your-project-for-healthy-contributions/setting-guidelines-for-repository-contributors -->
# Contributing

Thanks for your interest in bytefmt. Every kind of contribution is welcome, from a one-line typo fix to a new feature.

## Code of Conduct

This project is governed by its [Code of Conduct](CODE_OF_CONDUCT.md). By taking part in this project you agree to abide by its terms.

## Ways to contribute

- Report a bug you ran into.
- Suggest a feature or an improvement.
- Improve the documentation, including examples and typo fixes.
- Contribute code through a pull request.

## Reporting bugs

Search the existing issues first, in case someone has already reported the same problem. If nothing matches, open a new issue from the Issues tab (Issues, then New issue) and pick the **Bug report** form. Please fill in every required field: a report that says what you expected, what happened instead, and how to reproduce it (ideally the exact input passed to `ByteSize`, `ToBytes` or `ToMegabytes`) is usually fixed much faster.

## Suggesting features

Open a new issue from the Issues tab and pick the **Feature request** form. Describe the problem you want to solve before the solution you have in mind, so that other ways of solving it can be considered as well. bytefmt keeps a deliberately small, dependency-free API, so changes to exported behavior need a clear use case.

## Reporting security issues

Do not report security issues in public. Follow the steps in the [security policy](SECURITY.md) instead, so that the problem can be fixed before it becomes widely known.

## Submitting pull requests

1. Fork the repository and create a branch from `main`.
2. Follow the existing code style: run `gofmt`, `go vet` and `staticcheck` (configured in [`staticcheck.conf`](staticcheck.conf)).
3. Add or update tests that cover your change in `bytes_test.go`.
4. If the change is worth recording, add an entry under `Unreleased` in the [changelog](CHANGELOG.md).
5. Make sure the CI checks pass.
6. Open a pull request and fill in the template.

Smaller, focused pull requests are easier to review and get merged sooner than large ones. If you plan a bigger change, it is worth opening an issue first to agree on the approach.

> [!NOTE]
> This repository is a fork of [cloudfoundry/bytefmt](https://github.com/cloudfoundry/bytefmt). Changes that should reach every user of `code.cloudfoundry.org/bytefmt` belong upstream; contributions there require the Cloud Foundry Foundation CLA for [individuals](https://www.cloudfoundry.org/wp-content/uploads/2015/07/CFF_Individual_CLA.pdf) or [corporations](https://www.cloudfoundry.org/wp-content/uploads/2015/07/CFF_Corporate_CLA.pdf).

## Development setup

You need Go 1.26 or later. Dependencies are vendored in [`vendor/`](vendor), so no network access is needed to build or test.

```sh
git clone https://github.com/anyingiit/bytefmt.git
cd bytefmt
go test ./...                 # run the Ginkgo/Gomega suite
go test -bench . -run '^$'    # run the benchmarks
go vet ./...
go run honnef.co/go/tools/cmd/staticcheck@latest ./...
```

After changing dependencies, run `go mod tidy && go mod vendor` and commit both `go.mod`/`go.sum` and `vendor/`.

Optionally, install [pre-commit](https://pre-commit.com) and run `pre-commit install` once to run the checks in [`.pre-commit-config.yaml`](.pre-commit-config.yaml) before each commit. [`.editorconfig`](.editorconfig) keeps whitespace consistent across editors.

### Running tests in the Cloud Foundry CI container

The [`scripts/`](scripts) directory contains the upstream working group's Docker-based test harness. It expects [wg-app-platform-runtime-ci](https://github.com/cloudfoundry/wg-app-platform-runtime-ci) to be cloned next to this repository:

- `./scripts/create-docker-container.bash`: create a Docker container with the right mounts.
- `./scripts/test-in-docker.bash`: create the container and run all tests in one step.
- Inside the container, `/repo/scripts/docker/test.bash` runs all tests.

## Questions

For questions about using the project, start a thread in the [Discussions](https://github.com/anyingiit/bytefmt/discussions) tab instead of opening an issue. Issues are reserved for bug reports and feature requests.
