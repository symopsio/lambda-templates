variable "account_id" {
  description = "AWS Account ID"
}

variable "app" {
  description = "App name"
}

variable "filename" {
  description = "Local file with the initial function code"
}

variable "region" {
  default = "us-east-1"
}

variable "sym_account_id" {
  description = "Account ID for Sym Cross-Account Invocation"
  default     = "803477428605"
}
