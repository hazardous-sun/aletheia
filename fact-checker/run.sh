#!/bin/bash

printUsage() {
  echo "Usage: run.sh [OPTIONS]"
  echo "Pod:"
  echo -e "  -r --RESET \t \t Resets the project images"
  echo -e "  -d -- DELETE \t \t Deletes the previously stored data"
  echo -e "Postgres database:"
  echo -e "  --DB_HOST= \t \t Value used to identify the container where the database is running"
  echo -e "  --DB_PORT= \t \t Value used to identify the port where the database is running"
  echo -e "  --DB_NAME= \t \t Value passed to POSTGRES_DB during the database initialization"
  echo -e "  --DB_PASSWORD= \t Value passed to POSTGRES_PASSWORD during the database initialization"
  echo -e "  --DB_USER= \t \t Value passed to POSTGRES_USER during the database initialization"
  echo -e "  --SERVER_PORT= \t \t Value used to set the port the API server will use"
  echo "Miscelaneous:"
  echo -e "  -h --HELP \t \t Shows the script usage"
}

# Setting environment variables
export DB_HOST="news-db"
export DB_PORT="5432"
export DB_USER="postgres"
export DB_PASSWORD="1234"
export DB_NAME="postgres"
export SERVER_PORT="8000"

# Parse short options
while getopts ":drh" opt; do
  case $opt in
    d) # deletes previously stored data
      echo "Running 'rm pgdata -r'..."
      sudo rm pgdata -r
      ;;
    r) # resets project images
      echo "Clearing previous project images..."
      podman image rm news-db fact-check-api
      ;;
    h) # shows the usage of the script
      printUsage
      exit 2
      ;;
    \?) # handles invalid options
      echo "Invalid option: -$OPTARG" >&2
      printUsage
      exit 1
      ;;
  esac
done

# Shift the parsed options away, leaving only positional arguments
shift $((OPTIND - 1))

# Parse long options
for arg in "$@"; do
  case "$arg" in
    --DELETE)
      echo "Running 'rm pgdata -r'..."
      sudo rm pgdata -r
      ;;
    --RESET)
      echo "Clearing previous project images..."
      podman image rm news-db fact-check-api
      ;;
    --DB_HOST=*)
      VALUE="${arg#*=}"
      export DB_HOST="$VALUE"
      echo "DB_HOST value overwritten"
      ;;
    --DB_PORT=*)
      VALUE="${arg#*=}"
      export DB_PORT="$VALUE"
      echo "DB_PORT value overwritten"
      ;;
    --DB_USER=*)
      VALUE="${arg#*=}"
      export DB_USER="$VALUE"
      echo "DB_USER value overwritten"
      ;;
    --DB_PASSWORD=*)
      VALUE="${arg#*=}"
      export DB_PASSWORD="$VALUE"
      echo "DB_PASSWORD value overwritten"
      ;;
    --DB_NAME=*)
      VALUE="${arg#*=}"
      export DB_NAME="$VALUE"
      echo "DB_NAME value overwritten"
      ;;
    --SERVER_PORT=*)
      VALUE="${arg#*=}"
      export SERVER_PORT="$VALUE"
      echo "SERVER_PORT value overwritten"
      ;;
    --HELP)
      printUsage
      exit 2
      ;;
    *)
      # Handle unknown arguments
      echo "Unknown argument: $arg" >&2
      printUsage
      exit 1
      ;;
  esac
done

# Cleanup previously created containers of this project
podman-compose down

# Run service
podman-compose up -d