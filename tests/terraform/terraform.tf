terraform {
  required_version = ">= 1.3.6"

  required_providers {
    helm = {
      source  = "hashicorp/helm"
      version = "2.16.1"
    }
  }
}
