#!/usr/bin/env bash

ERROR="\033[31m"
NC="\033[0m"

printUsage() {
  echo "Usage: run.sh [OPTIONS]"
  echo "Options:"
  echo -e "  --PORT=\t Specify which port should be used for the FastAPI service (default: 7654)"
  echo -e "  -h --HELP\t Shows this help message"
}

# Default port
PORT="7654"

# Parse command-line options
while [[ $# -gt 0 ]]; do
  case "$1" in
    --PORT=*)
      VALUE="${1#*=}"
      if [ "$VALUE" == "" ]; then
        echo -e "${ERROR}Error: PORT value cannot be empty${NC}"
        printUsage
        exit 1
      fi
      PORT="$VALUE"
      echo "Using custom port: $PORT"
      ;;
    -h|--HELP)
      printUsage
      exit 0
      ;;
    *)
      echo -e "${ERROR}Error: Unknown argument '$1'${NC}" >&2
      printUsage
      exit 1
      ;;
  esac
  shift
done

# Build the Docker image
echo "Building container image..."
podman build . -t ai-analyzer

# Run the container with port mapping
echo "Starting container..."
podman run -d \
  --name ai-analyzer \
  -p $PORT:7654 \
  ai-analyzer

echo -e "\nService running on port $PORT"