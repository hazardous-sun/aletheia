#!/bin/bash

# flags used for building the GUI
flags=(0, 0, 0) # 1. context 2. image 3. video

printUsage() {
  echo "Usage: run.sh [OPTIONS]"
  echo "Inteface:"
  echo "  -C --CONTEXT \t Enables the context text field section"
  echo "  -I --IMAGE \t Enables the image check field section"
  echo "  -V --VIDEO \y Enables the video check field section"
  echo "Miscelaneous:"
  echo "  -h --HELP \t Shows the script usage"
}

enableSection() {
  flags["$1"]="1"
}

# Parse command-line options
while [[ $# -gt 0 ]]; do
  case "$1" in
    -C|--CONTEXT)
      echo "Enabling context section"
      enableSection 0
      ;;
    -I|--IMAGE)
      echo "Enabling image section"
      enableSection 1
      ;;
    -V|--VIDEO)
      echo "Enabling video section"
      enableSection 2
      ;;
    -h|--HELP)
      printUsage
      exit 0
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

# Setting environment variables
export CONTEXT=${flags[0]}
export IMAGE=${flags[1]}
export VIDEO=${flags[2]}

# Run the client application
go run src/cmd/main.go
