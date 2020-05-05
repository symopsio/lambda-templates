# sym-aws-python-approvals

Refer to the top-level [README](../README.md) for general information on this template.

## Python protobufs

[Protocol Buffer Basics: Python](https://developers.google.com/protocol-buffers/docs/pythontutorial) has examples of how to work with the [event objects](https://github.com/symopsio/types) that you receive from Sym.

## Local testing

[handler.py](handler.py) includes a main method with an example approval event.

 This lets you run the lambda locally using:
```
$ python handler.py
```

## Local cache

The [serverless-requirements-plugin](https://www.serverless.com/plugins/serverless-python-requirements/) will cache your dependencies. Run `sls requirements cleanCache` to clear the cache if needed.

## Serverless Template

This template was uses the [serverless-python-requirements](https://www.serverless.com/plugins/serverless-python-requirements/) plugin to load in dependencies.

## Get in touch

Please reach out to info@symops.io with any questions on this example or help getting started.
