// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {}
variable "user_ocid" {}
variable "fingerprint" {}
variable "private_key_path" {}
variable "region" {}
variable "compartment_id" {}

variable "organization_subscription_subscription_ids" {
  default = "subscriptionIds"
}

provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

data "oci_osub_organization_subscription_organization_subscriptions" "test_organization_subscriptions" {
  #Required
  compartment_id   = var.compartment_id
  subscription_ids = var.organization_subscription_subscription_ids

  #Optional
  x_one_origin_region = var.region
}
