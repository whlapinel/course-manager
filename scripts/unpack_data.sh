#!/bin/bash
tar -xzvf users.tar.gz internal/data
rm users.tar.gz
tar -xzvf database.tar.gz internal/data/database
rm database.tar.gz