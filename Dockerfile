ARG ALPINE_VERSION=3.13
ARG GO_VERSION=1.16
ARG XCPUTRANSLATE_VERSION=v0.6.0
ARG GOLANGCI_LINT_VERSION=v1.41.1

FROM --platform=${BUILDPLATFORM} qmcgaw/xcputranslate:${XCPUTRANSLATE_VERSION} AS xcputranslate
FROM --platform=${BUILDPLATFORM} qmcgaw/binpot:golangci-lint-${GOLANGCI_LINT_VERSION} AS golangci-lint

FROM --platform=${BUILDPLATFORM} golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS base
ENV CGO_ENABLED=0
RUN apk --update add git g++
WORKDIR /tmp/gobuild
COPY --from=xcputranslate /xcputranslate /usr/local/bin/xcputranslate
COPY --from=golangci-lint /bin /go/bin/golangci-lint
# Copy repository code and install Go dependencies
COPY go.mod go.sum ./
RUN go mod download
COPY cmd/ ./cmd/
COPY internal/ ./internal/

FROM --platform=${BUILDPLATFORM} base AS test
ENV CGO_ENABLED=1

FROM --platform=${BUILDPLATFORM} base AS lint
COPY .golangci.yml ./
RUN golangci-lint run --timeout=10m

FROM --platform=${BUILDPLATFORM} base AS tidy
RUN git init && \
  git config user.email ci@localhost && \
  git config user.name ci && \
  git add -A && git commit -m ci && \
  sed -i '/\/\/ indirect/d' go.mod && \
  go mod tidy && \
  git diff --exit-code -- go.mod

FROM --platform=${BUILDPLATFORM} base AS build
ARG TARGETPLATFORM
ARG VERSION=unknown
ARG CREATED="an unknown date"
ARG COMMIT=unknown
RUN GOARCH="$(xcputranslate translate -targetplatform ${TARGETPLATFORM} -field arch)" \
  GOARM="$(xcputranslate translate -targetplatform ${TARGETPLATFORM} -field arm)" \
  go build -trimpath -ldflags="-s -w \
  -X 'main.version=$VERSION' \
  -X 'main.buildDate=$CREATED' \
  -X 'main.commit=$COMMIT' \
  " -o entrypoint cmd/app/main.go

FROM alpine:${ALPINE_VERSION} AS alpine
RUN apk add ca-certificates

FROM scratch
ARG VERSION=unknown
ARG CREATED="an unknown date"
ARG COMMIT=unknown
LABEL \
  org.opencontainers.image.authors="quentin.mcgaw@gmail.com" \
  org.opencontainers.image.created=$CREATED \
  org.opencontainers.image.version=$VERSION \
  org.opencontainers.image.revision=$COMMIT \
  org.opencontainers.image.url="https://github.com/qdm12/cli-template" \
  org.opencontainers.image.documentation="https://github.com/qdm12/cli-template" \
  org.opencontainers.image.source="https://github.com/qdm12/cli-template" \
  org.opencontainers.image.title="cli-template" \
  org.opencontainers.image.description=""
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
ENTRYPOINT ["/cli-template"]
USER 1000
COPY --from=build --chown=1000 /tmp/gobuild/entrypoint /cli-template
