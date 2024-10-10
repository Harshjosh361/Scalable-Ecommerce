# First stage: Build the Go binary
FROM golang:1.22.2-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum to cache dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy the rest of the application files
COPY . .

# Build the Go app as a statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Second stage: Minimal runtime image
FROM alpine:latest

WORKDIR /root/

# Copy the Go binary from the builder stage
COPY --from=builder /app/main .

# Expose port 8083
EXPOSE 8083

# Run the Go binary
CMD ["./main"]
