resource "google_service_account" "firebase_hosting" {
  account_id   = "firebase-hosting"
  display_name = "firebase-hosting"
  description  = "フロントエンドサービスをfirebase hostingでデプロイする用"
}

resource "google_project_iam_member" "firebase_hosting_admin" {
  project = var.project_name
  member  = "serviceAccount:${google_service_account.firebase_hosting.email}"
  role    = "roles/firebasehosting.admin"
}

resource "google_project_iam_member" "firebase_hosting_api_keys_viewer" {
  project = var.project_name
  member  = "serviceAccount:${google_service_account.firebase_hosting.email}"
  role    = "roles/serviceusage.apiKeysViewer"
}