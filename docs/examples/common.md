### basic

```hcl
  module "main" {
    source = "hadenlabs/project/sentry"
    version = "0.0.0"

    providers = {
      sentry = sentry
    }

    name    = "terraform-sentry-project"
    organization    = "slug-organization"
    team    = "slug-team"
    platform = "go"
  }
```
