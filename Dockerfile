# Dockerfile for building/running Go (Golang) services
#
# https://docs.docker.com/engine/reference/run/
# https://golang.org/help/
# https://play.golang.org/
#
# multi-stage, builder
FROM golang:1.17 AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /src
COPY . .

RUN go build \
  -ldflags "-s -w -extldflags 'static'" \
  -installsuffix cgo \
  -tags netgo \
  -mod vendor \
  -o /bin/ginrest \
  .

# final
FROM alpine:latest
RUN apk --no-cache add ca-certificates && \
  update-ca-certificates && \
  addgroup -g 1001 appgroup && \
  adduser -H -D -s /bin/false -G appgroup -u 1001 appuser

USER 1001:1001
COPY --from=builder /bin/go-ginrest /bin/go-ginrest
ENTRYPOINT ["/bin/go-ginrest"]
