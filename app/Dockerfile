# syntax=docker/dockerfile:1.6

FROM --platform=$BUILDPLATFORM golang:1.25.0-alpine AS builder
WORKDIR /app

RUN apk add --no-cache git ca-certificates file

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG TARGETOS=linux
ARG TARGETARCH
ENV CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH
RUN go build -ldflags="-s -w" -o /app/main ./cmd/api
RUN file /app/main

FROM scratch
WORKDIR /app
COPY --from=builder /app/main /app/main

EXPOSE 3000
ENTRYPOINT ["/app/main"]
