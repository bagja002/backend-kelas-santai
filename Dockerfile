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

RUN apk --no-cache add ca-certificates tzdata

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

# Copy binary
COPY --from=builder /app/main .

# Copy config jika memang dipakai
COPY --from=builder /app/config.json .
COPY --from=builder /app/internal/templates/invoice.html ./internal/templates/

# 👉 PENTING: ubah owner /app dulu
RUN chown -R appuser:appgroup /app

# Switch ke non-root
USER appuser

# Sekarang user punya hak penuh ke /app
RUN mkdir -p public/uploads/courses
RUN mkdir -p internal/templates/

EXPOSE 7006
ENV TZ=Asia/Jakarta

CMD ["./main"]




