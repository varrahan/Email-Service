# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install git in case dependencies need it
RUN apk add --no-cache git

# Copy go.mod and go.sum first (for caching)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy all source files
COPY . .

# Build the binary
RUN go build -o email-service ./cmd/email-service

# Final stage: minimal image with just the binary
FROM alpine:latest

# Install CA certificates for HTTPS, SMTP, etc.
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/portfolio-backend .

# Run
CMD ["./portfolio-backend"]
