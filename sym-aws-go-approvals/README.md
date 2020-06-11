# sym-aws-go-approvals

Refer to the top-level [README](../README.md) for how to use this template.

## Module configuration

You'll need to update [`gomod.sh`] to properly configure the module path for your organization. Change `GITHUB_PATH` to the correct path for where you're checking your lambda in. You'll also need to update the import paths for the `internal` package in the approve and expire functions to match your updated module path.

## Serverless Template

This template was generated from the Serverless [aws-go-mod](https://github.com/serverless/serverless/tree/master/lib/plugins/create/templates/aws-go-mod) template.

## Get in touch

Please reach out to info@symops.io with any questions on this example or help getting started.
