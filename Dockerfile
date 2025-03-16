# Use Alpine as the base image
FROM alpine:latest

# Install dependencies
RUN apk add --no-cache go caddy sqlite

# Set working directory
WORKDIR /app

# Copy Go app files
COPY ./build .

# Start Caddy and the Go app
CMD caddy start && ./app
