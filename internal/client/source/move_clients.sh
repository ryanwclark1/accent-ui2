#!/bin/bash

# Define the base directory relative to the current script location
BASE_DIR="out/client"
SUB_DIRS=("test" "docs" "api")

# Navigate to the base directory
cd "$BASE_DIR" || exit

# Copy all .go files and README.md, preserving the directory structure
find . -type f \( -name "*.go" -o -name "*.md" -o -name "openapi.yaml" \) -exec bash -c 'mkdir -p "../../../${1%/*}" && cp -u "$1" "../../../$1"' _ {} \;

cd ../.. || exit

# Define the directory to search in
SEARCH_DIR="../"
SEARCH_DIR=$(realpath "$SEARCH_DIR")

# Define the directory to exclude from the search
EXCLUDE_DIR="${SEARCH_DIR}/${BASE_DIR}"

# Define the text to search for
SEARCH_TEXT="github.com/ryanwclark1/accent-ui2/internal/client/"

# Define the replacement text
REPLACE_TEXT="github.com/ryanwclark1/accent-ui2/internal/client/"

# Find all files in the directory containing the search text
# and replace the text in those files, excluding the specified directory.
find "$SEARCH_DIR" -type f ! -path "$EXCLUDE_DIR/*" -exec grep -lZ "$SEARCH_TEXT" {} + |
    while IFS= read -rd '' file; do
        sed -i "s|${SEARCH_TEXT}|${REPLACE_TEXT}|g" "$file"
        echo "Updated: $file"
    done

echo "Replacement complete."
