variable "cluster_name" {
  description = "Name of the Kubernetes cluster"
  type        = string
}

variable "region" {
  description = "Region for the cluster"
  type        = string
}

variable "environment" {
  description = "Environment (dev/staging/prod)"
  type        = string
}

variable "node_count" {
  description = "Number of worker nodes"
  type        = number
}

variable "k8s_version" {
  description = "Kubernetes version"
  type        = string
}
