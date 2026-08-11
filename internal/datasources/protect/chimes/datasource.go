package chimes

import (
	"context"

	"github.com/alienchow/unitf/internal/client"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &ChimesDataSource{}
	_ datasource.DataSourceWithConfigure = &ChimesDataSource{}
)

type ChimesDataSource struct {
	client client.ProtectClient
}

type ChimesDataSourceModel struct {
	Items []ChimeModel `tfsdk:"items"`
}

type ChimeModel struct {
	ID types.String `tfsdk:"id"`
}

func NewChimesDataSource() datasource.DataSource {
	return &ChimesDataSource{}
}

func (d *ChimesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_chimes"
}

func (d *ChimesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Data source for listing Chimes.",
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

func (d *ChimesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *ChimesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state ChimesDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	items, err := d.client.ListChimes(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Error Reading Chimes", err.Error())
		return
	}

	for _, i := range items {
		state.Items = append(state.Items, ChimeModel{
			ID: types.StringValue(i.ID),
		})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}
