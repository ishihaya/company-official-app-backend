locals {
  project_id = "company-official-app-dev"
  project_region = "asia-northeast1"
  gcs_backet_name = "tfstate-coa-dev"
}

terraform {
  required_version = "~> 1.0.0"
  # NOTE: tfstateをgcsで管理する
  backend "gcs" {
    bucket = local.gcs_backet_name
    prefix = "." # ディレクトリ
  }
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "3.68.0"
    }
  }
}

provider "google" {
  project = local.project_id
  region  = local.project_region
  zone    = "${local.project_region}-a"
}

module "root" {
  source = "./../.."
}