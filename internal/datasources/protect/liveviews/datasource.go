package liveviews

import (
	"context"

	"github.com/alienchow/unitf/internal/client"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &LiveviewsDataSource{}
	_ datasource.DataSourceWithConfigure = &LiveviewsDataSource{}
)

type LiveviewsDataSource struct {
	client client.ProtectClient
}

type LiveviewsDataSourceModel struct {
	Items []LiveviewModel `tfsdk:"items"`
}

type LiveviewModel struct {
	ID types.String `tfsdk:"id"`
}

func NewLiveviewsDataSource() datasource.DataSource {
	return &LiveviewsDataSource{}
}

func (d *LiveviewsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_liveviews"
}

func (d *LiveviewsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Data source for listing Liveviews.",
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
		},
	}
}

func (d *LiveviewsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(client.ProtectClient)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Data Source Configure Type", "Expected client.ProtectClient")
		return
	}
	d.client = c
}

func (d *LiveviewsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state LiveviewsDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	items, err := d.client.ListLiveviews(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Error Reading Liveviews", err.Error())
		return
	}

	for _, i := range items {
		state.Items = append(state.Items, LiveviewModel{
			ID: types.StringValue(i.ID),
		})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}
