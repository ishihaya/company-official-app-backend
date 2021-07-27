resource "google_sql_database_instance" "app_mysql_instance" {
  # NOTE: インスタンス再生成などの場合はsuffixの数値を増やして対応する
  name             = "app-mysql-01"
  database_version = "MYSQL_8_0"
  region           = local.project_region
  depends_on       = [module.root.clout_sql_instance_private_connection]
  settings {
    # NOTE: 実用レベルのプロジェクトや環境を除き冗長化を避けて料金を抑える
    tier              = "db-f1-micro"
    availability_type = "ZONAL"
    disk_autoresize   = false
    disk_size         = 10
    disk_type         = "PD_HDD"
    activation_policy = "ALWAYS" # 停止する場合は"NEVER"
    ip_configuration {
      ipv4_enabled = true # Cloud SQL Proxy のみ使用可能
      # NOTE: プライベートIPは現在使用していないが必要になった際にDB作り直しが必要なので事前に用意しておく
      private_network = module.root.cloud_sql_instance_private_ip # 作成したVPCネットワークに接続可能
    }
    backup_configuration {
      enabled                        = true
      binary_log_enabled             = true
      start_time                     = "00:00"
      transaction_log_retention_days = 1 # backup 保管日数
    }
    maintenance_window {
      day  = 1  # Monday
      hour = 21 # 06:00 JST
    }
    # kamipo traditional
    database_flags {
      name  = "sql_mode"
      value = "TRADITIONAL,NO_AUTO_VALUE_ON_ZERO,ONLY_FULL_GROUP_BY"
    }
    # https://cloud.google.com/sql/docs/mysql/flags?hl=ja
    database_flags {
      name  = "slow_query_log"
      value = "on"
    }
    database_flags {
      name  = "log_output"
      value = "FILE"
    }
    database_flags {
      name  = "innodb_lock_wait_timeout"
      value = "10"
    }
    database_flags {
      name  = "long_query_time"
      value = "1"
    }
  }
}