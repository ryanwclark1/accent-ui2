#!/bin/bash

# Get all YAML or YML files in the directory
files=$(find . -type f \( -name "accent*.yaml" -o -name "accent*.yml" \))
GO_POST_PROCESS_FILE="/workdir/go-post-process-file.sh"

# Loop through each YAML/YML file
for file in $files; do
    # Extract filename without extension
    filename=$(basename "$file")
    filename_no_ext="${filename%.*}"


    # Remove hyphens from the filename
    filename_no_accent="${filename_no_ext//accent-/}"
    filename_no_ext_hyphens="${filename_no_accent//-/}"

    # Generate code using openapi-generator-cli
    # https://openapi-generator.tech/docs/generators/go-server
    docker run -e --rm \
    -e GO_POST_PROCESS_FILE='/workdir/go-post-process-file.sh' \
    -v ${PWD}:/workdir openapi-generator-cli generate \
    --enable-post-process-file \
    -i /workdir/${filename} \
    -g go-server \
    --additional-properties=packageName=${filename_no_ext_hyphens}server,packageVersion=0.0.1,addResponseHeader=true,enumClassPrefix=falsem,featureCORS,hideGenerationTimestamp=false,onlyInterfaces=true,outputAsLibrary=true,router=chi,serverport=8081,sourceFolder=src \
    --git-user-id ryanwclark \
    --git-repo-id accent-voice \
    --global-property=generateAliasAsModel \
    -o /workdir/out/server/${filename_no_accent}


    # Change ownership of newly created files and directories recursively
    docker run \
    -v ${PWD}:/workdir alpine chown -R $(id -u):$(id -g) /workdir/out/server/${filename_no_accent}

    docker run \
    -v ${PWD}:/workdir davidanson/markdownlint-cli2 --fix "/workdir/out/server/${filename_no_accent}/docs/*.md" "#node_modules"
done

    docker run \
    -v ${PWD}:/workdir alpine chown -R $(id -u):$(id -g) /workdir/out
