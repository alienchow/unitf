package traffic_matching_lists

import (
	"context"

	"github.com/alienchow/unitf/internal/client"
	"github.com/alienchow/unitf/internal/datasources"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TrafficMatchingListsDataSourceModel struct {
	SiteID types.String               `tfsdk:"site_id"`
	Filter []datasources.FilterModel  `tfsdk:"filter"`
	Items  []TrafficmatchinglistModel `tfsdk:"items"`
}

type TrafficmatchinglistModel struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

func NewTrafficMatchingListsDataSource() datasource.DataSource {
	return &datasources.GenericDataSource[TrafficMatchingListsDataSourceModel]{
		TypeName: "network_traffic_matching_lists",
		TFSchema: schema.Schema{
			Description: "Data source for listing traffic_matching_lists.",
			Attributes: map[string]schema.Attribute{
				"site_id": schema.StringAttribute{
					Required:    true,
					Description: "Site ID.",
				},
				"filter": schema.ListNestedAttribute{
					Optional:    true,
					Description: "Filters.",
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Required: true,
							},
							"values": schema.ListAttribute{
								ElementType: types.StringType,
								Required:    true,
							},
						},
					},
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
		},
		ReadFunc: func(ctx context.Context, c *client.Client, model *TrafficMatchingListsDataSourceModel) error {
			items, err := c.Network.ListTrafficMatchingLists(ctx, model.SiteID.ValueString())
			if err != nil {
				return err
			}

			model.Items = make([]TrafficmatchinglistModel, 0, len(items))
			for _, n := range items {
				model.Items = append(model.Items, TrafficmatchinglistModel{
					ID:   types.StringValue(n.ID),
					Name: types.StringValue(n.Name),
				})
			}
			return nil
		},
	}
}
