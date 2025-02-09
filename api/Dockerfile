# Build stage
FROM golang:1.23.4 AS builder

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /main

# Expose port (adjust as needed)
EXPOSE 8080

# Command to run the application
CMD ["/main"]
