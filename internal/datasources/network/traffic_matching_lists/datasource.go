package traffic_matching_lists

import (
	"context"

	"github.com/alienchow/unitf/internal/client"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &TrafficMatchingListsDataSource{}
	_ datasource.DataSourceWithConfigure = &TrafficMatchingListsDataSource{}
)

type TrafficMatchingListsDataSource struct {
	client client.NetworkClient
}

type TrafficMatchingListsDataSourceModel struct {
	SiteID types.String               `tfsdk:"site_id"`
	Items  []TrafficmatchinglistModel `tfsdk:"items"`
}

type TrafficmatchinglistModel struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

func NewTrafficMatchingListsDataSource() datasource.DataSource {
	return &TrafficMatchingListsDataSource{}
}

func (d *TrafficMatchingListsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_traffic_matching_lists"
}

func (d *TrafficMatchingListsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Data source for listing traffic_matching_lists.",
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

func (d *TrafficMatchingListsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *TrafficMatchingListsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state TrafficMatchingListsDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	items, err := d.client.ListTrafficMatchingLists(ctx, state.SiteID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error Reading traffic_matching_lists", err.Error())
		return
	}

	state.Items = make([]TrafficmatchinglistModel, 0, len(items))
	for _, n := range items {
		state.Items = append(state.Items, TrafficmatchinglistModel{
			ID:   types.StringValue(n.ID),
			Name: types.StringValue(n.Name),
		})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}
