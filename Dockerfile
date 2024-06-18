# Start with a base image containing Go
FROM golang:1.16

# Set the working directory inside the container
WORKDIR /app

# Copy the current directory contents into the container
COPY . .

# Download Go modules
RUN go mod tidy

# Build the Go application
RUN go build -o main .

# Expose port 8080
EXPOSE 8080

# Run the executable
CMD ["./main"]
