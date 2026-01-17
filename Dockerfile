# ---------- Build stage ----------
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install git for go modules
RUN apk add --no-cache git

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o app ./cmd/main

# ---------- Runtime stage ----------
FROM alpine:3.19

WORKDIR /app

# Create user AND fix ownership in one layer
RUN addgroup -S appgroup \
 && adduser -S appuser -G appgroup

# Copy binary only

COPY --from=builder /app/app /app/app

# Copy config file
COPY --from=builder /app/etc/app.yaml /app/etc/app.yaml

# Fix permissions (important)
RUN chown -R appuser:appgroup /app /app/etc/app.yaml


USER appuser

# Application port
ENV SERVER_PORT=8000
EXPOSE 8000

ENTRYPOINT ["/app/app"]
