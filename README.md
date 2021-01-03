# REPONAME

## Features

## Usage

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
