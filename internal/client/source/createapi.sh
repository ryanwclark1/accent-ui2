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
    filename_no_ext_hyphens="${filename_no_ext//accent-/}"

    # Generate code using openapi-generator-cli
    # https://openapi-generator.tech/docs/generators/go
    docker run -e --rm \
    -e GO_POST_PROCESS_FILE='/workdir/go-post-process-file.sh' \
    -v ${PWD}:/workdir openapi-generator-cli generate \
    --enable-post-process-file \
    -i /workdir/${filename} \
    -g go \
    --additional-properties=packageName=${filename_no_ext_hyphens},packageVersion=0.0.1,prependFormOrBodyParameters=true,enumClassPrefix=true,isGoSubmodule=true,structPrefix=true,generateMarshalJSON=true,generateInterfaces=true,useOneOfDiscriminatorLookup=false,withAWSV4SignaturewithAWSV4Signature=false,withGoMod=false,withXml=false,hideGenerationTimestamp=true \
    --git-user-id ryanwclark \
    --git-repo-id accent-voice \
    --global-property=generateAliasAsModel \
    -o /workdir/out/client/${filename_no_ext_hyphens}


    # Change ownership of newly created files and directories recursively
    docker run \
    -v ${PWD}:/workdir alpine chown -R $(id -u):$(id -g) /workdir/out/client/${filename_no_ext_hyphens}

    docker run \
    -v ${PWD}:/workdir davidanson/markdownlint-cli2 --fix "/workdir/out/client/${filename_no_ext_hyphens}/docs/*.md" "#node_modules"
done

    docker run \
    -v ${PWD}:/workdir alpine chown -R $(id -u):$(id -g) /workdir/out
