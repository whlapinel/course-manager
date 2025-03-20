#!/bin/bash
docker pull whlapinel/course-manager:latest
docker compose -f compose.yaml -f compose.prod.yaml up