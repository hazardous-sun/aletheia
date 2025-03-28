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
  flags[index]="$1"
}

# Parse command-line options
while getopts ":CIVh" opt; do
  case $opt in
    C) # enables context section
      echo "Enabling context section"
      enableSection 0
      ;;
    I) # enables image section
      echo "Enabling image section"
      enableSection 1
      ;;
    V) # enables video section
      echo "Enabling video section"
      enableSection 2
      ;;
    h) # shows the usage of the script
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

# Parse the rest of the command-line options
for arg in "$@"; do
  case "$arg" in
    --CONTEXT) # enables context section
      echo "Enabling context section"
      enableSection 0
      ;;
    --IMAGE) # enables image section
      echo "Enabling image section"
      enableSection 1
      ;;
    --VIDEO) # enables video section
      echo "Enabling image section"
      enableSection 2
      ;;
    --HELP) # shows the usage of the script
      printUsage
      exit 2
      ;;
    *) # handles unknown arguments
      echo "Invalid option: $arg" >&2
      printUsage
      exit 1
      ;;
  esac
done

# Setting environment variables
export CONTEXT=flags[0]
export IMAGE=flags[1]
export VIDEO=flags[2]

# Run the client application
go run src/cmd/main.go
