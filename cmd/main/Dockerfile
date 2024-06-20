# Use the official Golang image as the base image for building the Go binary
FROM golang:1.22.3-alpine AS builder

# Set the current working directory inside the container
WORKDIR /build

# Copy go.mod and go.sum files first (to leverage Docker cache)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o ./playgobackend ./cmd/main/main.go

# Use a minimal base image for the final container
FROM alpine:latest

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go binary from the builder image
COPY --from=builder /build/playgobackend /app/playgobackend

# Expose the application port
EXPOSE 8080

# Set the default command
ENTRYPOINT ["/app/playgobackend"]
