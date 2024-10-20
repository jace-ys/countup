resource "spacelift_stack" "root" {
  space_id                         = "root"
  name                             = "root"
  description                      = "🚀 Root stack for managing other Spacelift stacks"
  repository                       = "countup"
  branch                           = "main"
  project_root                     = "infra/spacelift/root"
  terraform_workflow_tool          = "OPEN_TOFU"
  terraform_version                = "1.9.1"
  administrative                   = true
  protect_from_deletion            = true
  terraform_smart_sanitization     = true
  enable_well_known_secret_masking = true

  labels = ["feature:add_plan_pr_comment", "terraform-provider-google"]
}

resource "spacelift_gcp_service_account" "root" {
  stack_id = spacelift_stack.root.id
  token_scopes = [
    "https://www.googleapis.com/auth/compute",
    "https://www.googleapis.com/auth/cloud-platform",
    "https://www.googleapis.com/auth/ndev.clouddns.readwrite",
    "https://www.googleapis.com/auth/devstorage.full_control",
    "https://www.googleapis.com/auth/userinfo.email",
  ]
}

resource "google_project_iam_member" "spacelift_editor" {
  project = var.google_project
  role    = "roles/editor"
  member  = "serviceAccount:${spacelift_gcp_service_account.root.service_account_email}"
}

resource "google_project_iam_member" "spacelift_project_iam_admin" {
  project = var.google_project
  role    = "roles/resourcemanager.projectIamAdmin"
  member  = "serviceAccount:${spacelift_gcp_service_account.root.service_account_email}"
}

resource "google_project_iam_member" "spacelift_secret_accessor" {
  project = var.google_project
  role    = "roles/secretmanager.secretAccessor"
  member  = "serviceAccount:${spacelift_gcp_service_account.root.service_account_email}"
}
