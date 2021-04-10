output "instance" {
  description = "output instance of project"
  value       = sentry_project.this
}

output "id" {
  description = "output project id"
  value       = sentry_project.this.id
}

output "dsn_public" {
  description = "output dsn public sentry"
  value       = sentry_key.this.dsn_public
}