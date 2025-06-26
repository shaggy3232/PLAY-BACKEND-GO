variable "db_username" {
  type        = string
  description = "Database username"
}

variable "db_password" {
  type        = string
  description = "Database password"
  sensitive   = true
}

variable "domain_name" {
  type        = string
  description = "Database username"
}

variable "route53_zone_id" {
  type        = string
  description = "Database password"
  sensitive   = true
}