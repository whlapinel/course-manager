#!/bin/bash

# === CONFIG ===
BACKUP_FILE="$1"
TARGET_DIR="$HOME"  # or set a different root directory

# === CHECKS ===
if [ -z "$BACKUP_FILE" ]; then
  echo "❌ Usage: $0 <path-to-backup.tar.gz>"
  exit 1
fi

if [ ! -f "$BACKUP_FILE" ]; then
  echo "❌ Backup file not found: $BACKUP_FILE"
  exit 1
fi

echo "⚠️ WARNING: This will overwrite existing files in:"
echo "  - internal/data/database/"
echo "  - internal/data/users/"
read -p "Proceed with restore? (y/n): " confirm

if [ "$confirm" != "y" ]; then
  echo "❌ Restore cancelled."
  exit 1
fi

# === RESTORE ===
tar -xzf "$BACKUP_FILE" -C "$TARGET_DIR"

echo "✅ Restore complete."
