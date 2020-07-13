terraform {
  required_version = ">= 0.12.7"
}

provider "aws" {
  version = "~> 2.0"
  region  = var.region
}

data "aws_caller_identity" "current" { }

module "lambda" {
  source          = "../../modules/lambda/"
  account_id      = data.aws_caller_identity.current.account_id
  app             = var.app
  filename        = "${path.root}/../../../dist/${var.app}.zip"
  region          = var.region
  sym_account_id  = var.sym_account_id
}

module "cicd" {
  source          = "../../modules/cicd/"
  account_id      = data.aws_caller_identity.current.account_id
  app             = var.app
  region          = var.region
}
