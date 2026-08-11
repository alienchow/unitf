package lights

import (
	"context"

	"github.com/alienchow/unitf/internal/client"
	"github.com/alienchow/unitf/internal/datasources"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type LightsDataSourceModel struct {
	Items  []LightModel              `tfsdk:"items"`
	Filter []datasources.FilterModel `tfsdk:"filter"`
}

type LightModel struct {
	ID types.String `tfsdk:"id"`
}

func NewLightsDataSource() datasource.DataSource {
	return &datasources.GenericDataSource[LightsDataSourceModel]{
		TypeName: "protect_lights",
		TFSchema: schema.Schema{
			Description: "Data source for listing Lights.",
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
		ReadFunc: func(ctx context.Context, c *client.Client, model *LightsDataSourceModel) error {
			items, err := c.Protect.ListLights(ctx)
			if err != nil {
				return err
			}

			for _, i := range items {
				model.Items = append(model.Items, LightModel{
					ID: types.StringValue(i.ID),
				})
			}
			return nil
		},
	}
}
