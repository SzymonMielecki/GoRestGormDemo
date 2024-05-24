#!/bin/sh

set -e

host="$1"
shift
cmd="$@"

echo "Waiting for $host to be ready..."

until PGPASSWORD=$POSTGRES_PASSWORD psql -h "$host" -U "postgres" -c '\q'; do
    >&2 echo "Postgres is unavailable - sleeping"
    sleep 1
done

echo "$host is ready - executing command $cmd"

exec $cmd