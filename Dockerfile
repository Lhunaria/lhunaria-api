FROM golang:1.23-alpine AS builder
# Set the Current Working Directory inside the container
WORKDIR /app
# Copy go mod and sum files
COPY go.mod ./
# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
# Copy the source code into the container
COPY . .
# Build the Go app
RUN go build -o lhunaria-api

# ------------------------------------------------------#

FROM alpine:latest
# Set the Current Working Directory inside the container
WORKDIR /app
# Copy the Pre-built binary file from the builder stage
COPY --from=builder /app/lhunaria-api .
# Expose port 8080 to the outside world
EXPOSE 8080
# Command to run the executable
CMD ["./lhunaria-api"]
