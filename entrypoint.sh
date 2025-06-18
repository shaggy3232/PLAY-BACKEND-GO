#!/bin/sh
set -e

# Extract DATABASE_URL from JSON (injected by ECS secret)
export DATABASE_URL=$(echo $DATABASE_URL | jq -r '.DATABASE_URL')

# Run DB migrations
echo "Running Goose migrations..."
goose -dir /migrations postgres "$DATABASE_URL" up

# Start the backend server
echo "Starting Go backend..."
exec /app/playgobackend