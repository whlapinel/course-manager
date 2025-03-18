#!/bin/bash
export ENV=development
caddy run --config Caddyfile.dev &

# Use absolute path if needed
/app/cmd/web_app/web_app
