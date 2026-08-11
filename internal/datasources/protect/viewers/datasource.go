package viewers

import (
	"context"

	"github.com/alienchow/unitf/internal/client"
	"github.com/alienchow/unitf/internal/datasources"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ViewersDataSourceModel struct {
	Filter []datasources.FilterModel `tfsdk:"filter"`
	Items  []ViewerModel             `tfsdk:"items"`
}

type ViewerModel struct {
	ID types.String `tfsdk:"id"`
}

func NewViewersDataSource() datasource.DataSource {
	return &datasources.GenericDataSource[ViewersDataSourceModel]{
		TypeName: "protect_viewers",
		TFSchema: schema.Schema{
			Description: "Data source for listing Viewers.",
			Attributes: map[string]schema.Attribute{
				"filter": schema.ListNestedAttribute{
					Optional:    true,
					Description: "Filters to apply to the items.",
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Required:    true,
								Description: "Name of the field to filter on.",
							},
							"values": schema.ListAttribute{
								ElementType: types.StringType,
								Required:    true,
								Description: "List of values to match against.",
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
								Description: "Unique ID.",
							},
						},
					},
				},
			},
		},
		ReadFunc: func(ctx context.Context, c *client.Client, model *ViewersDataSourceModel) error {
			items, err := c.Protect.ListViewers(ctx)
			if err != nil {
				return err
			}
			for _, i := range items {
				model.Items = append(model.Items, ViewerModel{
					ID: types.StringValue(i.ID),
				})
			}
			return nil
		},
	}
}
