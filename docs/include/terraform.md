<!-- markdown-link-check-disable -->
<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| terraform | >= 0.13 |
| sentry | >=0.6.0 |

## Providers

| Name | Version |
|------|---------|
| sentry | >=0.6.0 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [sentry_key.this](https://registry.terraform.io/providers/jianyuan/sentry/latest/docs/resources/key) | resource |
| [sentry_project.this](https://registry.terraform.io/providers/jianyuan/sentry/latest/docs/resources/project) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| name | The name of the project. | `string` | n/a | yes |
| organization | The name of the organization. | `string` | n/a | yes |
| platform | The name of platform. | `string` | n/a | yes |
| team | The name of team. | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| dsn\_public | output dsn public sentry |
| id | output project id |
| instance | output instance of project |
<!-- END_TF_DOCS -->
<!-- markdown-link-check-enable -->