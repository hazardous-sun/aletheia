#!/bin/sh

set -e

host="$DB_HOST"
shift
port="$DB_PORT"
shift
cmd="$@"

until nc -z "$host" "$port"; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done

>&2 echo "Postgres is up - executing command"
exec $cmd