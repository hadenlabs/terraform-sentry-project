# ---------------------------------------------------------------------------------------------------------------------
# ENVIRONMENT VARIABLES
# Define these secrets as environment variables.
# ---------------------------------------------------------------------------------------------------------------------

# SENTRY_TOKEN

# ---------------------------------------------------------------------------------------------------------------------
# REQUIRED VARIABLES
# These variables must be set when using this module.
# ---------------------------------------------------------------------------------------------------------------------

# ---------------------------------------------------------------------------------------------------------------------
# OPTIONAL VARIABLES
# These variables have defaults, but may be overridden.
# ---------------------------------------------------------------------------------------------------------------------
variable "name" {
  type        = string
  description = "The name of the project."
}

variable "organization" {
  type        = string
  description = "The name of the organization."
}

variable "team" {
  type        = string
  description = "The name of team."
}

variable "platform" {
  type        = string
  description = "The name of platform."
}
