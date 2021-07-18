# NOTE: CDのsetup関連のためsetup.tfと同様手動でapplyする

resource "google_service_account" "piped" {
  description = "PipedをClooudRunで起動させるためのもの"
  account_id = "piped"
  display_name = "piped"
}

resource "google_project_iam_custom_role" "piped" {
  description = "PipedをClooudRunで起動させるためのもの"
  role_id = "piped_role"
  title = "Piped role"
  permissions = [
    # サービスアカウント
    "iam.serviceAccounts.actAs",
    "iam.serviceAccounts.get",
    "iam.serviceAccounts.list",
    "resourcemanager.projects.get",
  ]
}