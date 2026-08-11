package radius_profiles

import (
	"context"

	"github.com/alienchow/unitf/internal/client"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &RadiusProfilesDataSource{}
	_ datasource.DataSourceWithConfigure = &RadiusProfilesDataSource{}
)

type RadiusProfilesDataSource struct {
	client client.NetworkClient
}

type RadiusProfilesDataSourceModel struct {
	Items  []RadiusProfileModel `tfsdk:"items"`
	SiteID types.String         `tfsdk:"site_id"`
}

type RadiusProfileModel struct {
	ID types.String `tfsdk:"id"`
}

func NewRadiusProfilesDataSource() datasource.DataSource {
	return &RadiusProfilesDataSource{}
}

func (d *RadiusProfilesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_radius_profiles"
}

func (d *RadiusProfilesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Data source for listing RadiusProfiles.",
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

func (d *RadiusProfilesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *RadiusProfilesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state RadiusProfilesDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	items, err := d.client.ListRadiusProfiles(ctx, state.SiteID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error Reading RadiusProfiles", err.Error())
		return
	}

	for _, i := range items {
		state.Items = append(state.Items, RadiusProfileModel{
			ID: types.StringValue(i.ID),
		})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}
