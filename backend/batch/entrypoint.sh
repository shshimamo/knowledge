#!/bin/sh

# The first argument is the command
CMD=$1

# Check if CMD is empty
if [ -z "$CMD" ]; then
  echo "No command specified"
  exit 1
fi

# Shift all the arguments
shift

# Check if the command exists
if [ ! -f "/go/bin/$CMD" ]; then
  echo "Unknown command: $CMD"
  exit 1
fi

# Execute the command
/go/bin/$CMD "$@"