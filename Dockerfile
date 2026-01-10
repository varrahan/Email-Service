# Build stage
FROM golang:1.23.6-alpine AS builder

WORKDIR /app

# Install git in case dependencies need it
RUN apk add --no-cache git

# Copy go.mod and go.sum first for caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy all source files
COPY . .

# Build the binary
RUN go build -o email-service ./cmd/email-service

# Final stage: minimal image with just the binary
FROM alpine:latest

# Install CA certificates for HTTPS and SMTP.
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/email-service /email-service
COPY --from=builder /app/templates ./templates

# Run script
CMD ["/email-service"]
