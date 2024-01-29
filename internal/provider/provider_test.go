package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

const (
	providerConfig = `
provider "slack" {
  endpoint = "http://127.0.0.1:8030"
  token = "example"
}
`
)

var (
	testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"slack": providerserver.NewProtocol6WithError(New("test")()),
	}
)
