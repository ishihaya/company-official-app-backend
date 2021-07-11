locals {
  # TODO: 変数の値を入れる
  project_id = ""
  project_region = "asia-northeast1"
  gcs_backet_name = ""
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
  # TODO: 必要な変数を入れる
}