<!-- Source: Best-README-Template BLANK_README (Unlicense) — https://github.com/othneildrew/Best-README-Template -->
<a id="readme-top"></a>

# bytefmt

A tiny, dependency-free Go package for converting between byte counts and human-readable strings such as `100.5M` or `1.5GiB`.

[![CI](https://github.com/anyingiit/bytefmt/actions/workflows/ci.yml/badge.svg)](https://github.com/anyingiit/bytefmt/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/code.cloudfoundry.org/bytefmt.svg)](https://pkg.go.dev/code.cloudfoundry.org/bytefmt)
[![Go Report Card](https://goreportcard.com/badge/code.cloudfoundry.org/bytefmt)](https://goreportcard.com/report/code.cloudfoundry.org/bytefmt)
[![License](https://img.shields.io/github/license/anyingiit/bytefmt)](LICENSE)

[Report a bug](https://github.com/anyingiit/bytefmt/issues/new?template=bug_report.yml) · [Request a feature](https://github.com/anyingiit/bytefmt/issues/new?template=feature_request.yml)

<details>
  <summary>Table of Contents</summary>
  <ol>
    <li><a href="#about-the-project">About The Project</a></li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li>
      <a href="#usage">Usage</a>
      <ul>
        <li><a href="#formatting-bytes">Formatting bytes</a></li>
        <li><a href="#parsing-strings">Parsing strings</a></li>
        <li><a href="#supported-units">Supported units</a></li>
      </ul>
    </li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>

## About The Project

`bytefmt` formats a `uint64` byte count as a short human-readable string, and parses such strings back into bytes. It uses only the Go standard library and was originally written for Cloud Foundry.

- `ByteSize` picks the largest unit that keeps the value at or above 1 and prints at most one decimal place.
- `ToBytes` and `ToMegabytes` accept `K`, `KB` and `KiB` style suffixes, case-insensitively, and ignore surrounding whitespace.
- All units are **binary** (base 2): `1K == 1KB == 1KiB == 1024` bytes.

> [!NOTE]
> This repository is a fork of [cloudfoundry/bytefmt](https://github.com/cloudfoundry/bytefmt). The Go module path is unchanged, so import it as `code.cloudfoundry.org/bytefmt`.

See the [open issues](https://github.com/anyingiit/bytefmt/issues) for planned features and known issues, and the [changelog](CHANGELOG.md) for what changed in each release.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Getting Started

### Prerequisites

- Go 1.26 or later (see [`go.mod`](go.mod))

### Installation

```sh
go get code.cloudfoundry.org/bytefmt
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Usage

### Formatting bytes

```go
package main

import (
	"fmt"

	"code.cloudfoundry.org/bytefmt"
)

func main() {
	fmt.Println(bytefmt.ByteSize(100.5 * bytefmt.MEGABYTE)) // 100.5M
	fmt.Println(bytefmt.ByteSize(uint64(1024)))             // 1K
	fmt.Println(bytefmt.ByteSize(0))                        // 0B
}
```

### Parsing strings

```go
b, err := bytefmt.ToBytes("1.5GiB")   // 1610612736, nil
m, err := bytefmt.ToMegabytes("2g")   // 2048, nil
_, err = bytefmt.ToBytes("12")        // error: a unit is required
```

`ToBytes` returns an error when the unit is missing or unknown, or when the number is negative or malformed.

### Supported units

| Constant | Value | `ByteSize` suffix | Accepted by `ToBytes` |
| --- | --- | --- | --- |
| `BYTE` | 1 | `B` | `B` |
| `KILOBYTE` | 1024 | `K` | `K`, `KB`, `KiB` |
| `MEGABYTE` | 1024² | `M` | `M`, `MB`, `MiB` |
| `GIGABYTE` | 1024³ | `G` | `G`, `GB`, `GiB` |
| `TERABYTE` | 1024⁴ | `T` | `T`, `TB`, `TiB` |
| `PETABYTE` | 1024⁵ | `P` | `P`, `PB`, `PiB` |
| `EXABYTE` | 1024⁶ | `E` | `E`, `EB`, `EiB` |

The full API reference is on [pkg.go.dev](https://pkg.go.dev/code.cloudfoundry.org/bytefmt).

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Contributing

Contributions are welcome. Read [CONTRIBUTING.md](CONTRIBUTING.md) for how to open an issue or a pull request, and [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) for the standards expected of everyone taking part.

Please do not report security issues in public issues or pull requests. [SECURITY.md](SECURITY.md) explains how to report them privately.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## License

Distributed under the Apache License 2.0. See [LICENSE](LICENSE) and [NOTICE](NOTICE) for details.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Contact

Project link: [https://github.com/anyingiit/bytefmt](https://github.com/anyingiit/bytefmt)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Acknowledgments

- [Cloud Foundry](https://www.cloudfoundry.org/) and the [App Runtime Platform working group](https://github.com/cloudfoundry/community/blob/main/toc/working-groups/app-runtime-platform.md), who created and maintain the upstream [cloudfoundry/bytefmt](https://github.com/cloudfoundry/bytefmt).
- [Chef's Pick OSS Starter](https://github.com/chefs-pick-oss-starter/chefs-pick-oss-starter), whose community-standard picks shaped this repository's documentation and settings.

<p align="right">(<a href="#readme-top">back to top</a>)</p>
