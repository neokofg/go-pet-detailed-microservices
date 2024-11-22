terraform {
  required_providers {
    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = "~> 2.0"
    }
  }
}

provider "kubernetes" {
  config_path = "~/.kube/config"
}

# Frontend Cluster
module "frontend_cluster" {
  source = "./modules/cluster"

  cluster_name = "frontend-cluster"
  region       = var.region
}

# Auth Cluster
module "auth_cluster" {
  source = "./modules/cluster"

  cluster_name = "auth-cluster"
  region       = var.region
}

# Database Cluster
module "database_cluster" {
  source = "./modules/cluster"

  cluster_name = "database-cluster"
  region       = var.region
}
