# syntax=docker/dockerfile:1.6

FROM --platform=linux/arm64 golang:1.24.3-alpine AS builder
WORKDIR /app

RUN apk add --no-cache git ca-certificates file

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=arm64
RUN go build -ldflags="-s -w" -o /app/main .
RUN file /app/main  # sanity check → tem que sair "ELF 64-bit LSB executable, ARM aarch64"

FROM scratch
WORKDIR /app
COPY --from=builder /app/main /app/main

EXPOSE 3000
ENTRYPOINT ["/app/main"]
