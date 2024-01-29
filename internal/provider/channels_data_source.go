package provider

import (
	"context"
	"fmt"

	"github.com/berquerant/terraform-slack/swagger/client/channel"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &channelsDataSource{}
	_ datasource.DataSourceWithConfigure = &channelsDataSource{}
)

func NewChannelsDataSource() datasource.DataSource {
	return &channelsDataSource{}
}

type channelsDataSource struct {
	*dataSourceData
}

type channelsDataSourceModel struct {
	TeamID   types.String   `tfsdk:"team_id"`
	Channels []channelModel `tfsdk:"channels"`
}

func (d *channelsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_channels"
}

func (d *channelsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetch the list of channels.",
		Attributes: map[string]schema.Attribute{
			"team_id": schema.StringAttribute{
				Description: "Team ID.",
				Optional:    true,
			},
			"channels": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Channel ID.",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "Channel name.",
							Computed:    true,
						},
						"team_id": schema.StringAttribute{
							Description: "Team ID.",
							Computed:    true,
						},
						"is_private": schema.BoolAttribute{
							Description: "If true, private channel.",
							Computed:    true,
						},
						"is_archived": schema.BoolAttribute{
							Description: "If true, archived channel.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *channelsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state channelsDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	teamID := state.TeamID.ValueString()
	params := channel.NewFetchChannelsParams().
		WithTeamID(&teamID)
	response, err := d.client.Channel.FetchChannels(params, d.authInfo)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Slack Channels",
			"Could not read Slack channels: "+err.Error(),
		)
		return
	}

	channelModels := make([]channelModel, len(response.Payload.Channels))
	for i, c := range response.Payload.Channels {
		channelModels[i] = newChannelModel(c)
		if teamID != "" {
			channelModels[i].TeamID = types.StringValue(teamID)
		}
	}
	state.Channels = channelModels

	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (d *channelsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	data, ok := req.ProviderData.(*dataSourceData)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected dataSourceData, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	d.dataSourceData = data
}
