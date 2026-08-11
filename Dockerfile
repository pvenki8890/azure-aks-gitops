# Stage 1: Build the Go application
FROM golang:1.26-alpine AS builder

WORKDIR /build/app

COPY app/go.mod .
COPY app/main.go .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /build/azure-web-app .

# Stage 2: Run the application
FROM alpine:latest

WORKDIR /app

COPY --from=builder /build/azure-web-app .

EXPOSE 8081

CMD ["./azure-web-app"]
