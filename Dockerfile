# Dockerfile 
# for building/running Go (Golang) services
#
# docker build -t go-ginrest:1.20 nicholashoule:go-ginrest
# docker run go-ginrest:1.20
#
# https://docs.docker.com/engine/reference/run/
# https://golang.org/help/
# https://play.golang.org/
#
# Docker multi-stage build
ARG GO_VERSION=1.20
ARG ALPINE_VERSION=latest
# Golang
FROM "golang:${GO_VERSION}" AS builder
# ENV
ENV GOFLAGS=-mod=mod \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

WORKDIR /src
COPY . .

# RUN go build -ldflags "-s -w"
RUN GOFLAGS=-mod=mod go build -ldflags "-s -w" -o /src/bin .

# Alpine
FROM "alpine:${ALPINE_VERSION}"
LABEL service="go-ginrest"
LABEL github="github.com/nicholashoule"
LABEL version="1.0.0"
LABEL description="For running Go (Golang) services."

# RUN apk and add group/user
RUN apk --no-cache add ca-certificates && \
  update-ca-certificates && \
  addgroup -g 1001 appgroup && \
  adduser -H -D -s /bin/false -G appgroup -u 1001 appuser

USER 1001:1001
COPY --from=builder /src/bin/go-ginrest /opt/go/go-ginrest
COPY --from=builder /src/favicon.ico /
ENTRYPOINT ["/opt/go/go-ginrest"]
