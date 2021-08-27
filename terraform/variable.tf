variable "project_name" { type = string }
variable "project_region" { type = string }
variable "tfstate_bucket_name" { type = string }
variable "cloud_sql_instance" {
  type = object({
    connection_name = string
    name            = string
  })
}
variable "mysql_root_password" { type = string }
variable "mysql_user" { type = string }
variable "mysql_password" { type = string }
variable "mysql_migration_ro_password" { type = string }
variable "cloud_run_name" { type = string }
variable "cloud_sql_instance_connection_name" { type = string }
variable "container_image_name" { type = string }
variable "app_version" { type = string }
variable "mysql_database" { type = string }
variable "app_env" { type = string }