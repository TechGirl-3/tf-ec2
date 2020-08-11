output "instance_id" {
    description = "Instance ID"
    value = aws_instance.myinstance.*.id
}

output "instance-key" {
    description = "Instance Key"
    value = aws_instance.myinstance.*.key_name
}

output "instance-tags" {
    description = "Instance Tags"
    value = aws_instance.myinstance.*.tags
}
