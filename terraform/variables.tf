variable "region" {
  description = "Region for the clusters"
  type        = string
  default     = "us-west-1"
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
