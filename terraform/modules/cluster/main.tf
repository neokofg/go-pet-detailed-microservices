resource "twc_kubernetes_cluster" "cluster" {
  name           = var.cluster_name
  k8s_version    = var.k8s_version
  network_driver = "calico"
  region         = var.region

  preset_id = "k8s-2"

  worker_groups {
    name       = "worker-group"
    node_count = var.node_count
  }

  tags = {
    Environment = var.environment
    Terraform   = "true"
  }
}

resource "kubernetes_namespace" "cluster_namespace" {
  metadata {
    name = var.cluster_name
  }

  depends_on = [twc_kubernetes_cluster.cluster]
}
