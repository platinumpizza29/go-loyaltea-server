FROM --platform=linux/amd64 golang:1.24.3-alpine AS builder

WORKDIR /app

# Install git and certs
RUN apk add --no-cache git ca-certificates

# Download dependencies early
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the Go binary (static)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o loyaltea main.go

# ----------- Runtime Stage -----------
FROM --platform=linux/amd64 alpine:3.19

WORKDIR /app

# Install CA certificates
RUN apk add --no-cache ca-certificates

# Copy built binary
COPY --from=builder /app/loyaltea /app/loyaltea
RUN chmod +x /app/loyaltea

# Optional: copy swagger docs
COPY --from=builder /app/docs ./docs

# Expose Gin port
EXPOSE 8080

# Environment
ENV GIN_MODE=release

# Run the app
CMD ["./loyaltea"]
