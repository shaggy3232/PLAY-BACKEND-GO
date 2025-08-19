variable "app_name" {
  description = "Name of the Fly app"
  type        = string
}

variable "org" {
  description = "Fly.io organization"
  type        = string
  default     = "personal" # works if you're just using your personal account
}

variable "region" {
  description = "Primary region (e.g., ord, yyz, sea, fra)"
  type        = string
  default     = "ord"
}

variable "image" {
  description = "Docker image to deploy"
  type        = string
}

variable "database_url" {
  description = "The Supabase PostgreSQL connection string"
  type        = string
  sensitive   = true
}