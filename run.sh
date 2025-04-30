#!/usr/bin/env bash

# Colors for output
ERROR="\033[31m"
WARNING="\033[33m"
INFO="\033[36m"
NC="\033[0m"

# Default configuration
DB_HOST="aletheia-db"
DB_PORT="5432"
DB_USER="postgres"
DB_PASSWORD="1234"
DB_NAME="postgres"
SERVER_PORT="8000"
AI_PORT="7654"
CLEAR_DB=false
RESET_IMAGES=false
DEBUG=false

printUsage() {
  echo "Usage: run.sh [OPTIONS]"
  echo "Database Options:"
  echo -e "  -C --CLEAR \t\tClears the database"
  echo -e "  -R --RESET \t\tResets all project images"
  echo "Server API Configuration:"
  echo -e "  --SERVER_PORT=\tPort for the API server (default: 8000)"
  echo -e "  --DB_HOST=\t\tDatabase host (default: aletheia-db)"
  echo -e "  --DB_PORT=\t\tDatabase port (default: 5432)"
  echo -e "  --DB_NAME=\t\tDatabase name (default: postgres)"
  echo -e "  --DB_USER=\t\tDatabase user (default: postgres)"
  echo -e "  --DB_PASSWORD=\tDatabase password (default: 1234)"
  echo -e "  --DEBUG \t\tEnable debug mode (Delve debugger)"
  echo "AI Analyzer Configuration:"
  echo -e "  --AI_PORT=\t\tPort for AI service (default: 7654)"
  echo "Miscellaneous:"
  echo -e "  -h --HELP \t\tShows this help message"
}

clearDatabase() {
  echo -e "${INFO}Clearing database...${NC}"
  if [[ -d "server-api/pgdata" ]]; then
    if sudo rm -rf server-api/pgdata; then
      echo -e "${INFO}Successfully cleared the database!${NC}"
    else
      echo -e "${ERROR}Failed to clear database${NC}" >&2
      exit 1
    fi
  else
    echo -e "${WARNING}No database directory found${NC}"
  fi
}

clearImages() {
  echo -e "${INFO}Clearing project images...${NC}"
  for IMAGE in aletheia_aletheia-db aletheia_aletheia-api aletheia_aletheia-ai-analyzer; do
    if podman image exists "$IMAGE"; then
      if podman image rm "$IMAGE"; then
        echo -e "${INFO}Removed image: $IMAGE${NC}"
      else
        echo -e "${ERROR}Failed to remove image: $IMAGE${NC}" >&2
        exit 1
      fi
    fi
  done
}

# Cleanup any existing containers
echo -e "${INFO}Cleaning up existing containers...${NC}"
if ! podman-compose down --remove-orphans; then
  echo -e "${ERROR}Failed to clean up existing containers${NC}" >&2
  exit 1
fi

# Parse command-line options
while [[ $# -gt 0 ]]; do
  case "$1" in
    -C|--CLEAR)
      CLEAR_DB=true
      ;;
    -R|--RESET)
      RESET_IMAGES=true
      ;;
    --SERVER_PORT=*)
      SERVER_PORT="${1#*=}"
      echo -e "${INFO}Using custom API port: $SERVER_PORT${NC}"
      ;;
    --DB_HOST=*)
      DB_HOST="${1#*=}"
      echo -e "${INFO}Using custom DB host: $DB_HOST${NC}"
      ;;
    --DB_PORT=*)
      DB_PORT="${1#*=}"
      echo -e "${INFO}Using custom DB port: $DB_PORT${NC}"
      ;;
    --DB_NAME=*)
      DB_NAME="${1#*=}"
      echo -e "${INFO}Using custom DB name: $DB_NAME${NC}"
      ;;
    --DB_USER=*)
      DB_USER="${1#*=}"
      echo -e "${INFO}Using custom DB user: $DB_USER${NC}"
      ;;
    --DB_PASSWORD=*)
      DB_PASSWORD="${1#*=}"
      echo -e "${INFO}Using custom DB password${NC}"
      ;;
    --AI_PORT=*)
      AI_PORT="${1#*=}"
      echo -e "${INFO}Using custom AI port: $AI_PORT${NC}"
      ;;
    -h|--HELP)
      printUsage
      exit 0
      ;;
    --DEBUG)
      DEBUG=true
      echo -e "${INFO}Debug mode enabled${NC}"
      ;;
    *)
      echo -e "${ERROR}Error: Unknown option '$1'${NC}" >&2
      printUsage
      exit 1
      ;;
  esac
  shift
done

# Execute pre-run actions
if $CLEAR_DB; then
  clearDatabase
fi

if $RESET_IMAGES; then
  clearImages
fi

# Export environment variables
export DB_HOST DB_PORT DB_USER DB_PASSWORD DB_NAME SERVER_PORT AI_PORT DEBUG

# Start the services
echo -e "${INFO}Starting services with podman-compose...${NC}"
if ! podman-compose up -d; then
  echo -e "${ERROR}Failed to start services with podman-compose${NC}" >&2
  echo -e "${WARNING}Please check your podman and podman-compose installation and try again.${NC}"
  exit 1
fi

# Verify services are running
echo -e "${INFO}Service status:${NC}"
echo -e "${INFO}API Server:${NC} http://localhost:$SERVER_PORT"
echo -e "${INFO}AI Analyzer:${NC} http://localhost:$AI_PORT"

# Additional verification - check if containers are actually running
if ! podman-compose ps | grep -q "Up"; then
  echo -e "${ERROR}Some containers failed to start${NC}" >&2
  podman-compose ps
  exit 1
fi

echo -e "${INFO}All services started successfully!${NC}"