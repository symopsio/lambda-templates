#!/bin/bash
set -eu

if [ -f ./go.mod ]; then
    exit 0
fi

touch go.mod

if [ -z ${GITHUB_PATH} ]; then
    echo "Set GITHUB_PATH in your Makefile"
    exit 1
fi

CURRENT_DIR=$(basename $(pwd))

CONTENT=$(cat <<-EOD
module github.com/${GITHUB_PATH}/${CURRENT_DIR}

require github.com/aws/aws-lambda-go v1.17.0
EOD
)

echo "$CONTENT" > go.mod
