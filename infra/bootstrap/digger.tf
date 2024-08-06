resource "google_iam_workload_identity_pool" "digger" {
  project                   = var.google_project
  workload_identity_pool_id = "digger"
}

resource "google_iam_workload_identity_pool_provider" "digger" {
  project                            = var.google_project
  workload_identity_pool_id          = google_iam_workload_identity_pool.digger.workload_identity_pool_id
  workload_identity_pool_provider_id = "digger"
  display_name                       = "Digger"
  description                        = "OIDC identity pool provider for Digger"

  attribute_mapping = {
    "google.subject"       = "assertion.sub"
    "attribute.repository" = "assertion.repository"
    "attribute.ref"        = "assertion.ref"
  }

  attribute_condition = "assertion.repository == 'jace-ys/countup'"

  oidc {
    issuer_uri = "https://token.actions.githubusercontent.com"
  }
}

resource "google_service_account" "digger" {
  project      = var.google_project
  account_id   = "digger"
  display_name = "Digger"
}

resource "google_service_account_iam_binding" "digger_workload_identity_user" {
  service_account_id = google_service_account.digger.name
  role               = "roles/iam.workloadIdentityUser"
  members = [
    "principalSet://iam.googleapis.com/${google_iam_workload_identity_pool.digger.name}/attribute.repository/jace-ys/countup"
  ]
}
