variable "instance_count" {
  description = "Number of instances to create"
  default       = "1"
}

variable "instance_type" {
  description = "compute instance size"
  default      = "t2.micro"
}


variable "key_name" {
  description = "name of public key pair to associate with instance"
  default     = "gotest"
}

variable "ami" {
  description = "ID of AMI to use for the instance"
  default      = "ami-02354e95b39ca8dec"
}

variable "environment" {
  description = "ENV"
  default = "Test"
}

variable "name" {
  description = "Name"
  default = "testterra"
}
