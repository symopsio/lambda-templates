#!/usr/bin/env bash

# Rebuild and package the Lambda function implementation
# into a zip file appropriate for deploying to AWS Lambda

if [[ -d dist ]]; then
  echo "Cleaning existing dist directory..."
  rm -rf dist
fi

if [[ -f handler.zip ]]; then
  echo "Removing existing handler.zip..."
  rm handler.zip
fi

echo "Installing dependencies..."
pip install -r requirements.txt --target dist

echo "Creating handler.zip"
pushd dist || exit 1
zip -q -r ../handler.zip .
popd || exit 1

zip -r -g handler.zip handler.py
