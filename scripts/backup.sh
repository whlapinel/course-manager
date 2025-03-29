#!/bin/bash

# === CONFIG ===
DB_PATH="internal/data/database/course_manager.db"
USER_DIR="internal/data/users"
BACKUP_DIR="backup"
TIMESTAMP=$(date +"%Y-%m-%d_%H-%M-%S")
ARCHIVE_NAME="backup_$TIMESTAMP.tar.gz"

# === SAFETY CHECKS ===
if [ ! -f "$DB_PATH" ]; then
  echo "ERROR: Database file not found at $DB_PATH"
  exit 1
fi

if [ ! -d "$USER_DIR" ]; then
  echo "ERROR: User directory not found at $USER_DIR"
  exit 1
fi

mkdir -p "$BACKUP_DIR"

# === BACKUP ===
tar -czf "$BACKUP_DIR/$ARCHIVE_NAME" "$DB_PATH" "$USER_DIR"

# === DONE ===
echo "✅ Backup complete: $BACKUP_DIR/$ARCHIVE_NAME"

# === PUSH TO REMOTE ===
cd "$BACKUP_DIR" || exit 1

git add "$ARCHIVE_NAME"
git commit -m "Backup on $TIMESTAMP"
git push

echo "✅ Backup pushed to remote."

# === CLEANUP OLD BACKUPS (older than 7 days) ===
find "$BACKUP_DIR" -type f -name "backup_*.tar.gz" -mtime +7 -delete
