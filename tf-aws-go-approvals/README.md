# tf-aws-go-approvals

An example Sym approvals Lambda written in Go and deployed with Terraform.

## Prerequisites

* Terraform
* Admin AWS privileges

## Deployment Steps

1. Build the Lambdas and then deploy.

    ```
    $ make dist
    $ cd tf/examples/demo && terraform apply
    ```

## Updating The Code

Once the Lambdas are deployed, you can update the code using `make deploy`. This updates the function code without monkeying with the Terraform configuration.

## Changing the app name

By default the app name is `sym-approvals`. If you want to deploy with a different app name, you can create a `.tfvars` file and configure the new name in there. You'll also need to update the `Makefile`.

## Testing

Refer to the shared [test events](../test) that you can use to verify that the lambda is functioning. You can configure these test events in the AWS console and then verify that things are working as expected.

## Integrating with CICD

The example includes a CICD user that you can use to enable this workflow in your CICD system. `terraform show` allows you to see the AWS access keys you'll need.
