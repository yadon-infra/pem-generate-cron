# Build stage
FROM golang:1.20-alpine as builder

WORKDIR /app

COPY . .

RUN go build -o rsa-key-generator main.go

# Final stage
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/rsa-key-generator .

CMD ["./rsa-key-generator", "-dir", "/keys"]
