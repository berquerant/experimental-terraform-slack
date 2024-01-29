package channel

import (
	"net/http"

	"github.com/berquerant/terraform-slack/pkg/slackapi"
	"github.com/berquerant/terraform-slack/swagger/models"
	"github.com/go-openapi/runtime/middleware"
)

func newChannel(c *slackapi.Channel) *models.Channel {
	return &models.Channel{
		ID:         c.ID,
		Name:       c.Name,
		IsPrivate:  c.IsPrivate,
		IsArchived: c.IsArchived,
	}
}

func HandleUpdateChannel(params UpdateChannelParams, principal *models.Principal) middleware.Responder {
	api := slackapi.New(string(*principal))

	var req slackapi.UpdateChannelRequest
	req.ID = params.ID
	req.Name = params.Channel.Name
	req.IsArchived = params.Channel.IsArchived

	response, err := api.UpdateChannel(&req)
	if err != nil {
		e := models.NewErrorModel(
			err.Error(),
			"failed to update channel",
		)
		return NewUpdateChannelDefault(http.StatusInternalServerError).WithPayload(&e)
	}

	payload := &models.ChannelModel{
		BaseModel: models.NewBaseModel(true),
		Channel:   newChannel(&response.Channel),
	}
	return NewUpdateChannelOK().WithPayload(payload)
}

func HandleCreateChannel(params CreateChannelParams, principal *models.Principal) middleware.Responder {
	api := slackapi.New(string(*principal))

	var req slackapi.CreateChannelRequest
	if x := params.Channel.Name; x != nil {
		req.Name = *x
	}
	if x := params.Channel.IsPrivate; x != nil {
		req.IsPrivate = *x
	}
	if x := params.Channel.TeamID; x != nil {
		req.TeamID = *x
	}

	response, err := api.CreateChannel(&req)
	if err != nil {
		e := models.NewErrorModel(
			err.Error(),
			"failed to create channel",
		)
		return NewCreateChannelDefault(http.StatusInternalServerError).WithPayload(&e)
	}

	payload := &models.ChannelModel{
		BaseModel: models.NewBaseModel(true),
		Channel:   newChannel(&response.Channel),
	}
	return NewCreateChannelCreated().WithPayload(payload)
}

func HandleFetchChannel(params FetchChannelParams, principal *models.Principal) middleware.Responder {
	api := slackapi.New(string(*principal))

	var req slackapi.GetChannelRequest
	req.ID = params.ID

	response, err := api.GetChannel(&req)
	if err != nil {
		e := models.NewErrorModel(
			err.Error(),
			"failed to fetch channel",
		)
		return NewFetchChannelDefault(http.StatusInternalServerError).WithPayload(&e)
	}

	payload := &models.ChannelModel{
		BaseModel: models.NewBaseModel(true),
		Channel:   newChannel(&response.Channel),
	}
	return NewFetchChannelOK().WithPayload(payload)
}

func HandleFetchChannels(params FetchChannelsParams, principal *models.Principal) middleware.Responder {
	api := slackapi.New(string(*principal))

	var req slackapi.GetChannelsRequest
	if x := params.TeamID; x != nil {
		req.TeamID = *x
	}

	response, err := api.GetChannels(&req)
	if err != nil {
		e := models.NewErrorModel(
			err.Error(),
			"failed to fetch channels",
		)
		return NewFetchChannelsDefault(http.StatusInternalServerError).WithPayload(&e)
	}

	channels := make([]*models.Channel, len(response.Channels))
	for i, c := range response.Channels {
		channels[i] = newChannel(&c)
	}
	payload := &models.ChannelsModel{
		BaseModel: models.NewBaseModel(true),
		Channels:  channels,
	}
	return NewFetchChannelsOK().WithPayload(payload)
}
