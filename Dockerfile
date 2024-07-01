# Use the official Golang image from the Docker Hub
FROM golang:1.22.4

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code to the container
COPY . .

# Build the Go app
RUN go build -o app .

# Expose the port on which the app will run
EXPOSE 8080

# Command to run the application
CMD ["./app"]
