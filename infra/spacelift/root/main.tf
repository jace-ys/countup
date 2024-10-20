locals {
  environments = [for i in fileset("../../environments", "*/gcp/terraform.tf") : dirname(dirname(i))]
}

resource "spacelift_stack" "environment" {
  for_each = toset(local.environments)

  space_id                         = "root"
  name                             = each.key
  description                      = "🌳 Environment [${each.key}]"
  repository                       = "countup"
  branch                           = "main"
  project_root                     = "infra/environments/${each.key}/gcp"
  terraform_workflow_tool          = "OPEN_TOFU"
  terraform_version                = "1.9.1"
  administrative                   = false
  protect_from_deletion            = true
  terraform_smart_sanitization     = true
  enable_well_known_secret_masking = true

  labels = ["feature:add_plan_pr_comment", "terraform-provider-google"]
}

resource "spacelift_gcp_service_account" "environment" {
  for_each = toset(local.environments)

  stack_id = spacelift_stack.environment[each.key].id
  token_scopes = [
    "https://www.googleapis.com/auth/compute",
    "https://www.googleapis.com/auth/cloud-platform",
    "https://www.googleapis.com/auth/ndev.clouddns.readwrite",
    "https://www.googleapis.com/auth/devstorage.full_control",
    "https://www.googleapis.com/auth/userinfo.email",
  ]
}

resource "google_project_iam_member" "environment_editor" {
  for_each = toset(local.environments)

  project = data.google_project.this.project_id
  role    = "roles/editor"
  member  = "serviceAccount:${spacelift_gcp_service_account.environment[each.key].service_account_email}"
}

resource "google_project_iam_member" "environment_project_iam_admin" {
  for_each = toset(local.environments)

  project = data.google_project.this.project_id
  role    = "roles/resourcemanager.projectIamAdmin"
  member  = "serviceAccount:${spacelift_gcp_service_account.environment[each.key].service_account_email}"
}

resource "google_project_iam_member" "environment_secret_accessor" {
  for_each = toset(local.environments)

  project = data.google_project.this.project_id
  role    = "roles/secretmanager.secretAccessor"
  member  = "serviceAccount:${spacelift_gcp_service_account.environment[each.key].service_account_email}"
}
