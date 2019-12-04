# Default to GO 1.12.13
ARG GO_VERSION=1.12.13
# 第一層基底
FROM golang:${GO_VERSION}-alpine AS build_base
RUN apk add bash --no-cache ca-certificates git gcc g++ libc-dev

WORKDIR /go/src/goformat-v2
COPY . /go/src/goformat-v2
ENV GO111MODULE on

COPY go.mod .
COPY go.sum .

RUN go mod download

FROM build_base AS server_builder

COPY . .

# And compile the project
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go install -a -tags netgo -ldflags '-w -extldflags "-static"' ./cmd/goformat-server

#In this last stage, we start from a fresh Alpine image, to reduce the image size and not ship the Go compiler in our production artifacts.
FROM alpine AS goformat-v2
# We add the certificates to be able to verify remote weaviate instances
RUN apk add ca-certificates
# Finally we copy the statically compiled Go binary.
COPY --from=server_builder /go/bin/goformat-server /bin/goformat
ENTRYPOINT ["/bin/goformat"]
