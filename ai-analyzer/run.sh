#!/usr/bin/env bash

printUsage() {
  echo "Usage: run.sh [OPTIONS]"
  echo "Container:"
  echo -e "  --PORT= \t Specify which port should be used for communicating with the analyzer"
  echo "Miscelaneous:"
  echo -e "  -h --HELP \t Shows the script usage"
}

# Setting environment variables
export PORT="7654"

# Build the Docker image
podman build . -t ai-analyzer

