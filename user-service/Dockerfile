# Use a smaller base image
FROM golang:1.22.2-alpine as builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum to cache dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app as a statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Use a minimal base image for the final stage
FROM alpine:latest

# Set working directory in final image
WORKDIR /root/

# Copy the Go binary from the builder image
COPY --from=builder /app/main .

# Expose port 8080
EXPOSE 8080

# Run the Go app
CMD ["./main"]
