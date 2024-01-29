package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/berquerant/terraform-slack/swagger/client/channel"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                = &channelResource{}
	_ resource.ResourceWithConfigure   = &channelResource{}
	_ resource.ResourceWithImportState = &channelResource{}
)

func NewChannelResource() resource.Resource {
	return &channelResource{}
}

type channelResource struct {
	*resourceData
}

type channelResourceModel struct {
	ID         types.String `tfsdk:"id"`
	Name       types.String `tfsdk:"name"`
	TeamID     types.String `tfsdk:"team_id"`
	IsPrivate  types.Bool   `tfsdk:"is_private"`
	IsArchived types.Bool   `tfsdk:"is_archived"`
}

func (r *channelResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_channel"
}

func (r *channelResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages a channel.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Channel ID.",
				Computed:    true,
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "Channel name.",
				Required:    true,
			},
			"team_id": schema.StringAttribute{
				Description: "Team ID.",
				Optional:    true,
			},
			"is_private": schema.BoolAttribute{
				Description: "If true, private channel.",
				Optional:    true,
			},
			"is_archived": schema.BoolAttribute{
				Description: "If true, archived channel.",
				Optional:    true,
			},
		},
	}
}

func (r *channelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state channelResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	id := state.ID.ValueString()
	params := channel.NewFetchChannelParams()
	params.ID = id
	response, err := r.client.Channel.FetchChannel(params, r.authInfo)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Slack Channel",
			"Could not read Slack channel id "+state.ID.ValueString()+": "+err.Error(),
		)
		return
	}

	c := response.Payload.Channel
	state.Name = types.StringValue(c.Name)
	state.IsPrivate = types.BoolValue(c.IsPrivate)
	state.IsArchived = types.BoolValue(c.IsArchived)

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *channelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan channelResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var (
		body       channel.CreateChannelBody
		id         = plan.ID.ValueString()
		name       = plan.Name.ValueString()
		teamID     = plan.TeamID.ValueString()
		isPrivate  = plan.IsPrivate.ValueBool()
		isArchived = plan.IsArchived.ValueBool()
	)
	body.Name = &name
	body.TeamID = &teamID
	body.IsPrivate = &isPrivate
	params := channel.NewCreateChannelParams().
		WithChannel(body)
	created, err := r.client.Channel.CreateChannel(params, r.authInfo)
	if err != nil {
		if !strings.Contains(err.Error(), "name_taken") {
			resp.Diagnostics.AddError(
				"Error creating channel",
				"Clould not create channel, unexpected error: "+err.Error(),
			)
			return
		}
	} else {
		id = created.Payload.Channel.ID
	}

	if isArchived {
		body := channel.UpdateChannelBody{
			IsArchived: true,
		}
		params := channel.NewUpdateChannelParams().
			WithID(id).
			WithChannel(body)
		if _, err := r.client.Channel.UpdateChannel(params, r.authInfo); err != nil {
			resp.Diagnostics.AddError(
				"Error creating channel",
				"Clould not create channel, unexpected error: "+err.Error(),
			)
			return
		}
		plan.IsArchived = types.BoolValue(true)
	}

	plan.ID = types.StringValue(id)
	if teamID != "" {
		plan.TeamID = types.StringValue(teamID)
	}
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *channelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan channelResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var (
		id   = plan.ID.ValueString()
		body = channel.UpdateChannelBody{
			IsArchived: plan.IsArchived.ValueBool(),
			Name:       plan.Name.ValueString(),
		}
		params = channel.NewUpdateChannelParams().
			WithID(id).
			WithChannel(body)
	)

	if _, err := r.client.Channel.UpdateChannel(params, r.authInfo); err != nil {
		resp.Diagnostics.AddError(
			"Error updating channel",
			"Could not update channel, unexpected channel: "+err.Error(),
		)
		return
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *channelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.Diagnostics.AddWarning(
		"Warning deleting channel",
		"Not implemented",
	)
}

func (r *channelResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	data, ok := req.ProviderData.(*resourceData)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected resourceData, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.resourceData = data
}

func (r *channelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
