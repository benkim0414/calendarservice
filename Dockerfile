FROM golang:1.14.1-alpine as builder
ENV GO111MODULE=on
WORKDIR /module
COPY . /module/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build -a -tags netgo \
    -ldflags '-w -extldflags "-static"' \
    -o server \
    cmd/server/main.go

FROM alpine:latest
RUN apk add --no-cache ca-certificates
RUN GRPC_HEALTH_PROBE_VERSION=v0.3.2 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe
COPY --from=builder /module/server .
EXPOSE 50051
ENTRYPOINT ["/server"]
