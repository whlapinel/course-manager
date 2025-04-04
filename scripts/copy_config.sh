#!/bin/bash
scp compose.yaml root@159.223.174.80:/root/compose.yaml
scp compose.prod.yaml root@159.223.174.80:/root/compose.prod.yaml
scp -r sites/assets/ root@159.223.174.80:/root/sites
scp Taskfile.yml root@159.223.174.80:/root
scp ./scripts/backup.sh root@159.223.174.80:/root/scripts
scp ./scripts/restore_backup.sh root@159.223.174.80:/root/scripts
scp .env root@159.223.174.80:/root
