# Stage 1: Builder
FROM golang:1.24-alpine AS builder

# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache git ca-certificates tzdata

WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker cache
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the binary
# -ldflags="-w -s" reduces binary size by removing debug information
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main ./cmd/api/main.go

# Stage 2: Runner
FROM alpine:latest

# Install ca-certificates and tzdata for HTTPS and Timezone support
RUN apk --no-cache add ca-certificates tzdata

# Create a non-root user and group
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/main .

# Copy necessary configs if any (e.g., config.json if not using env vars exclusively, 
# but best practice suggests using env vars. If config.json is verified needed, uncomment below)
# COPY --from=builder /app/config.json . 

# Create directories for uploads with correct permissions
RUN mkdir -p public/uploads/courses && \
    chown -R appuser:appgroup public/uploads

# Switch to non-root user
USER appuser

# Expose port (adjust if necessary)
EXPOSE 3000

# Set timezone (Optional, defaults to UTC, but good for "TimeNowJakarta")
ENV TZ=Asia/Jakarta

# Run the binary
CMD ["./main"]
