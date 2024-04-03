#!/bin/bash

# Define the base directory relative to the current script location
BASE_DIR="out/client"
SUB_DIRS=("test" "docs" "api")

# Navigate to the base directory
cd "$BASE_DIR" || exit

# Copy all .go files and README.md, preserving the directory structure
find . -type f \( -name "*.go" -o -name "*.md" -o -name "openapi.yaml" \) -exec bash -c 'mkdir -p "../../../${1%/*}" && cp "$1" "../../../$1"' _ {} \;

