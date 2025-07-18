resource "aws_ecr_repository" "backend" {
  name                 = "play-backend"
  image_tag_mutability = "MUTABLE"



  encryption_configuration {
    encryption_type = "AES256"
  }
}