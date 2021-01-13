ARG ALPINE_VERSION=3.12
ARG GO_VERSION=1.15

FROM --platform=$BUILDPLATFORM golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS base
RUN apk --update add git
ENV CGO_ENABLED=0
WORKDIR /tmp/gobuild
COPY go.mod go.sum ./
RUN go mod download
COPY cmd/ ./cmd/
COPY internal/ ./internal/

FROM --platform=$BUILDPLATFORM base AS test
ENV CGO_ENABLED=1
RUN apk --update add g++
RUN go test -race ./...

FROM --platform=$BUILDPLATFORM base AS lint
ARG GOLANGCI_LINT_VERSION=v1.34.1
RUN wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | \
    sh -s -- -b /usr/local/bin ${GOLANGCI_LINT_VERSION}
COPY .golangci.yml ./
RUN golangci-lint run --timeout=10m

FROM --platform=$BUILDPLATFORM base AS build
COPY --from=qmcgaw/xcputranslate:v0.4.0 /xcputranslate /usr/local/bin/xcputranslate
ARG TARGETPLATFORM
ARG VERSION=unknown
ARG BUILD_DATE="an unknown date"
ARG COMMIT=unknown
COPY cmd/ ./cmd/
COPY internal/ ./internal/
RUN GOARCH="$(xcputranslate -targetplatform ${TARGETPLATFORM} -field arch)" \
    GOARM="$(xcputranslate -targetplatform ${TARGETPLATFORM} -field arm)" \
    go build -trimpath -ldflags="-s -w \
    -X 'main.version=$VERSION' \
    -X 'main.buildDate=$BUILD_DATE' \
    -X 'main.commit=$COMMIT' \
    " -o entrypoint cmd/app/main.go

FROM alpine:${ALPINE_VERSION} AS alpine
RUN apk add ca-certificates

FROM scratch
ARG VERSION=unknown
ARG BUILD_DATE="an unknown date"
ARG COMMIT=unknown
LABEL \
  org.opencontainers.image.authors="quentin.mcgaw@gmail.com" \
  org.opencontainers.image.created=$BUILD_DATE \
  org.opencontainers.image.version=$VERSION \
  org.opencontainers.image.revision=$COMMIT \
  org.opencontainers.image.url="https://github.com/qdm12/REPONAME" \
  org.opencontainers.image.documentation="https://github.com/qdm12/REPONAME" \
  org.opencontainers.image.source="https://github.com/qdm12/REPONAME" \
  org.opencontainers.image.title="REPONAME" \
  org.opencontainers.image.description=""
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
ENTRYPOINT ["/REPONAME"]
COPY --from=build /tmp/gobuild/entrypoint /REPONAME
