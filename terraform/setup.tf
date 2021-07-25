# 初期設定
# 初回は権限のあるユーザーでapplyする
# またはコンソール上で手動で操作してimportする

# NOTE: tfstate管理用のこのバケットは初回はコンソール上で作る必要がある
# terraform import module.root.google_storage_bucket.manage_tfstate PROJECT_NAME/CLOUD_STORAGE_BUCKET_NAME
resource "google_storage_bucket" "manage_tfstate" {
  name                        = var.tfstate_bucket_name
  project                     = var.project_name
  location                    = var.project_region
  default_event_based_hold    = false
  force_destroy               = false
  requester_pays              = false
  uniform_bucket_level_access = true
}

# NOTE: credentialはコンソール上で作成
resource "google_service_account" "github_actions_read_write_service_account" {
  description  = "Github Actionsを使用してGCPのリソースに読み書きが可能なサービスアカウント"
  account_id   = "github-actions-read-write"
  display_name = "github-actions-read-write"
}

resource "google_project_iam_custom_role" "github_actions_read_write_role" {
  description = "Github Actionsを使用してGCPのリソースに読み書きが可能なロール"
  role_id     = "github_actions_read_write_role"
  title       = "Github Actions Read Write"
  # NOTE: 必要な権限をここに全て追加していく
  permissions = [
    # gcs
    "storage.objects.create",
    "storage.objects.delete",
    "storage.objects.get",
    "storage.objects.list",
    "storage.buckets.get",
    # NOTE: 初回gcrにpushする際にこの権限がないと403が出る
    "storage.buckets.create",
    # service account
    "iam.serviceAccounts.actAs",
    "iam.serviceAccounts.get",
    "iam.serviceAccounts.list",
    "resourcemanager.projects.get",
    "iam.serviceAccounts.create",
    "iam.serviceAccounts.update",
    "iam.serviceAccounts.delete",
    # googleapi
    "serviceusage.services.enable",
    "serviceusage.services.disable",
    "serviceusage.services.get",
    "serviceusage.services.list",
    # role
    "iam.roles.create",
    "iam.roles.delete",
    "iam.roles.get",
    "iam.roles.list",
    "iam.roles.undelete",
    "iam.roles.update",
    # iam # 権限の付与を可能にする
    "resourcemanager.projects.setIamPolicy",
    "resourcemanager.projects.getIamPolicy",
  ]
}

resource "google_project_iam_member" "github_actions_read_write_iam" {
  project = var.project_name
  role    = "projects/${var.project_name}/roles/${google_project_iam_custom_role.github_actions_read_write_role.role_id}"
  member  = "serviceAccount:${google_service_account.github_actions_read_write_service_account.email}"
}

# NOTE: credentialはコンソール上で作成
resource "google_service_account" "github_actions_read_only_service_account" {
  description  = "Github Actionsを使用してGCPのリソースに読み込みのみ可能なサービスアカウント"
  account_id   = "github-actions-read-only"
  display_name = "github-actions-read-only"
}

resource "google_project_iam_custom_role" "github_actions_read_only_role" {
  description = "Github Actionsを使用してGCPのリソースに読み込みのみ可能なロール"
  role_id     = "github_actions_read_only_role"
  title       = "Github Actions Read Only"
  # NOTE: 必要な権限をここに全て追加していく
  permissions = [
    # gcs
    "storage.objects.create",
    "storage.objects.delete",
    "storage.objects.get",
    "storage.objects.list",
    "storage.buckets.get",
    # service account
    "iam.serviceAccounts.get",
    "iam.serviceAccounts.list",
    "resourcemanager.projects.get",
    # googleapi
    "serviceusage.services.get",
    "serviceusage.services.list",
    # role
    "iam.roles.get",
    "iam.roles.list",
    # iam
    "resourcemanager.projects.getIamPolicy",
  ]
}

# terraform import module.root.google_project_iam_member.github_actions_read_only_iam "projects/PROJECT_NAME/roles/github_actions_read_only_role serviceAccount:ACCOUNT_EMAIL"
resource "google_project_iam_member" "github_actions_read_only_iam" {
  project = var.project_name
  role    = "projects/${var.project_name}/roles/${google_project_iam_custom_role.github_actions_read_only_role.role_id}"
  member  = "serviceAccount:${google_service_account.github_actions_read_only_service_account.email}"
}