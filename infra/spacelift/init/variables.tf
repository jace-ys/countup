variable "google_project" {
  type    = string
  default = "emp-jace-1fb5"
}

variable "google_region" {
  type    = string
  default = "europe-west1"
}

variable "spacelift_endpoint" {
  type    = string
  default = "https://jace-ys.app.spacelift.io"
}

variable "spacelift_key_id" {
  type    = string
  default = "01JAN9TH8H60PCN3N30V1ZES81"
}

variable "spacelift_key_secret" {
  type      = string
  sensitive = true
}
