package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/berquerant/terraform-slack/swagger/client"
	"github.com/go-openapi/runtime"
	transport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/spf13/cobra"
)

func main() {
	if execute() != nil {
		os.Exit(1)
	}
}

var (
	rootCmd = &cobra.Command{
		Use: "slack-api-for-terraform-provider-client",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			cmd.SetOut(os.Stdout)
		},
	}
)

func execute() error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		log.Printf("execute error %v", err)
		return err
	}
	return nil
}

func init() {
	rootCmd.PersistentFlags().String("host", "127.0.0.1:8030", "")
	rootCmd.PersistentFlags().String("key", "", "API Key")
}

func newClient(cmd *cobra.Command) *client.SlackAPIForTerraformProvider {
	host, _ := cmd.Flags().GetString("host")
	cfg := client.DefaultTransportConfig().WithHost(host)
	return client.NewHTTPClientWithConfig(strfmt.Default, cfg)
}

func getAPIKey(cmd *cobra.Command) string {
	x, _ := cmd.Flags().GetString("key")
	return x
}

type config struct {
	c   *client.SlackAPIForTerraformProvider
	key string
}

func getConfig(cmd *cobra.Command) *config {
	return &config{
		c:   newClient(cmd),
		key: getAPIKey(cmd),
	}
}

func (c *config) authInfo() runtime.ClientAuthInfoWriter {
	return transport.APIKeyAuth("X-API-Key", "header", c.key)
}
