#!/bin/bash

# Cleanup previously created containers of this project
podman-compose down

# Setting environment variables
export DB_HOST="news-db"
export DB_PORT="5432"
export DB_USER="postgres"
export DB_PASSWORD="1234"
export DB_NAME="postgres"

# Parse command-line options
while getopts ":dr" opt; do
  case $opt in
    r) # reset project images
      echo "Clearing previous project images..."
      podman image rm news-db fact-check-api
      ;;
    d) # delete previously stored data
      echo "Running 'rm pgdata -r'..."
      sudo rm pgdata -r
      ;;
    \?)
      # Handle invalid options
      echo "Invalid option: -$OPTARG" >&2
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
    *)
      # Handle unknown arguments
      echo "Unknown argument: $arg" >&2
      ;;
  esac
done

# Run service
podman-compose up -d