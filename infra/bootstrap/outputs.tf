output "digger_service_account_email" {
  value = google_service_account.digger.email
}

output "digger_workload_identity_pool_provider_name" {
  value = google_iam_workload_identity_pool_provider.digger.name
}
