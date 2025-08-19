# ECS Cluster
resource "aws_ecs_cluster" "dev_cluster" {
  name = "dev-cluster"
}

# IAM Role for ECS Task Execution
resource "aws_iam_role" "ecs_task_execution_role" {
  name = "dev-ecs-task-execution-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [{
      Action = "sts:AssumeRole",
      Principal = {
        Service = "ecs-tasks.amazonaws.com"
      },
      Effect = "Allow",
      Sid    = ""
    }]
  })
}

# Attach AWS-managed policy to allow ECS tasks to pull from ECR and write logs
resource "aws_iam_role_policy_attachment" "ecs_execution_policy" {
  role       = aws_iam_role.ecs_task_execution_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}

# ECS Task Definition
resource "aws_ecs_task_definition" "dev_task" {
  family                   = "dev-task"
  requires_compatibilities = ["FARGATE"]
  network_mode             = "awsvpc"
  cpu                      = "256"
  memory                   = "512"
  execution_role_arn       = aws_iam_role.ecs_task_execution_role.arn

  container_definitions = jsonencode([
    {
      name      = "dev-backend"
      image     = "${aws_ecr_repository.backend.repository_url}:${var.image_tag}"
      essential = true
      portMappings = [
        {
          containerPort = 8080
          hostPort      = 8080
        }
      ]
      logConfiguration = {
        logDriver = "awslogs"
        options = {
          awslogs-group         = "/ecs/dev-backend"
          awslogs-region        = "ca-central-1"
          awslogs-stream-prefix = "ecs"
        }
      }
      environment = [
        {
          name  = "DATABASE_URL"
          value =  var.database_url
        },
        {
          name  = "GODEBUG"
          value = "netdns=go+v4"
        }
        
      ]
    }
  ])
}

resource "aws_ecs_service" "dev_service" {
  name            = "dev-service"
  cluster         = aws_ecs_cluster.dev_cluster.id
  task_definition = aws_ecs_task_definition.dev_task.arn
  desired_count   = 1
  launch_type     = "FARGATE"

  network_configuration {
    subnets         = [
      aws_subnet.public_subnet_1.id,
      aws_subnet.public_subnet_2.id
    ]
    security_groups  = [aws_security_group.dev_ecs_sg.id]
    assign_public_ip = true
  }
}

resource "aws_cloudwatch_log_group" "dev_backend" {
  name              = "/ecs/dev-backend"
  retention_in_days = 7
}