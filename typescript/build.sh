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

mkdir -p dist
cp package.json dist

echo "Compiling..."
npm run-script build

pushd dist || exit 1

  echo "Installing prod dependencies..."
  npm install --only=production

  echo "Creating handler.zip..."
  zip -q -r ../handler.zip .

popd || exit 1

echo "Build complete!"
