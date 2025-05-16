terraform {
  required_providers {
    oci = {
      source  = "terraform.local/local/oci"
      version = "1.0.0"
    }
    kubernetes = {
      source = "kubernetes"
    }
  }
}

provider "oci" {
}

provider "kubernetes" {
  # cluster_ca_certificate = base64decode(yamldecode(data.oci_containerengine_cluster_kube_config.kube_config.content)["clusters"][0]["cluster"]["certificate-authority-data"])
  host                   = "https://158.180.53.1:6443"
  token                  = data.oci_containerengine_cluster_auth.cluster.token
  insecure               = true
}


data "oci_containerengine_cluster_auth" "cluster" {
    cluster_id = "test"
}


data "kubernetes_all_namespaces" "allns" {
}


output "auth" {
  value = [
    data.oci_containerengine_cluster_auth.cluster
  ]
}


output "all-ns" {
  value = [
    data.kubernetes_all_namespaces.allns.namespaces,
  ]
}
