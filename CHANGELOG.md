<!-- Source: Keep a Changelog 1.1.0 (MIT) — https://keepachangelog.com/en/1.1.0/ -->
# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

Releases up to and including v0.90.0 predate this file; see the [tags](https://github.com/anyingiit/bytefmt/tags) and [upstream history](https://github.com/cloudfoundry/bytefmt/commits/main) for their changes.

## [Unreleased]

### Added

- Community health files: `CODE_OF_CONDUCT.md`, `SECURITY.md` and this changelog.
- GitHub Actions CI (`gofmt`, `go vet`, `staticcheck`, tests), Dependabot for Go modules and actions, and release-note categories.
- `.editorconfig`, `.gitignore` and an optional `.pre-commit-config.yaml`.
- Runnable examples for `ByteSize`, `ToBytes` and `ToMegabytes`, shown on pkg.go.dev.

### Changed

- Rewrote `README.md` with installation, usage and a table of supported units.
- Moved `CONTRIBUTING.md` to the repository root and added a local development workflow.
- Replaced the issue and pull request templates with standard GitHub issue forms.

### Removed

- Upstream CI sync notice (`.github/TEMPLATE-README.md`).

## [0.90.0] - 2026-09-15

### Changed

- Updated Go module dependencies.

[Unreleased]: https://github.com/anyingiit/bytefmt/compare/v0.90.0...HEAD
[0.90.0]: https://github.com/anyingiit/bytefmt/compare/v0.89.0...v0.90.0
