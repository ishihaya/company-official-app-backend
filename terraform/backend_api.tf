resource "google_service_account" "backend_api" {
  account_id   = "backend-api"
  display_name = "backend-api"
  description  = "Cloud Runサービスであるbackend-apiがGCPのリソースへアクセスする用"
}

resource "google_project_iam_custom_role" "backend_api" {
  role_id     = "backend_api_role"
  title       = "Backend API Role to Access GCP Resource"
  description = "Cloud Runサービスであるbackend-apiがGCPのリソースへアクセスする用"
  permissions = [
    # サービスアカウント
    "iam.serviceAccounts.actAs",
    "iam.serviceAccounts.get",
    "iam.serviceAccounts.list",
    "resourcemanager.projects.get",
    # Secret Manager
    "secretmanager.versions.access",
    # Cloud SQL
    "cloudsql.instances.connect",
    "cloudsql.instances.get",
  ]
}

resource "google_project_iam_member" "backend_api" {
  project = var.project_name
  role    = "projects/${var.project_name}/roles/${google_project_iam_custom_role.backend_api.role_id}"
  member  = "serviceAccount:${google_service_account.backend_api.email}"
}

resource "google_cloud_run_service" "company_official_app_backend" {
  provider = google-beta

  name                       = var.cloud_run_name
  project                    = var.project_name
  location                   = var.project_region
  autogenerate_revision_name = true
  metadata {
    annotations = {
      "run.googleapis.com/ingress" = "all"
      # NOTE: Secret Manager(beta)を使う場合初回はこのannotationsを定義しないと動作しない
      # https://github.com/hashicorp/terraform-provider-google/issues/9159
      "run.googleapis.com/launch-stage" = "BETA"
    }
  }
  traffic {
    percent         = 100
    latest_revision = true
  }
  template {
    metadata {
      annotations = {
        "run.googleapis.com/cloudsql-instances" = var.cloud_sql_instance_connection_name
        "autoscaling.knative.dev/minScale"      = "0"
        "autoscaling.knative.dev/maxScale"      = "10"
      }
    }
    spec {
      service_account_name = google_service_account.backend_api.email
      containers {
        image = "${var.container_image_name}:latest"
        env {
          name  = "MYSQL_USER"
          value = var.mysql_user
        }
        env {
          name  = "MYSQL_DATABASE"
          value = var.mysql_database
        }
        env {
          name  = "DB_SOCKET_PATH"
          value = "/cloudsql/${var.cloud_sql_instance_connection_name}"
        }

        # Secret Manager
        env {
          name = "MYSQL_PASSWORD"
          value_from {
            secret_key_ref {
              name = "mysql_password"
              key  = "latest"
            }
          }
        }
      }
    }
  }
}

resource "google_cloud_run_service_iam_member" "backend_api_noauth" {
  service  = google_cloud_run_service.company_official_app_backend.name
  project  = google_cloud_run_service.company_official_app_backend.project
  location = google_cloud_run_service.company_official_app_backend.location
  role     = "roles/run.invoker"
  member   = "allUsers"
}