#!/usr/bin/env bash
export DBSYNC_HOST="localhost:8080"
export DBSYNC_WORK_DIR="/tmp"
export DBSYNC_WORKERS=5
# Default database connection
export DBSYNC_DB_HOST="unix(/var/run/mysqld/mysqld1.sock)"
export DBSYNC_DB_USERNAME="root"
export DBSYNC_DB_PASSWORD="root"
# Other database connections
export DBSYNC_DB_HOST_1="unix(/var/run/mysqld/mysqld2.sock)"
export DBSYNC_DB_USERNAME_1="root"
export DBSYNC_DB_PASSWORD_1="root"
# Seperate database names with ,
export DBSYNC_DB_DATABASE_1="database"