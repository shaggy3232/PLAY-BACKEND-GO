#!/bin/sh

# Wait for the database to be ready
until pg_isready -h db -U postgres; do
  echo "Waiting for database..."
  sleep 2
done

# Run Goose migrations
goose -dir /migrations postgres "user=postgres password=postgres123 dbname=postgres sslmode=disable" up