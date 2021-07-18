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

resource "google_project_service" "cloud_run" {
  project                    = var.project_name
  service                    = "run.googleapis.com"
  disable_dependent_services = true
}

resource "google_project_service" "secret_manager" {
  project                    = var.project_name
  service                    = "secretmanager.googleapis.com"
  disable_dependent_services = true
}