#!/bin/bash

# Define the base directory relative to the current script location
BASE_DIR="out/client"

# Navigate to the base directory
cd "$BASE_DIR" || exit

# Copy all .go files and README.md, preserving the directory structure
find . -type f \( -name "*.go" -o -name "README.md" \) -exec bash -c 'mkdir -p "../${1%/*}" && cp "$1" "../$1"' _ {} \;

# Copy entire directories as specified
for dir in test docs api; do
  if [ -d "$dir" ]; then
    # Using rsync to copy directories to preserve structure
    rsync -av --progress "$dir" "../$dir"
  fi
done
