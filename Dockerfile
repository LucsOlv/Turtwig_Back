# Build stage
FROM golang:1.24-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git make

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Install Wire and Swag
RUN go install github.com/google/wire/cmd/wire@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Copy the source code
COPY . .

# Generate Wire and Swagger files
RUN cd internal/app && wire
RUN swag init

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/api

# Final stage
FROM alpine:latest

# Add ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/bin/api .
COPY --from=builder /app/docs ./docs

# Expose port
EXPOSE 8080

# Command to run
CMD ["./api"]
