package slackapi

import "github.com/slack-go/slack"

func New(token string) API {
	client := slack.New(token)
	return &api{
		client: client,
	}
}

type api struct {
	client *slack.Client
}

var (
	// Ensure that api implements API.
	_ API = &api{}
)

type API interface {
	ChannelAPI
}
