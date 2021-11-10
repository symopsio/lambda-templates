# Sym Serverless Integrations

Sym's serverless integrations let you use Sym workflows to manage resources that the Sym platform does not directly integrate with. Sym invokes your custom function with the appropriate metadata and then your function implementation updates your backend systems.

These integrations can also be used as an "escape hatch" to trigger escalations in services Sym does not yet support.

There are two serverless integrations supported today:

- [HTTP](https://docs.symops.com/docs/http)
- [AWS Lambda](https://docs.symops.com/docs/aws-lambda)

The templates in this repo will be focused on deploying AWS Lambdas, but the code can easily be reused to parse inbound Sym requests from an HTTP server of your choosing.

## Schema

The schema of Sym's requests to you can be found [here](https://sym.stoplight.io/docs/sym-reporting), with an example payload at [`payload.json`](payload.json).

## Terraform Starter Template

Refer to [tf-aws-go-approvals](tf-aws-go-approvals/README.md) for an example Sym Lambda deployed with Terraform.

## Serverless Framework Templates

In addition to our Terraform starter, we provide templates using the [Serverless Framework](https://serverless.com) to manage function creation and deployment. Use the `--template-url` parameter to `sls create` to specify one of our custom templates for your project:

### Typescript

```
$ sls create --name <your-project-name> --template-url https://github.com/symopsio/serverless-templates/tree/master/sym-aws-typescript-approvals
$ cd <your-project-name>
$ npm install
```

### Python

```
$ sls create --name <your-project-name> --template-url https://github.com/symopsio/serverless-templates/tree/master/sym-aws-python-approvals
$ cd <your-project-name>
$ npm install
```

Note: the python template includes the `dockerizePip` option set to `true`, but if you don't have Docker installed locally you can change this to `false`.

### Go

```
$ sls create --name <your-project-name> --template-url https://github.com/symopsio/serverless-templates/tree/master/sym-aws-go-approvals
$ cd <your-project-name>
$ make
```

## Implementing your custom handler

Your handler needs to provide implementations of `escalate` and `deescalate` functions.

* `escalate` occurs when a requester is approved (either by another user or automatically) for a given resource
  * An `escalate` event can be identified by an `event.type` field in the payload equal to `"escalate"`
* `deescalate` occurs when the time limit on an approval is reached
  * A `deescalate` event can be identified by an `event.type` field in the payload equal to `"deescalate"`


## AWS Integration

### IAM Role

Both our Terraform and Serverless starter templates define an IAM Role that allows Sym to invoke Lambda functions in your account. In the Terraform template this is the `sym_execute_role_arn` output, and in `serverless.yml` this is the `SymCrossAccountInvocationRole`.

### Sym Configuration

Sym will need the following configuration items in order for your function to get invoked.

Function Names support any of the formats defined for the [FunctionName](https://docs.aws.amazon.com/lambda/latest/dg/API_Invoke.html#API_Invoke_RequestSyntax) parameter of the Lambda `Invoke` API.

* `ApprovalFunctionName`: The name of the approve Lambda function
* `ExpirationFunctionName`: The name the expiration Lambda function
* `RoleARN`: The ARN of the IAM Role that Sym assumes in order to invoke the function

## Get in touch

Please [reach out](https://docs.symops.com/docs/support) with any questions on this example or help getting started.
