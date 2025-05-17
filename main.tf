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



resource "oci_containerengine_cluster" "test_cluster" {
    #Required
    compartment_id = "ocid1.compartment.oc1..aaaaaaaaqi3if6t4n24qyabx5pjzlw6xovcbgugcmatavjvapyq3jfb4diqq"
    kubernetes_version = "v1.31.1"
    name = "oke"
    vcn_id = "ocid1.vcn.oc1.eu-frankfurt-1.amaaaaaawe6j4fqaei4rixhlmbpck3rplwqeze7hlqlpbmty6qphmvxutkha"

    #Optional
    cluster_pod_network_options {
        #Required
        cni_type = "FLANNEL_OVERLAY"
    }
    endpoint_config {

        #Optional
        is_public_ip_enabled = true
        nsg_ids = [
                  "ocid1.networksecuritygroup.oc1.eu-frankfurt-1.aaaaaaaaqtib37pkm2jmardbn4apuzazg4afoj464x3jkwch3lxieqr6k5ya"
                ]
        subnet_id = "ocid1.subnet.oc1.eu-frankfurt-1.aaaaaaaablm5kdm6zrgvkjmlcj7nsvgycdqywebstuz5ygicq7bxdcsdy5kq"
    }
    options {

        #Optional
        add_ons {

            #Optional
            is_kubernetes_dashboard_enabled = false
            is_tiller_enabled = false
        }
        admission_controller_options {

            #Optional
            is_pod_security_policy_enabled = false
        }
        ip_families = ["IPv4"]
        kubernetes_network_config {

            #Optional
            pods_cidr = "10.244.0.0/16"
            services_cidr =  "10.96.0.0/16"
        }
        open_id_connect_token_authentication_config {
            #Required
            is_open_id_connect_auth_enabled = true


            configuration_file = base64encode(yamlencode(
              {
                "apiVersion" = "apiserver.config.k8s.io/v1beta1"
                "kind"       = "AuthenticationConfiguration"
                "jwt" = [
                  {
                    "issuer" = {
                      "url" = "https://dev-hdp8m6tn.eu.auth0.com/",
                      "audiences" = [
                        "1sd1mdsadsasKpEX6dsasdad3yKy1M1263"
                      ],
                      "audienceMatchPolicy" = "MatchAny"
                    }
                    "claimMappings" = {
                      "username" = {
                        "claim"  = "email_address"
                        "prefix" = "oidc:"
                      }
                      "groups" = {
                        "prefix" = "oidc:"
                      }
                    }
                }
                ]
              }
            ))
        }                   
        # open_id_connect_discovery {

        #     #Optional
        #     is_open_id_connect_discovery_enabled = true
        # }
        persistent_volume_config {

            #Optional
            defined_tags ={
                      "oke.role"= "persistent_volume",
                      "oke.state_id"= "uhyedk"
                    }
            freeform_tags = {}
        }
        service_lb_config {

            #Optional
            defined_tags ={
                      "oke.role"= "service_lb",
                      "oke.state_id"= "uhyedk"
                    }
            freeform_tags = {}
        }
        service_lb_subnet_ids = [
                  "ocid1.subnet.oc1.eu-frankfurt-1.aaaaaaaamnyznan4ae2235ppxdep34zax57yjf5tgcqtwkhshwtxgobxdcja"
                ]
    }
    type = "ENHANCED_CLUSTER"
}