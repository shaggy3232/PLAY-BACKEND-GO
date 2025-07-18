#!/bin/sh
set -e

# No jq needed if it's just a raw string
echo "Running Goose migrations..."
echo $DATABASE_URL
goose -dir /migrations postgres "$DATABASE_URL" up

echo "Starting Go backend..."
exec /app/playgobackend
