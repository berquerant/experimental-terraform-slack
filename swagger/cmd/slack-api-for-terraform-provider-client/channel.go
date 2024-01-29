package main

import (
	"encoding/json"

	"github.com/berquerant/terraform-slack/swagger/client/channel"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(channelCmd)
	initChannelList()
	initChannelGet()
	initChannelCreate()
	initChannelUpdate()
}

var channelCmd = &cobra.Command{
	Use: "channel",
}

func initChannelList() {
	channelCmd.AddCommand(channelListCmd)
	channelListCmd.Flags().String("team", "", "team id")
}

var channelListCmd = &cobra.Command{
	Use: "list",
	RunE: func(cmd *cobra.Command, args []string) error {
		c := getConfig(cmd)
		teamID, _ := cmd.Flags().GetString("team")
		params := channel.NewFetchChannelsParams().
			WithTeamID(&teamID)
		r, err := c.c.Channel.FetchChannels(params, c.authInfo())
		if err != nil {
			return err
		}
		v, _ := json.Marshal(r)
		cmd.Printf("%s\n", v)
		return nil
	},
}

func initChannelGet() {
	channelCmd.AddCommand(channelGetCmd)
	channelGetCmd.Flags().String("id", "", "channel id")
}

var channelGetCmd = &cobra.Command{
	Use: "get",
	RunE: func(cmd *cobra.Command, args []string) error {
		c := getConfig(cmd)
		id, _ := cmd.Flags().GetString("id")
		params := channel.NewFetchChannelParams().
			WithID(id)
		r, err := c.c.Channel.FetchChannel(params, c.authInfo())
		if err != nil {
			return err
		}
		v, _ := json.Marshal(r)
		cmd.Printf("%s\n", v)
		return nil
	},
}

func initChannelCreate() {
	channelCmd.AddCommand(channelCreateCmd)
	channelCreateCmd.Flags().String("name", "", "channel name")
	channelCreateCmd.Flags().Bool("private", false, "if true, private channel")
	channelCreateCmd.Flags().String("team", "", "channel team")
}

var channelCreateCmd = &cobra.Command{
	Use: "create",
	RunE: func(cmd *cobra.Command, args []string) error {
		c := getConfig(cmd)
		name, _ := cmd.Flags().GetString("name")
		isPrivate, _ := cmd.Flags().GetBool("private")
		teamID, _ := cmd.Flags().GetString("team")
		body := channel.CreateChannelBody{
			Name:      &name,
			IsPrivate: &isPrivate,
			TeamID:    &teamID,
		}
		params := channel.NewCreateChannelParams().WithChannel(body)
		r, err := c.c.Channel.CreateChannel(params, c.authInfo())
		if err != nil {
			return err
		}
		v, _ := json.Marshal(r)
		cmd.Printf("%s\n", v)
		return nil
	},
}

func initChannelUpdate() {
	channelCmd.AddCommand(channelUpdateCmd)
	channelUpdateCmd.Flags().String("id", "", "channel id")
	channelUpdateCmd.Flags().String("name", "", "channel name")
	channelUpdateCmd.Flags().Bool("archive", false, "if true, archive channel")
}

var channelUpdateCmd = &cobra.Command{
	Use: "update",
	RunE: func(cmd *cobra.Command, args []string) error {
		c := getConfig(cmd)
		id, _ := cmd.Flags().GetString("id")
		name, _ := cmd.Flags().GetString("name")
		archive, _ := cmd.Flags().GetBool("archive")
		body := channel.UpdateChannelBody{
			Name:       name,
			IsArchived: archive,
		}
		params := channel.NewUpdateChannelParams().
			WithID(id).
			WithChannel(body)
		r, err := c.c.Channel.UpdateChannel(params, c.authInfo())
		if err != nil {
			return err
		}
		v, _ := json.Marshal(r)
		cmd.Printf("%s\n", v)
		return nil
	},
}
