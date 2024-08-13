# Start from the official Golang image for building the binary
FROM golang:1.22.1 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod  file for dependency resolution
COPY go.mod ./

# Copy the rest of the source code
COPY . .

# Create a smaller image for the final executable
FROM debian:bullseye-slim

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage
COPY . .

# Expose the application port
EXPOSE 8080

# Command to run the binary
CMD ["go","run","."]
