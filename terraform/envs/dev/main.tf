terraform {
  required_version = "~> 1.0.0"
  # NOTE: tfstateをgcsで管理する
  backend "gcs" {
    bucket = "tfstate-coa-dev"
    prefix = "." # ディレクトリ
  }
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "3.74.0"
    }
  }
}

locals {
  project_id     = "company-official-app-dev"
  project_region = "asia-northeast1"
}

provider "google" {
  project = local.project_id
  region  = local.project_region
  zone    = "${local.project_region}-a"
}

# secrets
variable "mysql_root_password" {}
variable "mysql_password" {}
variable "mysql_migration_ro_password" {}

module "root" {
  source                      = "./../.."
  project_name                = local.project_id
  project_region              = local.project_region
  tfstate_bucket_name         = "tfstate-coa-dev"
  cloud_sql_instance          = google_sql_database_instance.app_mysql_instance
  mysql_root_password         = var.mysql_root_password
  mysql_user                  = "app"
  mysql_password              = var.mysql_password
  mysql_migration_ro_password = var.mysql_migration_ro_password
}