terraform {
  required_providers {
    slack = {
      source = "berquerant/slack"
    }
  }
}

provider "slack" {
  endpoint = "127.0.0.1:8030"
}

resource "slack_channel" "example" {
  name = "test-example-new-tf"
}

data "slack_channels" "example" {}

output "channels" {
  value = data.slack_channels.example
}
