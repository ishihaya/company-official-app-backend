resource "google_project_service" "service_usage" {
  service                    = "serviceusage.googleapis.com"
  project                    = var.project_name
  disable_dependent_services = true
}

resource "google_project_service" "iam" {
  service                    = "iam.googleapis.com"
  project                    = var.project_name
  disable_dependent_services = true
}

resource "google_project_service" "resource_manager" {
  service                    = "cloudresourcemanager.googleapis.com"
  project                    = var.project_name
  disable_dependent_services = true
}

resource "google_project_service" "cloud_sql_admin" {
  service                    = "sqladmin.googleapis.com"
  project                    = var.project_name
  disable_dependent_services = true
}

resource "google_project_service" "service_networking" {
  service                    = "servicenetworking.googleapis.com"
  project                    = var.project_name
  disable_dependent_services = true
}

resource "google_project_service" "cloud_run" {
  service = "run.googleapis.com"
    project                    = var.project_name
  disable_dependent_services = true
}

resource "google_project_service" "secret_manager" {
  service = "secretmanager.googleapis.com"
      project                    = var.project_name
  disable_dependent_services = true
}
