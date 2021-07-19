resource "google_container_cluster" "coa" {
  name             = "coa-cluster"
  location         = var.project_region
  enable_autopilot = true
}