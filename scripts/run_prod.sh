#!/bin/bash
export ENV=production

# Load environment variables from .env.production
set -a
source .env.production
set +a

# Start Caddy and run web_app
caddy run --config Caddyfile.prod &

# Use absolute path to avoid working directory issues
/app/cmd/web_app/web_app
