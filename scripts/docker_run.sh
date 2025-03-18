#!/bin/bash
docker run \
  -v $(pwd)/.env.development:/app/.env.development \
  -v $(pwd)/.env.production:/app/.env.production \
  -v $(pwd)/internal/data/database/users/:/app/internal/data/database/users/ \
  -v $(pwd)/sites/assets/:/app/sites/assets/ \
  -v $(pwd)/internal/data/database/migrations/:/app/internal/data/database/migrations/ \
  -v $(pwd)/internal/data/database/course_manager.db:/app/internal/data/database/course_manager.db \
  -p 443:443 \
  -it \
  whlapinel/course-manager