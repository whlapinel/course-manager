# Stage 1: Build Caddy and App
FROM golang:1.24 as builder

# Install xcaddy and build Caddy
RUN go install github.com/caddyserver/xcaddy/cmd/xcaddy@latest && \
    xcaddy build --with github.com/caddy-dns/digitalocean --output /go/bin/caddy

# Set working directory for app build
WORKDIR /app

# ✅ Step 1: Copy dependency files first to cache go.mod and go.sum
COPY go.mod go.sum ./
RUN go mod download

# ✅ Step 2: Copy the rest of the source code separately
COPY cmd cmd
COPY internal internal

# ✅ Step 3: Build the app binary
RUN go build -o /app/cmd/web_app/web_app ./cmd/web_app

# Stage 2: Final Image
FROM ubuntu:latest

# Install runtime dependencies
RUN apt update && apt install -y wget tar curl ca-certificates p11-kit

# Install Marp
RUN wget -O marp.tar.gz https://github.com/marp-team/marp-cli/releases/download/v4.1.2/marp-cli-v4.1.2-linux.tar.gz \
    && tar -xzf marp.tar.gz -C /usr/local/bin/ \
    && chmod +x /usr/local/bin/marp 

# ✅ Copy Caddy and app from builder stage
COPY --from=builder /go/bin/caddy /usr/bin/caddy
COPY --from=builder /app/cmd/web_app/web_app /app/cmd/web_app/web_app

# Set working directory
WORKDIR /app

# ✅ Copy config and scripts AFTER the build stage to avoid cache invalidation
COPY ./Caddyfile.dev . 
COPY ./Caddyfile.prod . 
COPY ./scripts/run.sh .

# Set permissions for scripts
RUN chmod +x ./run.sh /app/cmd/web_app/web_app

# Expose HTTPS port
EXPOSE 443

# Start the app
CMD ["./run.sh"]
