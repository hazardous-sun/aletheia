#!/usr/bin/env bash

ERROR="\033[31m"
NC="\033[0m"

printUsage() {
  echo "Usage: run.sh [OPTIONS]"
  echo "Container:"
  echo -e "  --PORT= \t Specify which port should be used for communicating with the analyzer"
  echo "Miscelaneous:"
  echo -e "  -h --HELP \t Shows the script usage"
}

# Setting environment variables
export PORT="7654"

# Parse command-line options
while [[ $# -gt 0 ]]; do
  case "$1" in
    --PORT=*)
      VALUE="${1#*=}"
      if [ "$VALUE" == "" ]; then
        echo -e "${ERROR}error: PORT value cannot be empty${NC}"
        printUsage
        exit 1
      fi
      export PORT="$VALUE"
      echo "PORT value overwritten"
      ;;
    *)
      echo "Invalid argument: $1" >&2
      printUsage
      exit 1
      ;;
  esac
  shift
done

# Build the Docker image
podman build . -t ai-analyzer

# Run the container with port mapping
podman run -d \
  --name ai-analyzer \
  -p $PORT:7654 \
  ai-analyzer