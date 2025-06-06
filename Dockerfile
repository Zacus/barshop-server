# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install required build tools
RUN apk add --no-cache git make

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o barshop-server ./cmd/server

# Final stage
FROM alpine:latest

WORKDIR /app

# Install necessary runtime dependencies
RUN apk --no-cache add ca-certificates tzdata

# Copy the binary from builder
COPY --from=builder /app/barshop-server .
COPY --from=builder /app/config.yaml .

# Create necessary directories
RUN mkdir -p /app/logs

# Set environment variables
ENV GIN_MODE=release

EXPOSE 8080

CMD ["./barshop-server"] 