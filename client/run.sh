#!/bin/bash

printUsage() {
  echo "Usage: run.sh [OPTIONS]"
  echo "Inteface:"
  echo -e "  -I --IMAGE \t Enables the image check field section"
  echo -e "  -P --PROMPT \t Enables the context text field section"
  echo -e "  -V --VIDEO \y Enables the video check field section"
  echo "Miscelaneous:"
  echo -e "  -h --HELP \t Shows the script usage"
}

# Set the port used to communicate with the server
export PORT="8000"

# Build the client application
go build -v -o client src/cmd/main.go

# Make the compiled code executable
chmod +x ./client

# Run the client application
./client "$@"

# Remove compiled code
rm client