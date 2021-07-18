# NOTE: CDのsetup関連のためsetup.tfと同様手動でapplyする

resource "google_service_account" "piped" {
  description  = "PipedをClooudRunで起動させるためのもの"
  account_id   = "piped-account"
  display_name = "piped"
}

resource "google_project_iam_custom_role" "piped" {
  description = "PipedをClooudRunで起動させるためのもの"
  role_id     = "piped_role"
  title       = "Piped role"
  permissions = [
    # サービスアカウント
    "iam.serviceAccounts.actAs",
    "iam.serviceAccounts.get",
    "iam.serviceAccounts.list",
    "resourcemanager.projects.get",
  ]
}

resource "google_project_iam_member" "piped" {
  project = var.project_name
  role    = "projects/${var.project_name}/roles/${google_project_iam_custom_role.piped.role_id}"
  member  = "serviceAccount:${google_service_account.piped.email}"
}

resource "google_cloud_run_service" "piped" {
  name                       = "piped"
  project                    = var.project_name
  location                   = var.project_region
  autogenerate_revision_name = false
  template {
    metadata {
      name = "piped"
      annotations = {
        "autoscaling.knative.dev/maxScale"  = "1" # This must be 1
        "autoscaling.knative.dev/minScale"  = "1" # This must be 1
        "run.googleapis.com/ingress"        = "internal"
        "run.googleapis.com/ingress-status" = "internal"
      }
    }
    spec {
      container_concurrency = 1 # This must be 1
      containers {
        image = "gcr.io/pipecd/piped:v0.12.0"
        args = [
          "piped",
          "--metrics=true",
          # TODO: secretを作成してconfigを読み込むようにする
          # "--config-file=/etc/piped-config/config.yaml",
        ]
        ports {
          container_port = 9085
        }
        resources {
          limits = {
            "cpu"    = "1000m"
            "memory" = "512Mi"
          }
        }
      }
    }
  }
}