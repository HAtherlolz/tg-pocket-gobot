# Start from a base image with Go installed
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download and cache the Go module dependencies
RUN go mod download

# Copy the source code to the container
COPY . .

# Build the Go application
RUN go build -o app cmd/bot/main.go

# Expose the port that the application listens on
EXPOSE 8080

# Set the entry point for the container
CMD ["./app"]