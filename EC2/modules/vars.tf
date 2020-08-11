variable "instance_count" {
  description = "Number of instances to create"
}

variable "instance_type" {
  description = "compute instance size"
}

variable "key_name" {
  description = "name of public key pair to associate with instance"
  default = ""
}

variable "ami" {
  description = "ID of AMI to use for the instance"
}

variable "tags" {
  description = "A mapping of tags to assign to the resource"
  type        = map(string)
}
