module "main" {
  source = "../.."

  name         = var.name
  organization = var.organization
  team         = var.team
  platform     = var.platform
}
