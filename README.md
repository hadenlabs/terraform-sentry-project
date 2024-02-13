<!--


  ** DO NOT EDIT THIS FILE
  **
  ** 1) Make all changes to `provision/generator/README.yaml`
  ** 2) Run`task readme` to rebuild this file.
  **
  ** (We maintain HUNDREDS of open source projects. This is how we maintain our sanity.)
  **


  -->

[![Latest Release](https://img.shields.io/github/release/hadenlabs/terraform-sentry-project)](https://github.com/hadenlabs/terraform-sentry-project/releases) [![Lint](https://img.shields.io/github/workflow/status/hadenlabs/terraform-sentry-project/lint-code)](https://github.com/hadenlabs/terraform-sentry-project/actions?workflow=lint-code) [![CI](https://img.shields.io/github/workflow/status/hadenlabs/terraform-sentry-project/ci)](https://github.com/hadenlabs/terraform-sentry-project/actions?workflow=ci) [![Test](https://img.shields.io/github/workflow/status/hadenlabs/terraform-sentry-project/test)](https://github.com/hadenlabs/terraform-sentry-project/actions?workflow=test) [![pre-commit](https://img.shields.io/badge/pre--commit-enabled-brightgreen?logo=pre-commit&logoColor=white)](https://github.com/pre-commit/pre-commit) [![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-yellow)](https://conventionalcommits.org) [![KeepAChangelog](https://img.shields.io/badge/changelog-Keep%20a%20Changelog%20v1.0.0-orange)](https://keepachangelog.com) [![Terraform Version](https://img.shields.io/badge/terraform-1.x%20|%200.15%20|%200.14%20|%200.13%20|%200.12.20+-623CE4.svg?logo=terraform)](https://github.com/hashicorp/terraform/releases)

# terraform-sentry-project

Terraform module to provision an sentry project.

## Requirements

This is a list of plugins that need to be installed previously to enjoy all the goodies of this configuration:

- [gomplate](https://github.com/hairyhenderson/gomplate)
- [terraform](https://github.com/hashicorp/terraform)
- [python](https://www.python.org)
- [taskfile](https://github.com/go-task/task)

## Usage

# How to use this project

```hcl
  module "main" {
    source = "hadenlabs/project/sentry"
    version = "0.1.0"

    providers = {
      sentry = sentry
    }

    name    = "terraform-sentry-project"
    organization    = "slug-organization"
    team    = "slug-team"
    platform = "go"
  }
```

Full working examples can be found in [examples](./examples) folder.

## Examples

<!-- Space: Projects -->
<!-- Parent: TerraformSentryProject -->
<!-- Title: Examples TerraformSentryProject -->
<!-- Label: Examples -->
<!-- Include: ./../disclaimer.md -->
<!-- Include: ac:toc -->

### Common

```hcl
  module "main" {
    source = "hadenlabs/project/sentry"
    version = "0.1.0"

    providers = {
      sentry = sentry
    }

    name    = "terraform-sentry-project"
    organization    = "slug-organization"
    team    = "slug-team"
    platform = "go"
  }
```

 <!-- markdown-link-check-disable -->
<!-- BEGIN_TF_DOCS -->

## Requirements

| Name      | Version |
| --------- | ------- |
| terraform | >= 0.13 |
| sentry    | >=0.6.0 |

## Providers

| Name   | Version |
| ------ | ------- |
| sentry | >=0.6.0 |

## Modules

No modules.

## Resources

| Name | Type |
| --- | --- |
| [sentry_key.this](https://registry.terraform.io/providers/jianyuan/sentry/latest/docs/resources/key) | resource |
| [sentry_project.this](https://registry.terraform.io/providers/jianyuan/sentry/latest/docs/resources/project) | resource |

## Inputs

| Name         | Description                   | Type     | Default | Required |
| ------------ | ----------------------------- | -------- | ------- | :------: |
| name         | The name of the project.      | `string` | n/a     |   yes    |
| organization | The name of the organization. | `string` | n/a     |   yes    |
| platform     | The name of platform.         | `string` | n/a     |   yes    |
| team         | The name of team.             | `string` | n/a     |   yes    |

## Outputs

| Name       | Description                |
| ---------- | -------------------------- |
| dsn_public | output dsn public sentry   |
| id         | output project id          |
| instance   | output instance of project |

<!-- END_TF_DOCS -->
<!-- markdown-link-check-enable -->

## Help

**Got a question?**

File a GitHub [issue](https://github.com/hadenlabs/terraform-sentry-project/issues).

## Contributing

See [Contributing](./docs/contributing.md).

## Module Versioning

This Module follows the principles of [Semantic Versioning (SemVer)](https://semver.org/).

Using the given version number of `MAJOR.MINOR.PATCH`, we apply the following constructs:

1. Use the `MAJOR` version for incompatible changes.
1. Use the `MINOR` version when adding functionality in a backwards compatible manner.
1. Use the `PATCH` version when introducing backwards compatible bug fixes.

### Backwards compatibility in `0.0.z` and `0.y.z` version

- In the context of initial development, backwards compatibility in versions `0.0.z` is **not guaranteed** when `z` is increased. (Initial development)
- In the context of pre-release, backwards compatibility in versions `0.y.z` is **not guaranteed** when `y` is increased. (Pre-release)

## Copyright

Copyright © 2018-2024 [Hadenlabs](https://hadenlabs.com)

## Trademarks

All other trademarks referenced herein are the property of their respective owners.

## License

The code and styles are licensed under the LGPL-3.0 license [See project license.](LICENSE).

## Don't forget to 🌟 Star 🌟 the repo if you like terraform-sentry-project

[Your feedback is appreciated](https://github.com/hadenlabs/terraform-sentry-project/issues)
