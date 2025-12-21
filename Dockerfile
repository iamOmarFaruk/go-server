# Multi-stage build for Go web application
FROM golang:1.21-alpine AS builder

# Install git (for go modules)
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy go mod files and download dependencies
COPY go.mod ./
# Copy go.sum if it exists (optional)
COPY go.su[m] ./
RUN go mod download

# Copy source code
COPY . .

# Generate go.sum and build the application
RUN go mod tidy && CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/server

# Final stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

# Create non-root user
RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

WORKDIR /app

# Copy the binary from builder stage
COPY --from=builder /app/main .

# Copy templates and static files
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static

# Change ownership of the app directory
RUN chown -R appuser:appgroup /app

# Switch to non-root user
USER appuser

# Expose port
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

# Run the binary
CMD ["./main"]