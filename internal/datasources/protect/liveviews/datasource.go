package liveviews

import (
	"context"

	"github.com/alienchow/unitf/internal/client"
	"github.com/alienchow/unitf/internal/datasources"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type LiveviewsDataSourceModel struct {
	Items  []LiveviewModel           `tfsdk:"items"`
	Filter []datasources.FilterModel `tfsdk:"filter"`
}

type LiveviewModel struct {
	ID types.String `tfsdk:"id"`
}

func NewLiveviewsDataSource() datasource.DataSource {
	return &datasources.GenericDataSource[LiveviewsDataSourceModel]{
		TypeName: "protect_liveviews",
		TFSchema: schema.Schema{
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
				"filter": schema.ListNestedAttribute{
					Optional:    true,
					Description: "Filters to apply.",
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
			},
		},
		ReadFunc: func(ctx context.Context, c *client.Client, model *LiveviewsDataSourceModel) error {
			items, err := c.Protect.ListLiveviews(ctx)
			if err != nil {
				return err
			}

			for _, i := range items {
				model.Items = append(model.Items, LiveviewModel{
					ID: types.StringValue(i.ID),
				})
			}
			return nil
		},
	}
}
