terraform {
  required_version = "1.5.7"

  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "6.7.0"
    }
    spacelift = {
      source  = "spacelift-io/spacelift"
      version = "1.19.0"
    }
  }
}

provider "google" {
  project = var.google_project
  region  = var.google_region

  add_terraform_attribution_label = true
}

provider "spacelift" {
  api_key_endpoint = var.spacelift_endpoint
  api_key_id       = var.spacelift_key_id
  api_key_secret   = var.spacelift_key_secret
}
