package mc_lag_domains

import (
	"context"

	"github.com/alienchow/unitf/internal/client"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &McLagDomainsDataSource{}
	_ datasource.DataSourceWithConfigure = &McLagDomainsDataSource{}
)

type McLagDomainsDataSource struct {
	client client.NetworkClient
}

type McLagDomainsDataSourceModel struct {
	Items  []McLagDomainModel `tfsdk:"items"`
	SiteID types.String       `tfsdk:"site_id"`
}

type McLagDomainModel struct {
	ID types.String `tfsdk:"id"`
}

func NewMcLagDomainsDataSource() datasource.DataSource {
	return &McLagDomainsDataSource{}
}

func (d *McLagDomainsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mc_lag_domains"
}

func (d *McLagDomainsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Data source for listing McLagDomains.",
		Attributes: map[string]schema.Attribute{
			"items": schema.ListNestedAttribute{
				Computed:    true,
				Description: "List of items.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed:    true,
							Description: "Unique ID.",
						},
					},
				},
			},
			"site_id": schema.StringAttribute{
				Required:    true,
				Description: "Site ID.",
			},
		},
	}
}

func (d *McLagDomainsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *McLagDomainsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state McLagDomainsDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	items, err := d.client.ListMcLagDomains(ctx, state.SiteID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error Reading McLagDomains", err.Error())
		return
	}

	for _, i := range items {
		state.Items = append(state.Items, McLagDomainModel{
			ID: types.StringValue(i.ID),
		})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}
