variable "bucket_name" {
  type        = string
  description = "Bucket Name"
  default     = "liusha-k8s-nonprod-test"
}

variable "bucket_location" {
  type    = string
  default = "europe-west2"
}

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
