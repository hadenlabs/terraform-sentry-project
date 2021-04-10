output "project" {
  description = "output project name and url"
  value = {
    name = github_organization_project.project.name
    url  = github_organization_project.project.url
  }
}

output "dsn_public" {
  description = "output dsn public sentry"
  value       = sentry_key.this.dsn_public
}