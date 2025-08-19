resource "aws_security_group" "dev_ecs_sg" {
  name        = "dev-ecs-sg"
  description = "Allow HTTP traffic"
  vpc_id      = aws_vpc.dev_vpc.id

  ingress {
    description = "Allow HTTP"
    from_port   = 8080
    to_port     = 8080
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]  # public access
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]  # allow all outbound
  }

  tags = {
    Name = "dev-ecs-sg"
  }
}
