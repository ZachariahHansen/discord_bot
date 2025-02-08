# Build stage
FROM golang:1.21-alpine AS builder

# Install git for private dependencies if needed
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./main.go

# Final stage
FROM alpine:latest

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/main .

# Expose port (adjust as needed)
EXPOSE 8080

# Command to run the application
CMD ["./main"]