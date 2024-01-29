# experimental-terraform-slack

Test to write a terraform provider using swagger.

## Use the Provider locally

Write `~/.terraformrc`.

```
provider_installation {

  dev_overrides {
      "berquerant/slack" = "/path/to/experimental-terraform-slack/tmp"
  }

  # For all other providers, install them directly from their origin provider
  # registries as normal. If you omit this, Terraform will _only_ use
  # the dev_overrides block, and so no other providers will be available.
  direct {}
}
```

Run API server.

``` shell
make all
docker compose up -d
```

Try to get a list of channels.

``` shell
export SLACK_TOKEN="..."
cd examples/provider-install-verification
terraform plan
```
