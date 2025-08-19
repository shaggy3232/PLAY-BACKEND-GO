terraform {
  required_providers {
    fly = {
      source  = "fly-apps/fly"
      version = "~> 0.0.21"
    }
  }
}

provider "fly" {
  # Terraform will use the FLY_API_TOKEN env var
  # export FLY_API_TOKEN="your-token-here"
}

resource "fly_app" "app" {
  name     = var.app_name
  org      = var.org
  region   = var.region
}

resource "fly_machine" "server" {
  app       = fly_app.app.name
  region    = var.region
  name      = "${var.app_name}-machine"

  config {
    image   = var.image  # Example: "docker.io/username/myapp:latest"
    guest {
      cpu_kind = "shared"
      cpus     = 1
      memory_mb = 256   # free tier
    }
  }
}
