#!/bin/bash

printUsage() {
  echo "Usage: run.sh [OPTIONS]"
  echo "Pod:"
  echo -e "  -r \t \t \t Resets the project images"
  echo -e "  -d \t \t \t Deletes the previously stored data"
  echo -e "Postgres database:"
  echo -e "  --DB_HOST= \t \t Value used to identify the container where the database is running"
  echo -e "  --DB_PORT= \t \t Value used to identify the port where the database is running"
  echo -e "  --DB_NAME= \t \t Value passed to POSTGRES_DB during the database initialization"
  echo -e "  --DB_PASSWORD= \t Value passed to POSTGRES_PASSWORD during the database initialization"
  echo -e "  --DB_USER= \t \t Value passed to POSTGRES_USER during the database initialization"
  echo "Miscelaneous:"
  echo -e "  -h --HELP \t \t Shows the script usage"
}

# Parse command-line options
while getopts ":dhr" opt; do
  case $opt in
    d) # deletes previously stored data
      echo "Running 'rm pgdata -r'..."
      sudo rm pgdata -r
      ;;
    r) # resets project images
      echo "Clearing previous project images..."
      podman image rm news-db fact-check-api
      ;;
    h) # shows the usage of the scrip
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

# Overwrite default values for the environment variables
for arg in "$@"; do
  case "$arg" in
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

# Setting environment variables
export DB_HOST="news-db"
export DB_PORT="5432"
export DB_USER="postgres"
export DB_PASSWORD="1234"
export DB_NAME="postgres"

# Run service
podman-compose up -d