# Use the official Golang image as the base image.
FROM golang:1.18-alpine AS build

# Set the working directory in the container.
WORKDIR /app

# Copy the current directory contents into the container at /app.
COPY . .

# Download and install dependencies.
RUN go mod download

# Build the binary.
RUN go build -o gobot .

# Use a smaller image to deploy the application.
FROM alpine:latest

# Set the working directory in the container.
WORKDIR /app

# Copy the binary from the build stage.
COPY --from=build /app/gobot .

# Run the binary.
CMD ["/app/gobot"]