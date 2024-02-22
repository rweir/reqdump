# syntax=docker/dockerfile:1

# https://medium.com/@kittipat_1413/optimizing-multi-stage-builds-with-dockerfile-in-golang-a2ee8ed37ec6

# Stage 1: Build stage
FROM golang:1.22-alpine AS build

# Set the working directory
WORKDIR /app

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp .

# Stage 2: Final stage
FROM alpine:edge

# Set the working directory
WORKDIR /app

# Copy the binary from the build stage
COPY --from=build /app/myapp .

# Set the entrypoint command
ENTRYPOINT ["/app/myapp"]

EXPOSE 8080
