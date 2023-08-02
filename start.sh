#!/bin/sh

set -e

echo "run db migration"
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "seed user"
psql -d $DB_SOURCE -a -f /app/seed/user.sql

echo "start the app"
exec "$@"
