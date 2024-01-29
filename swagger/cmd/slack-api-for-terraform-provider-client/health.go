package main

import (
	"encoding/json"

	"github.com/berquerant/terraform-slack/swagger/client/health"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(healthCmd)
}

var healthCmd = &cobra.Command{
	Use: "health",
	RunE: func(cmd *cobra.Command, args []string) error {
		c := getConfig(cmd)
		params := health.NewCheckAliveParams()
		r, err := c.c.Health.CheckAlive(params, c.authInfo())
		if err != nil {
			return err
		}
		v, _ := json.Marshal(r)
		cmd.Printf("%s\n", v)
		return nil
	},
}
