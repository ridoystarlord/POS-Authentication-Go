# Build stage
FROM golang:1.20-alpine AS builder

# Set environment variables for Go
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Install necessary dependencies
RUN apk --no-cache add git

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project
COPY . .

# Build the application
RUN go build -o pos-authentication .

# Runtime stage
FROM alpine:latest

# Set working directory
WORKDIR /app

# Install runtime dependencies
RUN apk --no-cache add ca-certificates

# Copy the binary from the builder
COPY --from=builder /app/fiber-app .

# Expose application port
EXPOSE 8000

# Start the application
CMD ["./pos-authentication"]
