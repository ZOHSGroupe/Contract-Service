# Use the official Golang image
FROM golang:1.18

# Set the working directory in the container
WORKDIR /go/src/app

# Copy the entire project to the working directory
COPY . .

# Expose the port the application runs on
EXPOSE $GO_DOCKER_PORT

# Build the Go application
RUN go build -o main main.go

# Command to run the executable
CMD ["./main"]