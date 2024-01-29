terraform {
  required_providers {
    slack = {
      source = "berquerant/slack"
    }
  }
}

provider "slack" {
  endpoint = "127.0.0.1:8030"
  token    = "example"
}
