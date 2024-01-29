package provider

import (
	"context"
	"os"

	"github.com/berquerant/terraform-slack/swagger/client"
	"github.com/berquerant/terraform-slack/swagger/client/health"
	"github.com/go-openapi/runtime"
	transport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ provider.Provider = &SlackProvider{}

type SlackProvider struct {
	version string
}

type SlackProviderModel struct {
	Endpoint types.String `tfsdk:"endpoint"`
	Token    types.String `tfsdk:"token"`
}

func (p *SlackProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "slack"
	resp.Version = p.version
}

func (p *SlackProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Slack",
		Attributes: map[string]schema.Attribute{
			"endpoint": schema.StringAttribute{
				MarkdownDescription: "API endpoint",
				Optional:            true,
			},
			"token": schema.StringAttribute{
				MarkdownDescription: "API token",
				Optional:            true,
			},
		},
	}
}

func (p *SlackProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config SlackProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if config.Endpoint.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("endpoint"),
			"Unknown Slack API Endpoint",
			"The provider cannot create the Slack API client as there is an unknown configuration value for the Slack API host. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the SLACK_ENDPOINT environment variable.",
		)
	}
	if config.Token.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("token"),
			"Unknown Slack API Token",
			"The provider cannot create the Slack API client as there is an unknown configuration value for the Slack API token. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the SLACK_TOKEN environment variable.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	endpoint := os.Getenv("SLACK_ENDPOINT")
	token := os.Getenv("SLACK_TOKEN")

	if !config.Endpoint.IsNull() {
		endpoint = config.Endpoint.ValueString()
	}
	if !config.Token.IsNull() {
		token = config.Token.ValueString()
	}

	if endpoint == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("endpoint"),
			"Missing Slack API Endpoint",
			"The provider cannot create the Slack API client as there is a missing or empty value for the Slack API host. "+
				"Set the endpoint value in the configuration or use the SLACK_ENDPOINT environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}
	if token == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("token"),
			"Missing Slack API Token",
			"The provider cannot create the Slack API client as there is a missing or empty value for the Slack API token. "+
				"Set the token value in the configuration or use the SLACK_TOKEN environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	cfg := client.DefaultTransportConfig().WithHost(endpoint)
	client := client.NewHTTPClientWithConfig(strfmt.Default, cfg)
	authInfo := transport.APIKeyAuth("X-API-Key", "header", token)

	if _, err := client.Health.CheckAlive(health.NewCheckAliveParams(), authInfo); err != nil {
		resp.Diagnostics.AddError(
			"API is unhealthy",
			"The provider cannot find healthy API: "+err.Error(),
		)
	}
	if resp.Diagnostics.HasError() {
		return
	}

	resp.DataSourceData = &dataSourceData{
		client:   client,
		authInfo: authInfo,
	}
	resp.ResourceData = &resourceData{
		client:   client,
		authInfo: authInfo,
	}

	tflog.Info(ctx, "Configured Slack client", map[string]any{"endpoint": endpoint})
}

type dataSourceData struct {
	client   *client.SlackAPIForTerraformProvider
	authInfo runtime.ClientAuthInfoWriter
}

type resourceData struct {
	client   *client.SlackAPIForTerraformProvider
	authInfo runtime.ClientAuthInfoWriter
}

func (p *SlackProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewChannelResource,
	}
}

func (p *SlackProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewChannelsDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &SlackProvider{
			version: version,
		}
	}
}
