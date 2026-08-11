package cameras

import (
	"context"

	"github.com/alienchow/unitf/internal/client"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &CamerasDataSource{}
	_ datasource.DataSourceWithConfigure = &CamerasDataSource{}
)

type CamerasDataSource struct {
	client client.ProtectClient
}

type CamerasDataSourceModel struct {
	Items []CameraModel `tfsdk:"items"`
}

type CameraModel struct {
	ID types.String `tfsdk:"id"`
}

func NewCamerasDataSource() datasource.DataSource {
	return &CamerasDataSource{}
}

func (d *CamerasDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cameras"
}

func (d *CamerasDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Data source for listing Cameras.",
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

func (d *CamerasDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *CamerasDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state CamerasDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	items, err := d.client.ListCameras(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Error Reading Cameras", err.Error())
		return
	}

	for _, i := range items {
		state.Items = append(state.Items, CameraModel{
			ID: types.StringValue(i.ID),
		})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}
