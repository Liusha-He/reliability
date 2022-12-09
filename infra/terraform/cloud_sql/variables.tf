variable "project_id" {
  type    = string
  default = "liusha-k8s-nonprod-test"
}

variable "region" {
  type    = string
  default = "europe-west2"
}

variable "zone" {
  type    = string
  default = "europe-west2-a"
}

variable "instance_name" {
  type    = string
  default = "master-instance"
}

variable "database_version" {
  type    = string
  default = "mysql_5_7"
}

variable "db_password" {
  type    = string
  default = "test_password"
}

variable "db_instance_name" {
  type    = string
  default = "mysql-instance"
}

variable "db_name" {
  type    = string
  default = "meetup"
}
