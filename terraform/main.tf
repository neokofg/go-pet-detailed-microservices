# Frontend Cluster
module "frontend_cluster" {
  source = "./modules/cluster"

  cluster_name    = "frontend-cluster"
  region         = var.region
  environment    = var.environment
  node_count     = var.cluster_node_count
  k8s_version    = var.k8s_version
}

# Auth Cluster
module "auth_cluster" {
  source = "./modules/cluster"

  cluster_name    = "auth-cluster"
  region         = var.region
  environment    = var.environment
  node_count     = var.cluster_node_count
  k8s_version    = var.k8s_version
}

# Database Cluster
module "database_cluster" {
  source = "./modules/cluster"

  cluster_name    = "database-cluster"
  region         = var.region
  environment    = var.environment
  node_count     = var.cluster_node_count
  k8s_version    = var.k8s_version
}
