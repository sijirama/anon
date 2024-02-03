#!/bin/bash

# Default value for environment
environment=""

# Parse command line options
while [ "$#" -gt 0 ]; do
  case "$1" in
    -env)
      environment="$2"
      shift 2
      ;;
    *)
      echo "Invalid option: $1"
      exit 1
      ;;
  esac
done

# Check if the environment is specified
if [ -z "$environment" ]; then
  echo "Please specify an environment using -env dev or -env prod."
  exit 1
fi

# Run commands based on the environment flag
if [ "$environment" == "dev" ]; then
  echo "Running in development mode..."
  npx tailwindcss -i ./dist/main.css -o ./dist/tailwind.css && go run .
elif [ "$environment" == "prod" ]; then
  echo "Running in production mode..."
  npx tailwindcss -i ./dist/main.css -o ./dist/tailwind.css && go build -o anon && ./anon
else
  echo "Invalid environment specified. Use -env dev or -env prod."
  exit 1
fi

