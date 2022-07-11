# Sym Lambda Templates

Sym's serverless integrations let you use Sym workflows to manage resources that the Sym platform does not directly integrate with. Sym invokes your custom function with the appropriate metadata and then your function implementation updates your backend systems.

For full documentation, check out the [AWS Lambda Access Strategy Docs](https://docs.symops.com/docs/aws-lambda).

## Quickstart

Check out our [AWS Lambda Strategy Example](https://github.com/symopsio/examples/tree/main/aws_lambda_strategy) to see Sym's Lambda Strategy in action!

## Schema

The schema of Sym's requests to you can be found [here](https://json-schema.app/view/%23?url=https%3A%2F%2Fgist.githubusercontent.com%2Fdruiz%2Feff2a439a3239812b5aa05754bc7d70a%2Fraw%2F1888293393757de0722b7cf30206f75c7e048e49%2Flog_entry_schema.json), with example payloads in the [`test`](test) directory.

## Deploying

Each template includes a `build.sh` that you can use to package your AWS Lambda function for deployment. Once your function is packaged, you can upload it to S3 or use the `update-function-code` API:

```
$ aws lambda update-function-code \
  --function-name ${FUNCTION_NAME} \
  --zip-file fileb://handler.zip
```

## Python

### Lambda Configuration

You can try the template out by starting with this Terraform snippet or the equivalent in another deployment tool:

```
resource "aws_lambda_function" "this" {
  ...
  s3_bucket = "sym-releases"
  s3_key = "lambda-templates/python.zip"
  handler = "handler.handle"
  runtime = "python3.9"
  ...
}
```

### Local Testing

[`handler.py`](python/handler.py) lets you run the handler locally using:

```
$ pip install -r requirements.txt
$ # -e for an escalate payload, -d for a deescalate payload
$ python handler.py [-e | -d]
```

### Packaging

[`build.sh`](python/build.sh) builds a `handler.zip` that is appropriate for deploying to an AWS Lambda function.

## TypeScript

### Lambda Configuration

You can try the template out by starting with this Terraform snippet or the equivalent in another deployment tool:

```
resource "aws_lambda_function" "this" {
  ...
  s3_bucket = "sym-releases"
  s3_key = "lambda-templates/typescript.zip"
  handler  = "index.handler"
  runtime  = "nodejs14.x"
  ...
}
```

### Local Testing

[`index.ts`](typescript/src/index.ts) lets you run the handler locally using:

```
$ npm install
$ # -e for an escalate payload, -d for a deescalate payload
$ npx ts-node src/index.ts [-e | -d]
```

### Packaging

[`build.sh`](typescript/build.sh) builds a `handler.zip` that is appropriate for deploying to an AWS Lambda function.

## Go

### Lambda Configuration

You can try the template out by starting with this Terraform snippet or the equivalent in another deployment tool:

```
resource "aws_lambda_function" "this" {
  ...
  s3_bucket = "sym-releases"
  s3_key = "lambda-templates/go.zip"
  handler  = "bin/handler"
  runtime  = "go1.x"
  ...
}
```

### Local Testing

[`cmd/local`](go/cmd/local/main.go) lets you run the handler locally using:

```
$ cd cmd/local
$ # -e for an escalate payload, -d for a deescalate payload
$ go run . [-e | -d]
```

### Packaging

[`Makefile`](go/Makefile) builds a `handler.zip` that is appropriate for deploying to an AWS Lambda function.

## Get in touch

Please [reach out](https://docs.symops.com/docs/support) with any questions on this example or help getting started.
