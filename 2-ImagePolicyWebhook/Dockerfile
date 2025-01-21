# Start with a minimal Go image
FROM golang:1.20 as builder

# Set working directory
WORKDIR /app

# Copy Go modules and source code
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the application
RUN go build -o webhook .

# Final image
FROM debian:bullseye-slim

# Set up working directory
WORKDIR /app

# Copy TLS certificates (ensure they're added at build time)
COPY certs /app/certs

# Copy the built binary
COPY --from=builder /app/webhook /app/webhook

# Expose the webhook port
EXPOSE 8080

# Run the webhook
CMD ["/app/webhook"]

