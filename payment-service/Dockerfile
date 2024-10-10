# First stage: Build the Go binary
FROM golang:1.22.2-alpine AS builder

# Set the working directory inside the builder stage
WORKDIR /app

# Copy go.mod and go.sum to cache dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy the rest of the application files
COPY . .

# Build the Go app as a statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Second stage: Create a minimal runtime image
FROM alpine:latest

# Set working directory in the final image
WORKDIR /root/

# Copy the Go binary from the builder stage
COPY --from=builder /app/main .

# Expose port 8082
EXPOSE 8082

# Run the Go binary
CMD ["./main"]
