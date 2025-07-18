variable "image_tag" {
  description = "Docker image tag to deploy"
  type        = string
}

variable "database_url" {
  description = "The Supabase PostgreSQL connection string"
  type        = string
  sensitive   = true
}