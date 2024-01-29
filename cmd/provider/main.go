package main

import (
	"context"
	"flag"
	"log"

	"github.com/berquerant/terraform-slack/internal/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

var (
	version string = "dev"
)

func main() {
	var debug bool
	flag.BoolVar(&debug, "debug", false, "")
	flag.Parse()

	opts := providerserver.ServeOpts{
		// TODO: Update Address
		Address: "registry.terraform.io/hashicorp/slack",
		Debug:   debug,
	}

	if err := providerserver.Serve(context.Background(), provider.New(version), opts); err != nil {
		log.Fatal(err)
	}
}
