# Sym Go Lambda Skeleton

Refer to the top-level [README](../README.md) for general information on this skeleton.

## Schema

The schema of Sym's requests to you can be found [here](https://sym.stoplight.io/docs/sym-reporting), and is represented as an interface in [sym_handler/models.ts](sym_handler/models.ts).


## Local Testing

[handler.go](handler.go) includes a `main()` that starts the Lambda RPC server.

This lets you run the lambda locally using:

```
$ go run .
```
## Need Help?

Please [reach out](https://docs.symops.com/docs/support) with any questions on this example or help getting started.
