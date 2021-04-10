resource "sentry_project" "this" {
  organization = var.organization
  name         = lower(var.name)
  team         = var.team
  slug         = var.slug
  platform     = var.platform
  resolve_age  = 720
}

resource "sentry_key" "this" {
  depends_on = [
    sentry_project.this
  ]

  organization = var.organization
  project      = var.name
  name         = var.name
}
