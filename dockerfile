# Use the official Golang image from Docker Hub as the base image
FROM golang:latest

# Install Vim and Git (you may not need these in production)
RUN apt-get update && \
    apt-get install -y vim git && \
    rm -rf /var/lib/apt/lists/*

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code from the host machine to the container's working directory
COPY . .

# Expose port 5000 to the outside world (you may want to use a different port if needed)
EXPOSE 5000

# Build and run the Go application
CMD ["go", "run", "main.go"]
