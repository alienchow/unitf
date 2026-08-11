package speakers

import (
	"context"

	"github.com/alienchow/unitf/internal/client"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &SpeakersDataSource{}
	_ datasource.DataSourceWithConfigure = &SpeakersDataSource{}
)

type SpeakersDataSource struct {
	client client.ProtectClient
}

type SpeakersDataSourceModel struct {
	Items []SpeakerModel `tfsdk:"items"`
}

type SpeakerModel struct {
	ID types.String `tfsdk:"id"`
}

func NewSpeakersDataSource() datasource.DataSource {
	return &SpeakersDataSource{}
}

func (d *SpeakersDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_speakers"
}

func (d *SpeakersDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Data source for listing Speakers.",
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

func (d *SpeakersDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *SpeakersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state SpeakersDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	items, err := d.client.ListSpeakers(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Error Reading Speakers", err.Error())
		return
	}

	for _, i := range items {
		state.Items = append(state.Items, SpeakerModel{
			ID: types.StringValue(i.ID),
		})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}
