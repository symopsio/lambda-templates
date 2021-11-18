# Sym TypeScript Lambda Skeleton

Refer to the top-level [README](../README.md) for general information on this skeleton.

## Schema

The schema of Sym's requests to you can be found [here](https://sym.stoplight.io/docs/sym-reporting), and is represented as an interface in [sym_handler/models.ts](sym_handler/models.ts).


## Local Testing

[sym_handler/handler.ts](sym_handler/handler.ts) includes a `require.main` method with an example escalation event.

This lets you run the lambda locally using:

```
$ npm install
$ npx ts-node sym_handler/handler.ts
```
## Need Help?

Please [reach out](https://docs.symops.com/docs/support) with any questions on this example or help getting started.
