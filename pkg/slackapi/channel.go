package slackapi

import (
	"github.com/slack-go/slack"
)

type GetChannelsRequest struct {
	TeamID string
}

type GetChannelsResponse struct {
	Channels []Channel
}

type Channel struct {
	ID         string
	Name       string
	IsPrivate  bool
	IsArchived bool
}

type ChannelAPI interface {
	GetChannels(req *GetChannelsRequest) (*GetChannelsResponse, error)
	GetChannel(req *GetChannelRequest) (*GetChannelResponse, error)
	CreateChannel(req *CreateChannelRequest) (*CreateChannelResponse, error)
	UpdateChannel(req *UpdateChannelRequest) (*UpdateChannelResponse, error)
}

func newChannel(c *slack.Channel) Channel {
	return Channel{
		ID:         c.ID,
		Name:       c.Name,
		IsPrivate:  c.IsPrivate,
		IsArchived: c.IsArchived,
	}
}

// API: https://api.slack.com/methods/conversations.list
func (a *api) GetChannels(req *GetChannelsRequest) (*GetChannelsResponse, error) {
	var (
		params = &slack.GetConversationsParameters{
			TeamID: req.TeamID,
		}
		res GetChannelsResponse
	)

	for {
		channels, cursor, err := a.client.GetConversations(params)
		if err != nil {
			return nil, err
		}
		for _, c := range channels {
			res.Channels = append(res.Channels, newChannel(&c))
		}
		if cursor == "" {
			return &res, nil
		}
		params.Cursor = cursor
	}
}

type GetChannelRequest struct {
	ID string
}

type GetChannelResponse struct {
	Channel Channel
}

func (a *api) getChannel(id string) (*Channel, error) {
	channel, err := a.client.GetConversationInfo(&slack.GetConversationInfoInput{
		ChannelID:         id,
		IncludeLocale:     false,
		IncludeNumMembers: false,
	})
	if err != nil {
		return nil, err
	}
	response := newChannel(channel)
	return &response, nil
}

// API: https://api.slack.com/methods/conversations.info
func (a *api) GetChannel(req *GetChannelRequest) (*GetChannelResponse, error) {
	var (
		res GetChannelResponse
	)
	channel, err := a.getChannel(req.ID)
	if err != nil {
		return nil, err
	}
	res.Channel = *channel
	return &res, nil
}

type CreateChannelRequest struct {
	Name      string
	IsPrivate bool
	TeamID    string
}

type CreateChannelResponse struct {
	Channel Channel
}

// API: https://api.slack.com/methods/conversations.create
func (a *api) CreateChannel(req *CreateChannelRequest) (*CreateChannelResponse, error) {
	var (
		params = slack.CreateConversationParams{
			ChannelName: req.Name,
			IsPrivate:   req.IsPrivate,
			TeamID:      req.TeamID,
		}
		res CreateChannelResponse
	)

	channel, err := a.client.CreateConversation(params)
	if err != nil {
		return nil, err
	}
	res.Channel = newChannel(channel)
	return &res, nil
}

type UpdateChannelRequest struct {
	ID         string
	Name       string
	IsArchived bool
}

type UpdateChannelResponse struct {
	Channel Channel
}

func (a *api) updateChannelArchived(id string, archived bool) error {
	fact, err := a.client.GetConversationInfo(&slack.GetConversationInfoInput{
		ChannelID:         id,
		IncludeLocale:     false,
		IncludeNumMembers: false,
	})
	if err != nil {
		return err
	}

	switch {
	case !fact.IsArchived && archived:
		if err := a.client.ArchiveConversation(id); err != nil {
			return err
		}
	case fact.IsArchived && !archived:
		if err := a.client.UnArchiveConversation(id); err != nil {
			return err
		}
	}
	return nil
}

// API: https://api.slack.com/methods/conversations.rename
// API: https://api.slack.com/methods/conversations.archive
// API: https://api.slack.com/methods/conversations.unarchive
func (a *api) UpdateChannel(req *UpdateChannelRequest) (*UpdateChannelResponse, error) {
	var (
		res UpdateChannelResponse
	)

	if err := a.updateChannelArchived(req.ID, req.IsArchived); err != nil {
		return nil, err
	}

	if req.Name != "" {
		if _, err := a.client.RenameConversation(req.ID, req.Name); err != nil {
			return nil, err
		}
	}

	channel, err := a.getChannel(req.ID)
	if err != nil {
		return nil, err
	}
	res.Channel = *channel
	return &res, nil
}
