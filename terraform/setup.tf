# 初期設定
# 初回はコンソール上で手動で操作してimportする

# NOTE: credentialはコンソール上で作成
# terraform import module.root.google_service_account.github_actions_read_write_service_account projects/PROJECT_NAME/serviceAccounts/ACCOUNT_MAIL
resource "google_service_account" "github_actions_read_write_service_account" {
    description = "Github Actionsを使用してGCPのリソースに読み書きが可能なサービスアカウント"
    account_id = "github-actions-read-write"
    display_name = "github-actions-read-write"
}

# terraform import module.root.google_project_iam_custom_role.github_actions_read_write_role projects/PROJECT_NAME/roles/github_actions_read_write_role
resource "google_project_iam_custom_role" "github_actions_read_write_role" {
    description = "Github Actionsを使用してGCPのリソースに読み書きが可能なロール"
    role_id = "github_actions_read_write_role"
    title = "Github Actions Read Write"
    # NOTE: 必要な権限をここに全て追加していく
    permissions = [ 
        # gcs
        "storage.objects.create",
        "storage.objects.delete",
        "storage.objects.get",
        "storage.objects.list",
        "storage.buckets.get",
    ]
}