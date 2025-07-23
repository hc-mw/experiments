# Use the latest Go 1.24 image as the base
FROM golang:1.24 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first to take advantage of Docker caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code
COPY . .

# Ensure the output binary is built for Linux
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o todo-app cmd/main/main.go

# Use a minimal base image for the final container
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Install required packages (use bash for debugging if needed)
RUN apk --no-cache add ca-certificates

# Copy the compiled binary from the builder stage
COPY --from=builder /app/todo-app .

# Ensure the binary is executable
RUN chmod +x ./todo-app

# Set default command to start the server
CMD ["./todo-app", "serve"]
