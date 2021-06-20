output "project" {
  description = "show instance module"
  value       = module.main.instance
}

output "id" {
  description = "show data project"
  value       = module.main.id
}

output "dsn" {
  description = "show dsn sentry."
  value       = module.main.dsn_public
}
