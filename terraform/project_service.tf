resource "google_project_service" "service_usage" {
  service = "serviceusage.googleapis.com"
  project = var.project_name
  disable_dependent_services = true
}

resource "google_project_service" "iam" {
  service = "iam.googleapis.com"
  project = var.project_name
  disable_dependent_services = true
}

resource "google_project_service" "resource_manager" {
  service = "cloudresourcemanager.googleapis.com"
  project = var.project_name
  disable_dependent_services = true
}