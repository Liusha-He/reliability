resource "google_storage_bucket" "tf-state" {
  name     = var.bucket_name
  location = var.bucket_location
}
