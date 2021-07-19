# cli-template

*SHORT_DESCRIPTION*

[![Build status](https://github.com/qdm12/cli-template/actions/workflows/ci.yml/badge.svg)](https://github.com/qdm12/cli-template/actions/workflows/ci.yml)

[![dockeri.co](https://dockeri.co/image/qmcgaw/cli-template)](https://hub.docker.com/r/qmcgaw/cli-template)

![Last release](https://img.shields.io/github/release/qdm12/cli-template?label=Last%20release)
![Last Docker tag](https://img.shields.io/docker/v/qmcgaw/cli-template?sort=semver&label=Last%20Docker%20tag)
[![Last release size](https://img.shields.io/docker/image-size/qmcgaw/cli-template?sort=semver&label=Last%20released%20image)](https://hub.docker.com/r/qmcgaw/cli-template/tags?page=1&ordering=last_updated)
![GitHub last release date](https://img.shields.io/github/release-date/qdm12/cli-template?label=Last%20release%20date)
![Commits since release](https://img.shields.io/github/commits-since/qdm12/cli-template/latest?sort=semver)

[![Latest size](https://img.shields.io/docker/image-size/qmcgaw/cli-template/latest?label=Latest%20image)](https://hub.docker.com/r/qmcgaw/cli-template/tags)

[![GitHub last commit](https://img.shields.io/github/last-commit/qdm12/cli-template.svg)](https://github.com/qdm12/cli-template/commits/main)
[![GitHub commit activity](https://img.shields.io/github/commit-activity/y/qdm12/cli-template.svg)](https://github.com/qdm12/cli-template/graphs/contributors)
[![GitHub closed PRs](https://img.shields.io/github/issues-pr-closed/qdm12/cli-template.svg)](https://github.com/qdm12/cli-template/pulls?q=is%3Apr+is%3Aclosed)
[![GitHub issues](https://img.shields.io/github/issues/qdm12/cli-template.svg)](https://github.com/qdm12/cli-template/issues)
[![GitHub closed issues](https://img.shields.io/github/issues-closed/qdm12/cli-template.svg)](https://github.com/qdm12/cli-template/issues?q=is%3Aissue+is%3Aclosed)

[![Lines of code](https://img.shields.io/tokei/lines/github/qdm12/cli-template)](https://github.com/qdm12/cli-template)
![Code size](https://img.shields.io/github/languages/code-size/qdm12/cli-template)
![GitHub repo size](https://img.shields.io/github/repo-size/qdm12/cli-template)
![Go version](https://img.shields.io/github/go-mod/go-version/qdm12/cli-template)

[![MIT](https://img.shields.io/github/license/qdm12/cli-template)](https://github.com/qdm12/cli-template/master/LICENSE)
![Visitors count](https://visitor-badge.laobi.icu/badge?page_id=cli-template.readme)

## Quick links

- Problem or suggestion?
  - [Start a discussion](https://github.com/qdm12/cli-template/discussions)
  - [Create an issue](https://github.com/qdm12/cli-template/issues)
  - [Check the Wiki](https://github.com/qdm12/cli-template/wiki)
- Happy?
  - Sponsor me on [github.com/sponsors/qdm12](https://github.com/sponsors/qdm12)
  - Donate to [paypal.me/qmcgaw](https://www.paypal.me/qmcgaw)
  - Drop me [an email](mailto:quentin.mcgaw@gmail.com)

## Features

## Usage

```sh

```

### Binary

```sh
VERSION=v0.1.0
ARCH=amd64

wget -O cli-template "https://github.com/qdm12/cli-template/releases/download/$VERSION/cli-template_$VERSION_linux_$ARCH"
chmod 500 cli-template

./cli-template -help
```

### Docker

```sh
docker run -it --rm -v "/yourrepopath:/repository" qmcgaw/cli-template:v0.1.0 -help
```

## Platforms supported

- `linux/amd64`
- `linux/386`
- `linux/arm64`
- `linux/arm/v6`
- `linux/arm/v7`
- `linux/s390x`
- `linux/ppc64le`
- `linux/riscv64`
- If you need one more, please [create an issue](https://github.com/qdm12/cli-template/issues/new)

## Build it yourself

Install Go, then either

- Download it on your machine:

  ```sh
  go get github.com/qdm12/cli-template/cmd/cli-template
  ```

- Clone this repository and build it:

  ```sh
  GOARCH=amd64 go build cmd/cli-template/main.go
  ```
