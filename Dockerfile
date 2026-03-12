# Build stage
FROM golang:1.26-alpine AS builder

WORKDIR /app

# Copy dependency files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/app ./cmd/app/main.go

# Final stage
FROM alpine:latest

WORKDIR /bin

# Copy binary from builder
COPY --from=builder /bin/app .

# Expose ports
EXPOSE 8081 50051

# Run the binary
CMD ["./app"]
