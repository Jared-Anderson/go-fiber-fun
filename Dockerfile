# Use the official Go image as a base image
FROM golang:latest

# Set environment variables
ENV ENVIRONMENT=Development
ENV DB_CONNECTION_STRING=another_value

# Set the working directory inside the container
WORKDIR /app

# Copy your Go application source code into the container
COPY . .

# Build the Go application
RUN go build -o myapp

# Expose a port if your Go application listens on a specific port
EXPOSE 8080

# Define the command to run your Go application
CMD ["./myapp"]
