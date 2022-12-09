provider "google" {
  project = "liusha-k8s-nonprod-test"
  region  = var.region
}

terraform {
  backend "gcs" {
    bucket = "liusha-k8s-nonprod-test"
    prefix = "terraform/state"
  }
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 4.0"
    }
  }
}