# Build Stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Install necessary build tools
RUN apk add --no-cache git

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy application source code
COPY . .

# Build the application
RUN go build -o main .

# Final Stage
FROM alpine:latest

WORKDIR /app

# Add a non-root user
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

# Copy the compiled binary from the builder
COPY --from=builder /app/main .

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["/app/main"]
