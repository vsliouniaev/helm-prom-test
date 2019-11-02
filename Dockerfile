FROM golang:alpine as build

# This is expected to be overridden by looper
ARG VERSION=0.0.0
ARG PACKAGE="github.com/vsliouniaev/helm-test-image"

WORKDIR /go/src/${PACKAGE}

COPY . .
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GO111MODULE=off

RUN test -z $(go fmt ./... 2>&1)
RUN go test ./...
RUN go build -a -o helm-prom-test -ldflags \
    "-X ${PACKAGE}/core.Version=${VERSION} -X ${PACKAGE}/core.BuildTime=$(date -u +%FT%TZ)"
RUN mv helm-prom-test /helm-prom-test

FROM gcr.io/distroless/static:latest
WORKDIR /
COPY --from=build /helm-prom-test .
ENTRYPOINT ["/helm-prom-test"]
