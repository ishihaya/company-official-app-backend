resource "google_project_service" "gke" {
  project = var.project_name
}