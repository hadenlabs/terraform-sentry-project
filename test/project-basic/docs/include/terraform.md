<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >=0.13.0 |
| <a name="requirement_github"></a> [github](#requirement\_github) | >=4.5.0 |

## Providers

No providers.

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_main"></a> [main](#module\_main) | ../.. |  |

## Resources

No resources.

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
| <a name="output_dsn"></a> [dsn](#output\_dsn) | show dsn sentry. |
| <a name="output_id"></a> [id](#output\_id) | show data project |
| <a name="output_project"></a> [project](#output\_project) | show instance module |
<!-- END_TF_DOCS -->