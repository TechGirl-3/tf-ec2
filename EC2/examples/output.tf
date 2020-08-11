output "instance_id" {
  description = "Instance ID"
  value       = module.ec2.instance_id
}

output "instance-key" {
  description = "Instance Key"
  value       = module.ec2.instance-key
}

output "instance-tags" {
  description = "Instance Tags"
  value = module.ec2.instance-tags
}
