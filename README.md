# REPONAME

*SHORT_DESCRIPTION*

[![Build status](https://github.com/qdm12/REPONAME/actions/workflows/ci.yml/badge.svg)](https://github.com/qdm12/REPONAME/actions/workflows/ci.yml)

[![dockeri.co](https://dockeri.co/image/qmcgaw/REPONAME)](https://hub.docker.com/r/qmcgaw/REPONAME)

![Last release](https://img.shields.io/github/release/qdm12/REPONAME?label=Last%20release)
![Last Docker tag](https://img.shields.io/docker/v/qmcgaw/REPONAME?sort=semver&label=Last%20Docker%20tag)
[![Last release size](https://img.shields.io/docker/image-size/qmcgaw/REPONAME?sort=semver&label=Last%20released%20image)](https://hub.docker.com/r/qmcgaw/REPONAME/tags?page=1&ordering=last_updated)
![GitHub last release date](https://img.shields.io/github/release-date/qdm12/REPONAME?label=Last%20release%20date)
![Commits since release](https://img.shields.io/github/commits-since/qdm12/REPONAME/latest?sort=semver)

[![Latest size](https://img.shields.io/docker/image-size/qmcgaw/REPONAME/latest?label=Latest%20image)](https://hub.docker.com/r/qmcgaw/REPONAME/tags)

[![GitHub last commit](https://img.shields.io/github/last-commit/qdm12/REPONAME.svg)](https://github.com/qdm12/REPONAME/commits/main)
[![GitHub commit activity](https://img.shields.io/github/commit-activity/y/qdm12/REPONAME.svg)](https://github.com/qdm12/REPONAME/graphs/contributors)
[![GitHub closed PRs](https://img.shields.io/github/issues-pr-closed/qdm12/REPONAME.svg)](https://github.com/qdm12/REPONAME/pulls?q=is%3Apr+is%3Aclosed)
[![GitHub issues](https://img.shields.io/github/issues/qdm12/REPONAME.svg)](https://github.com/qdm12/REPONAME/issues)
[![GitHub closed issues](https://img.shields.io/github/issues-closed/qdm12/REPONAME.svg)](https://github.com/qdm12/REPONAME/issues?q=is%3Aissue+is%3Aclosed)

[![Lines of code](https://img.shields.io/tokei/lines/github/qdm12/REPONAME)](https://github.com/qdm12/REPONAME)
![Code size](https://img.shields.io/github/languages/code-size/qdm12/REPONAME)
![GitHub repo size](https://img.shields.io/github/repo-size/qdm12/REPONAME)
![Go version](https://img.shields.io/github/go-mod/go-version/qdm12/REPONAME)

[![MIT](https://img.shields.io/github/license/qdm12/REPONAME)](https://github.com/qdm12/REPONAME/master/LICENSE)
![Visitors count](https://visitor-badge.laobi.icu/badge?page_id=REPONAME.readme)

## Quick links

- Problem or suggestion?
  - [Start a discussion](https://github.com/qdm12/REPONAME/discussions)
  - [Create an issue](https://github.com/qdm12/REPONAME/issues)
  - [Check the Wiki](https://github.com/qdm12/REPONAME/wiki)
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

wget -O REPONAME "https://github.com/qdm12/REPONAME/releases/download/$VERSION/REPONAME_$VERSION_linux_$ARCH"
chmod 500 REPONAME

./REPONAME -help
```

### Docker

```sh
docker run -it --rm -v "/yourrepopath:/repository" qmcgaw/REPONAME:v0.1.0 -help
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
- If you need one more, please [create an issue](https://github.com/qdm12/REPONAME/issues/new)

## Build it yourself

Install Go, then either

- Download it on your machine:

  ```sh
  go get github.com/qdm12/REPONAME/cmd/REPONAME
  ```

- Clone this repository and build it:

  ```sh
  GOARCH=amd64 go build cmd/REPONAME/main.go
  ```
