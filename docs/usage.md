# How to use this project

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

Full working examples can be found in [examples](./examples) folder.
