resource "spacelift_context" "terraform_provider_google" {
  space_id    = spacelift_stack.root.id
  name        = "terraform-provider-google"
  description = "Configuration for terraform-provider-google"
  labels      = ["autoattach:terraform-provider-google"]
}

resource "spacelift_environment_variable" "google_project" {
  context_id = spacelift_context.terraform_provider_google.id
  name       = "GOOGLE_PROJECT"
  value      = var.google_project
  write_only = false
}

resource "spacelift_environment_variable" "google_region" {
  context_id = spacelift_context.terraform_provider_google.id
  name       = "GOOGLE_REGION"
  value      = var.google_region
  write_only = false
}
