terraform {
  required_providers {
    oci = {
      source  = "terraform.local/local/oci"
      version = "1.0.0"
    }
  }
}

provider "oci" {
}

provider "kubernetes" {
  cluster_ca_certificate = base64decode(yamldecode(data.oci_containerengine_cluster_kube_config.kube_config.content)["clusters"][0]["cluster"]["certificate-authority-data"])
  host                   = yamldecode(data.oci_containerengine_cluster_kube_config.kube_config.content)["clusters"][0]["cluster"]["server"]
  token                  = data.oci_containerengine_cluster_auth.cluster.token
}

data "oci_containerengine_cluster_auth" "cluster" {
    cluster_id = "ocid1.cluster.oc1.eu-frankfurt-1.aaaaaaaainvo5g44icgg4kgf7ndyxpsv6xkrsw3egtmpeirlgckswsk3r2xq"
}

data "oci_containerengine_cluster_kube_config" "kube_config" {
    cluster_id = "ocid1.cluster.oc1.eu-frankfurt-1.aaaaaaaainvo5g44icgg4kgf7ndyxpsv6xkrsw3egtmpeirlgckswsk3r2xq"
}


resource "time_sleep" "wait" {
  depends_on = [data.oci_containerengine_cluster_auth.cluster]

  create_duration = "10s"
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
