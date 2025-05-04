# Step 1: Build the Go binary
FROM golang:1.24.2 AS builder

WORKDIR /app

# Copy go.mod and download dependencies first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 go build -o server ./src/main.go

# Step 2: Create a minimal final image
FROM alpine:latest

# Install SSL certificates (for HTTPS and Postgres connections)
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/server .

# Copy configuration files into the container
COPY ./config /root/config
