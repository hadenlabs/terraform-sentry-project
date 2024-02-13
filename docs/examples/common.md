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
    version = "0.1.1"

    providers = {
      sentry = sentry
    }

    name    = "terraform-sentry-project"
    organization    = "slug-organization"
    team    = "slug-team"
    platform = "go"
  }
```
