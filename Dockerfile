# Use the official Golang image as the base image
FROM golang:alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main ./cmd

# Expose port 50051 to the outside world
EXPOSE 50051

# Command to run the executable
CMD ["./main"]