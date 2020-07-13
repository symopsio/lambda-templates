output "approve_function_arn" {
  description = "The arn of the approve function"
  value       = module.lambda.approve_function_arn
}

output "cicd_keys" {
  description = "CICD Access Keys"
  value       = module.cicd.cicd_keys
  sensitive   = true
}

output "expire_function_arn" {
  description = "The arn of the expire function"
  value       = module.lambda.expire_function_arn
}

output "sym_execute_role_arn" {
  description = "The ARN of the cross-account invocation role"
  value       = module.lambda.sym_execute_role_arn
}
