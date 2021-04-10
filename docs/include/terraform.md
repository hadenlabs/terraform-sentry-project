<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 0.13 |
| <a name="requirement_sentry"></a> [sentry](#requirement\_sentry) | ~> 0.6.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_sentry"></a> [sentry](#provider\_sentry) | ~> 0.6.0 |

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
| <a name="input_name"></a> [name](#input\_name) | The name of the project. | `string` | n/a | yes |
| <a name="input_organization"></a> [organization](#input\_organization) | The name of the organization. | `string` | n/a | yes |
| <a name="input_platform"></a> [platform](#input\_platform) | The name of platform. | `string` | n/a | yes |
| <a name="input_team"></a> [team](#input\_team) | The name of team. | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_dsn_public"></a> [dsn\_public](#output\_dsn\_public) | output dsn public sentry |
| <a name="output_project"></a> [project](#output\_project) | output project name and url |
<!-- END_TF_DOCS -->