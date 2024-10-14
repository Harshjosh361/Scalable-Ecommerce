# First stage: Build the Go binary
FROM golang:1.22.2-alpine AS builder


WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# Build the Go app as a statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Second stage: Create a minimal runtime image
FROM alpine:latest

# Set working directory in the final image
WORKDIR /root/

# Copy the Go binary from the builder stage
COPY --from=builder /app/main .

EXPOSE 8082

CMD ["./main"]
