# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Use Go 1.23 bookworm as base image
FROM golang:1.24.2 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Get the language/compact package
# RUN go get golang.org/x/text/internal/language/compact
RUN go mod vendor

# Build the Go app
RUN go build -o main .

# Step 5: Use a minimal image to run the app
FROM ubuntu:22.04

# Set DEBIAN_FRONTEND to non-interactive to avoid interactive prompts
ARG R_ADDRS
ARG R_USERNAME
ARG R_PASSWORD

ENV ROUTEROS_ADDRESS=$R_ADDRS
ENV ROUTEROS_USERNAME=$R_USERNAME
ENV ROUTEROS_PASSWORD=$R_PASSWORD
ENV DEBIAN_FRONTEND=noninteractive

# Install Node.js and other necessary packages
RUN apt-get update && \
    apt-get install -y curl gnupg2 lsb-release && \
    curl -sSL https://deb.nodesource.com/setup_16.x | bash - && \
    apt-get install -y nodejs

# Step 6: Set the working directory
WORKDIR /

# Step 7: Copy the compiled binary from the builder image
COPY --from=builder /app/main ./bin/

# Step 8: Expose the port (optional)
EXPOSE 38000

# Run the executable
CMD ["./bin/main"]