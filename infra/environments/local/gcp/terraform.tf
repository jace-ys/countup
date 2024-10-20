terraform {
  required_version = "1.9.1"

  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "6.34.0"
    }
  }
}

provider "google" {
  add_terraform_attribution_label = true
}

data "google_project" "this" {}

output "google_project" {
  value = data.google_project.this.project_id
}
