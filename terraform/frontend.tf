resource "google_service_account" "firebase_hosting" {
  account_id   = "firebase-hosting"
  display_name = "firebase-hosting"
  description  = "フロントエンドサービスをfirebase hostingでデプロイする用"
}

resource "google_project_iam_member" "firebase_hosting" {
}