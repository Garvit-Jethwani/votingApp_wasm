FROM golang:alpine as builder
LABEL maintainer="mgdevstack" \
    vendor="Zettabytes" \
    owner="zbio"
ADD main.go  .
COPY common /usr/local/go/src/github.com/roost-io/roost-example/voting_app/service-test-suite/common
COPY ballot /usr/local/go/src/github.com/roost-io/roost-example/voting_app/service-test-suite/ballot
RUN GO111MODULE=off CGO_ENABLED=0 GOFLAGS=-mod=vendor  GOPATH=/usr/local/go/src/ GOARCH=wasm GOOS=js go build -o app

FROM scratch
COPY --from=build /build/hello_world.wasm /hello_world.wasm
ENTRYPOINT [ "hello_world.wasm" ]
