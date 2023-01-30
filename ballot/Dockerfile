FROM golang:alpine as builder
LABEL maintainer="mgdevstack" \
    vendor="Zettabytes" \
    owner="zbio"
ADD main.go  .
RUN GO111MODULE=off CGO_ENABLED=0 GOOS=js GOARCH=wasm go build -o app.wasm

FROM scratch
LABEL maintainer="mgdevstack" \
    vendor="Zettabytes" \
    owner="zbio"
COPY --from=builder /go/app.wasm /app.wasm
USER nobody
ENTRYPOINT ["/app.wasm"]
