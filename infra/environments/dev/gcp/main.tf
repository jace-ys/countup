data "google_project" "this" {
  project_id = var.google_project
}

output "google_project" {
  value = data.google_project.this
}

output "foo" {
  value = "bar"
}
