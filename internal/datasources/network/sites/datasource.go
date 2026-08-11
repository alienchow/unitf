package sites

import (
	"context"

	"github.com/alienchow/unitf/internal/client"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &SitesDataSource{}
	_ datasource.DataSourceWithConfigure = &SitesDataSource{}
)

type SitesDataSource struct {
	client client.NetworkClient
}

type SitesDataSourceModel struct {
	Sites []SiteModel `tfsdk:"sites"`
}

type SiteModel struct {
	ID                types.String `tfsdk:"id"`
	Name              types.String `tfsdk:"name"`
	InternalReference types.String `tfsdk:"internal_reference"`
}

func NewSitesDataSource() datasource.DataSource {
	return &SitesDataSource{}
}

func (d *SitesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sites"
}

func (d *SitesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Data source for listing local UniFi Sites.",
		Attributes: map[string]schema.Attribute{
			"sites": schema.ListNestedAttribute{
				Computed:    true,
				Description: "List of sites.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed:    true,
							Description: "Unique site UUID.",
						},
						"name": schema.StringAttribute{
							Computed:    true,
							Description: "Name of the site.",
						},
						"internal_reference": schema.StringAttribute{
							Computed:    true,
							Description: "Internal reference handle.",
						},
					},
				},
			},
		},
	}
}

func (d *SitesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *SitesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	sites, err := d.client.ListSites(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Error Reading UniFi Sites", err.Error())
		return
	}

	var state SitesDataSourceModel
	for _, s := range sites {
		state.Sites = append(state.Sites, SiteModel{
			ID:                types.StringValue(s.ID),
			Name:              types.StringValue(s.Name),
			InternalReference: types.StringValue(s.InternalReference),
		})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}
