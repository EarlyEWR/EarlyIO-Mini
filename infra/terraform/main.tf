resource "kubernetes_namespace" "ns" {
  metadata {
    name = var.namespace
  }
}

resource "kubernetes_config_map" "example" {
  metadata {
    name      = "learning-config"
    namespace = var.namespace
  }
  data = {
    "WELCOME" = "Terraform + Kubernetes + Go"
  }
}
