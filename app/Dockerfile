FROM golang:1.24.3-alpine AS builder    
WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o main 

FROM scratch
WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 3000

CMD ["/app/main"]
