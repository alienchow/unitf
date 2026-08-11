package firewall_policies

import (
	"context"

	"github.com/alienchow/unitf/internal/client"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &FirewallPoliciesDataSource{}
	_ datasource.DataSourceWithConfigure = &FirewallPoliciesDataSource{}
)

type FirewallPoliciesDataSource struct {
	client client.NetworkClient
}

type FirewallPoliciesDataSourceModel struct {
	SiteID types.String `tfsdk:"site_id"`
	Items  []FirewallpolicyModel `tfsdk:"items"`
}

type FirewallpolicyModel struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

func NewFirewallPoliciesDataSource() datasource.DataSource {
	return &FirewallPoliciesDataSource{}
}

func (d *FirewallPoliciesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_firewall_policies"
}

func (d *FirewallPoliciesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Data source for listing firewall_policies.",
		Attributes: map[string]schema.Attribute{
			"site_id": schema.StringAttribute{
				Required:    true,
				Description: "Site ID.",
			},
			"items": schema.ListNestedAttribute{
				Computed:    true,
				Description: "List of items.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed:    true,
							Description: "Unique UUID.",
						},
						"name": schema.StringAttribute{
							Computed:    true,
							Description: "Name.",
						},
					},
				},
			},
		},
	}
}

func (d *FirewallPoliciesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(client.NetworkClient)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Data Source Configure Type", "Expected client.NetworkClient")
		return
	}
	d.client = c
}

func (d *FirewallPoliciesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state FirewallPoliciesDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	items, err := d.client.ListFirewallPolicies(ctx, state.SiteID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error Reading firewall_policies", err.Error())
		return
	}

	state.Items = make([]FirewallpolicyModel, 0, len(items))
	for _, n := range items {
		state.Items = append(state.Items, FirewallpolicyModel{
			ID:   types.StringValue(n.ID),
			Name: types.StringValue(n.Name),
		})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}
