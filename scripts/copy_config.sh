#!/bin/bash
scp scripts/remote.sh root@159.223.174.80:/root/scripts
scp compose.prod.yaml root@159.223.174.80:/root/compose.yaml
scp Taskfile.yml root@159.223.174.80:/root
scp .env root@159.223.174.80:/root
