resource "google_compute_instance" "apps" {
  count        = 4
  name         = "apps-${count.index + 1}"
  machine_type = "f1-micro"

  boot_disk {
    initialize_params {
      image = "ubuntu-os-cloud/ubuntu-1804-lts"
    }
  }

  network_interface {
    network = "main"

    subnetwork = "private"
    access_config {
    }
  }
}

resource "random_id" "db_name_suffix" {
  byte_length = 4
}

locals {
  onprem = [
    "86.178.136.146"
  ]
}

resource "google_sql_user" "admin" {
  name     = "admin"
  password = var.db_password
  instance = google_sql_database_instance.mysql-instance.name
}

resource "google_sql_database_instance" "mysql-instance" {
  name             = var.db_instance_name
  database_version = var.database_version
  region           = var.region

  settings {
    tier = "db-f1-micro"
    ip_configuration {
      ipv4_enabled = true

      dynamic "authorized_networks" {
        for_each = google_compute_instance.apps
        iterator = apps
        content {
          name  = apps.value.name
          value = apps.value.network_interface.0.access_config.0.nat_ip
        }
      }

      dynamic "authorized_networks" {
        for_each = local.onprem
        iterator = onprem
        content {
          name  = "onprem-${onprem.key}"
          value = onprem.value
        }
      }
    }
  }
}

resource "google_sql_database" "meetup" {
  instance = var.db_instance_name
  name     = var.db_name
}
