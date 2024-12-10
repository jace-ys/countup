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
  add_terraform_attribution_label = true
}

provider "spacelift" {
}

data "google_project" "this" {}
