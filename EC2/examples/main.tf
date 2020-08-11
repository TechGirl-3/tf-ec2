module "ec2" {
  source                      = "../modules"
  instance_count              = var.instance_count
  ami                         = var.ami
  instance_type               = var.instance_type
  key_name                    = var.key_name
  tags = {
    "Env"      = var.environment
    "Name" = var.name
  }
}
