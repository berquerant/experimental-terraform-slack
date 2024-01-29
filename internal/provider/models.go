package provider

import (
	"github.com/berquerant/terraform-slack/swagger/models"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type channelModel struct {
	ID         types.String `tfsdk:"id"`
	Name       types.String `tfsdk:"name"`
	TeamID     types.String `tfsdk:"team_id"`
	IsPrivate  types.Bool   `tfsdk:"is_private"`
	IsArchived types.Bool   `tfsdk:"is_archived"`
}

func newChannelModel(c *models.Channel) channelModel {
	return channelModel{
		ID:         types.StringValue(c.ID),
		Name:       types.StringValue(c.Name),
		IsPrivate:  types.BoolValue(c.IsPrivate),
		IsArchived: types.BoolValue(c.IsArchived),
	}
}
