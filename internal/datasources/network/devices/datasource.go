package devices

import (
	"context"

	"github.com/alienchow/unitf/internal/client"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &DevicesDataSource{}
	_ datasource.DataSourceWithConfigure = &DevicesDataSource{}
)

type DevicesDataSource struct {
	client client.NetworkClient
}

type DevicesDataSourceModel struct {
	Items  []DeviceModel `tfsdk:"items"`
	SiteID types.String  `tfsdk:"site_id"`
}

type DeviceModel struct {
	ID types.String `tfsdk:"id"`
}

func NewDevicesDataSource() datasource.DataSource {
	return &DevicesDataSource{}
}

func (d *DevicesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices"
}

func (d *DevicesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Data source for listing Devices.",
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

func (d *DevicesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *DevicesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state DevicesDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	items, err := d.client.ListDevices(ctx, state.SiteID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error Reading Devices", err.Error())
		return
	}

	for _, i := range items {
		state.Items = append(state.Items, DeviceModel{
			ID: types.StringValue(i.ID),
		})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}
