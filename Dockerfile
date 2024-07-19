# Use an official Golang runtime as the base image
FROM golang:1.22.3

# Set the working directory in the container
WORKDIR /app

# Copy the Go application source code to the working directory
COPY . .

# Cache Go dependencies
COPY go.mod go.sum ./
RUN go mod download

# Build the Go application
RUN go build -o main

# Expose the port on which the Go application will run
EXPOSE 8010

# Start the Go application
CMD ["./main"]
