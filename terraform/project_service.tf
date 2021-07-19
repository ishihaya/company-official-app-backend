resource "google_project_service" "service_usage" {
  project                    = var.project_name
  service                    = "serviceusage.googleapis.com"
  disable_dependent_services = true
}

resource "google_project_service" "iam" {
  project                    = var.project_name
  service                    = "iam.googleapis.com"
  disable_dependent_services = true
}

resource "google_project_service" "kubernetes" {
  project                    = var.project_name
  service                    = "container.googleapis.com"
  disable_dependent_services = true
}