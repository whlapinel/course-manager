#!/bin/bash

rsync -avz compose.yaml root@159.223.174.80:/root/compose.yaml
rsync -avz compose.prod.yaml root@159.223.174.80:/root/compose.prod.yaml
rsync -avz Taskfile.yml root@159.223.174.80:/root
rsync -avz ./scripts/backup.sh root@159.223.174.80:/root/scripts
rsync -avz ./scripts/restore_backup.sh root@159.223.174.80:/root/scripts
rsync -avz .env root@159.223.174.80:/root
