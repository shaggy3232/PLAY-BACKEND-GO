output "ecs_service_name" {
  description = "The name of the ECS service"
  value       = aws_ecs_service.dev_service.name
}

# If you have a load balancer or other resource for public URL/IP, output it here instead.
# Example placeholder:
# output "load_balancer_dns" {
#   value = aws_lb.app_alb.dns_name
# }


output "ecr_repo_url" {
  description = "The URL of your ECR repo to push images"
  value       = aws_ecr_repository.backend.repository_url
}