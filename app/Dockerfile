# syntax=docker/dockerfile:1.6

FROM --platform=linux/amd64 golang:1.24.3-alpine AS builder
WORKDIR /app

RUN apk add --no-cache git ca-certificates

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG TARGETOS TARGETARCH
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o /app/main .

FROM scratch
WORKDIR /app
COPY --from=builder /app/main /app/main

EXPOSE 3000
ENTRYPOINT ["/app/main"]
