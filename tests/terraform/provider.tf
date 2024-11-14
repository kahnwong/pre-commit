provider "helm" {
  registry {
    url      = "oci://localhost:5000"
    username = "username"
    password = "password"
  }
}
