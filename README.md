# Sym Serverless Integrations

Sym's serverless integrations let you use Sym workflows to manage resources that the Sym platform does not directly integrate with. Sym invokes your custom function with the appropriate metadata and then your function implementation updates your backend systems.

These integrations can also be used as an "escape hatch" to trigger escalations in services Sym does not yet support.

There are two serverless integrations supported today:

- [HTTP](https://docs.symops.com/docs/http)
- [AWS Lambda](https://docs.symops.com/docs/aws-lambda)

The templates in this repo will be focused on deploying AWS Lambdas, but the code can easily be reused to parse inbound Sym requests from an HTTP server of your choosing.

## Schema

The schema of Sym's requests to you can be found [here](https://sym.stoplight.io/docs/sym-reporting), with an example payload at [`payload.json`](payload.json).

## Get in touch

Please [reach out](https://docs.symops.com/docs/support) with any questions on this example or help getting started.
