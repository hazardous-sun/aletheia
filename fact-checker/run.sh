#!/bin/bash

printUsage() {
  echo "Usage: run.sh [OPTIONS]"
  echo "Pod:"
  echo -e "  -C --CLEAR \t \t Clears the database"
  echo -e "  -R --RESET \t \t Resets the project images"
  echo -e "Fact check server connection:"
  echo -e "  --SERVER_PORT= \t Value used to set the port the API server will use"
  echo -e "Postgres database connection:"
  echo -e "  --DB_HOST= \t \t Value used to identify the container where the database is running"
  echo -e "  --DB_PORT= \t \t Value used to identify the port where the database is running"
  echo -e "  --DB_NAME= \t \t Value passed to POSTGRES_DB during the database initialization"
  echo -e "  --DB_PASSWORD= \t Value passed to POSTGRES_PASSWORD during the database initialization"
  echo -e "  --DB_USER= \t \t Value passed to POSTGRES_USER during the database initialization"
  echo "Miscelaneous:"
  echo -e "  -h --HELP \t \t Shows the script usage"
}

clearDatabase() {
  echo "Clearing database..."

  if [[ -d "pgdata" ]]; then
    RESULT=$(sudo rm pgdata -r)
    if [[ "$RESULT" != 0 ]]; then
        echo "An error occurred while clearing the database"
        exit 1
    fi
  fi

  echo "Successfully cleared the database!"
}

clearImages() {
  echo "Clearing previous project images..."

  # Check if the images exist
  for IMAGE in news-db fact-check-api; do
    if podman image exists "$IMAGE"; then
      removeImage "$IMAGE"
    fi
  done

  echo "Successfully cleread previous project images!"
}

removeImage() {
  RESULT=$(podman image rm $1)

  if [[ "$RESULT" == 0 ]]; then
    echo "Successfully removed previous $1 image"
  else
    echo "An error occurred while trying to remove image $1"
    exit 1
  fi
}

# Setting environment variables
export DB_HOST="news-db"
export DB_PORT="5432"
export DB_USER="postgres"
export DB_PASSWORD="1234"
export DB_NAME="postgres"
export SERVER_PORT="8000"

# Cleanup previously created containers of this project
podman-compose down --remove-orphans

# Parse command-line options
while [[ $# -gt 0 ]]; do
  case "$1" in
    -C|--CLEAR) # clears the database
      clearDatabase
      ;;
    -R|--RESET)
      clearImages
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
    -h|--HELP)
      printUsage
      exit 2
      ;;
    -*)
      echo "Invalid option: $1" >&2
      printUsage
      exit 1
      ;;
    *)
      echo "Invalid argument: $1" >&2
      printUsage
      exit 1
      ;;
  esac
  shift
done

# Run service
podman-compose up -d