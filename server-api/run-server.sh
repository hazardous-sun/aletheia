#!/bin/bash

# Setting environment variables
export DB_HOST="news-db"
export DB_PORT="5432"
export DB_USER="postgres"
export DB_PASSWORD="1234"
export DB_NAME="postgres"

# Cleanup previously created containers of this project
podman-compose down

# Run service
podman-compose up -d