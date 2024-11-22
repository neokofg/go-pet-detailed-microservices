variable "timeweb_token" {
  description = "Timeweb Cloud API token"
  type        = string
  sensitive   = true
}

variable "region" {
  description = "Region for the clusters"
  type        = string
  default     = "ru-1"  # Timeweb Cloud регион (ru-1 - Москва)
}

variable "environment" {
  description = "Environment (dev/staging/prod)"
  type        = string
  default     = "dev"
}

variable "cluster_node_count" {
  description = "Number of nodes in each cluster"
  type        = number
  default     = 3
}

variable "k8s_version" {
  description = "Kubernetes version"
  type        = string
  default     = "1.28"
}
