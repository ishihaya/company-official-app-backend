resource "google_service_account" "github_actions_mysql_client" {
  description  = "Github ActionsがMySQLに接続する用"
  account_id   = "github-actions-mysql-client"
  display_name = "github-actions-mysql-client"
}

resource "google_project_iam_member" "github_actions_mysql_client" {
  project = var.project_name
  role    = "roles/cloudsql.client"
  member  = "serviceAccount:${google_service_account.github_actions_mysql_client.email}"
}

resource "google_sql_database" "app_mysql_database" {
  name       = "app-mysql-database"
  instance   = var.cloud_sql_instance.name
  depends_on = [var.cloud_sql_instance]
  charset    = "utf8mb4"
  collation  = "utf8mb4_bin"
}

# rootユーザー
resource "google_sql_user" "app_mysql_root_user" {
  instance   = var.cloud_sql_instance.name
  depends_on = [var.cloud_sql_instance]
  name     = "root"
  password = var.mysql_root_password
}

# アプリケーションで使用するユーザー
resource "google_sql_user" "app_mysql_user" {
  instance   = var.cloud_sql_instance.name
  depends_on = [var.cloud_sql_instance]
  name     = var.mysql_user
  password = var.mysql_password
}

# migrate dry-run 使用時にテーブルとカラム情報を読み込むユーザー
resource "google_sql_user" "app_mysql_migration_ro_user" {
  instance   = var.cloud_sql_instance.name
  depends_on = [var.cloud_sql_instance]
  name     = "migration_ro"
  password = var.mysql_migration_ro_password
}