resource "google_compute_network" "private_network" {
  name                    = "private-network"
  description             = "CloudSQLのプライベートネットワーク接続用"
  auto_create_subnetworks = false
}

output "cloud_sql_instance_private_ip" {
  value = google_compute_network.private_network.id
}

resource "google_compute_global_address" "private_ip_address" {
  name          = "private-ip-address"
  description   = "CloudSQLのプライベートネットワーク接続用"
  address_type  = "INTERNAL" # プライベート(内部)IPアドレス
  ip_version    = "IPV4"
  network       = google_compute_network.private_network.id
  purpose       = "VPC_PEERING"
  prefix_length = 20
}

resource "google_service_networking_connection" "private_connection" {
  network                 = google_compute_network.private_network.id
  service                 = google_project_service.service_networking.service
  reserved_peering_ranges = [google_compute_global_address.private_ip_address.name] # IPアドレス範囲を指定
}

output "clout_sql_instance_private_connection" {
  value = google_service_networking_connection.private_connection
}