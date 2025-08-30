FROM --platform=$BUILDPLATFORM golang:1.23-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG TARGETOS TARGETARCH
ENV CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH
RUN go build -ldflags="-s -w" -o /app/main .

FROM --platform=$TARGETPLATFORM scratch
WORKDIR /app
COPY --from=builder /app/main /app/main

EXPOSE 3000
ENTRYPOINT ["/app/main"]
