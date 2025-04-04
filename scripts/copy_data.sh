#!/bin/bash
tar -czvf users.tar.gz internal/data/users
tar -czvf database.tar.gz internal/data/database/course_manager.db
scp ./scripts/unpack_data.sh root@159.223.174.80:/root/scripts/unpack_data.sh
scp database.tar.gz root@159.223.174.80:/root/
scp users.tar.gz root@159.223.174.80:/root/
rm database.tar.gz
rm users.tar.gz
