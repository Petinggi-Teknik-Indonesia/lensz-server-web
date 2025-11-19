# ---------------------------
# Stage 1 — Build
# ---------------------------
FROM golang:1.25-alpine AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

# Install dependencies required for go build
RUN apk add --no-cache git

# Copy mod files first for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the backend source
COPY . .

# Build the Go binary
RUN go build -o server ./cmd || go build -o server .


# ---------------------------
# Stage 2 — Runtime
# ---------------------------
FROM alpine:3.19

# Required for PostgreSQL TLS connections
RUN apk add --no-cache ca-certificates

WORKDIR /app

# Copy the compiled binary
COPY --from=builder /app/server .

# Optional: copy .env (keep only if you really want env from file)
# COPY .env .env

# Your app exposes whatever port you set via SERVER_PORT
EXPOSE 8080


CMD ["./server"]
