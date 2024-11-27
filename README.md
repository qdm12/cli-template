# cli-template

SHORT_DESCRIPTION

![Last release](https://img.shields.io/github/release/qdm12/cli-template?label=Last%20release)
![Last Docker tag](https://img.shields.io/docker/v/qmcgaw/cli-template?sort=semver&label=Last%20Docker%20tag)
![GitHub last release date](https://img.shields.io/github/release-date/qdm12/cli-template?label=Last%20release%20date)
![Commits since release](https://img.shields.io/github/commits-since/qdm12/cli-template/latest?sort=semver)

[![Latest size](https://img.shields.io/docker/image-size/qmcgaw/cli-template/latest?label=Latest%20image)](https://hub.docker.com/r/qmcgaw/cli-template/tags)
[![Last release size](https://img.shields.io/docker/image-size/qmcgaw/cli-template?sort=semver&label=Last%20released%20image)](https://hub.docker.com/r/qmcgaw/cli-template/tags?page=1&ordering=last_updated)

[![GitHub last commit](https://img.shields.io/github/last-commit/qdm12/cli-template.svg)](https://github.com/qdm12/cli-template/commits/main)
[![GitHub commit activity](https://img.shields.io/github/commit-activity/y/qdm12/cli-template.svg)](https://github.com/qdm12/cli-template/graphs/contributors)
[![GitHub closed PRs](https://img.shields.io/github/issues-pr-closed/qdm12/cli-template.svg)](https://github.com/qdm12/cli-template/pulls?q=is%3Apr+is%3Aclosed)
[![GitHub issues](https://img.shields.io/github/issues/qdm12/cli-template.svg)](https://github.com/qdm12/cli-template/issues)
[![GitHub closed issues](https://img.shields.io/github/issues-closed/qdm12/cli-template.svg)](https://github.com/qdm12/cli-template/issues?q=is%3Aissue+is%3Aclosed)

[![MIT](https://img.shields.io/github/license/qdm12/cli-template)](https://github.com/qdm12/cli-template/main/LICENSE)
![Code size](https://img.shields.io/github/languages/code-size/qdm12/cli-template)
![GitHub repo size](https://img.shields.io/github/repo-size/qdm12/cli-template)
![Go version](https://img.shields.io/github/go-mod/go-version/qdm12/cli-template)
![Visitors count](https://visitor-badge.laobi.icu/badge?page_id=cli-template.readme)
[![Build status](https://github.com/qdm12/cli-template/actions/workflows/ci.yml/badge.svg)](https://github.com/qdm12/cli-template/actions/workflows/ci.yml)

[![dockeri.co](https://dockeri.co/image/qmcgaw/cli-template)](https://hub.docker.com/r/qmcgaw/cli-template)

## Quick links

- Problem or suggestion?
  - [Start a discussion](https://github.com/qdm12/cli-template/discussions)
  - [Create an issue](https://github.com/qdm12/cli-template/issues)
- Happy?
  - Sponsor me on [github.com/sponsors/qdm12](https://github.com/sponsors/qdm12)
  - Donate to [paypal.me/qmcgaw](https://www.paypal.me/qmcgaw)
  - Drop me [an email](mailto:quentin.mcgaw@gmail.com)

## Features

- Pre-built standalone binary programs built for Linux, OSX and Windows
- Docker images available on [Docker Hub](https://hub.docker.com/r/qmcgaw/cli-template) and [GitHub Container Registry](https://github.com/users/qmcgaw/packages/container/package/cli-template)
- Docker images are multi-arch: `linux/amd64`, `linux/386`, `linux/arm64`, `linux/arm/v6`, `linux/arm/v7`.

## Usage

You can run the program either as a standalone binary or as a Docker container:

- Standalone binary
  1. Download the binary matching your platform from [the releases page](https://github.com/qdm12/cli-template/releases).
  You can also install the binary from source using [Go](https://golang.org/dl/) with `go install github.com/qdm12/cli-template/cmd/cli-template`
  2. On Linux and OSX, make the binary executable with `chmod +x cli-template`
  3. Run the binary with `./cli-template`
- Docker container

  ```sh
  docker run --rm qmcgaw/cli-template
  ```

  There also Docker tags matching each Github release, for example `:0.1.0` for release `v0.1.0`.
  The latest image tag matches the tip of the main branch.

### Options

| Flag | Environment variable | Default | Description |
|---|---|---|---|
| `--some-path` | `SOME_PATH` | `./path` | Path to some directory |
