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

module "root" {
  source              = "./../.."
  project_name        = local.project_id
  project_region      = local.project_region
  tfstate_bucket_name = "tfstate-coa-dev"
}