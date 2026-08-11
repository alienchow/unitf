package sensors

import (
	"context"

	"github.com/alienchow/unitf/internal/client"
	"github.com/alienchow/unitf/internal/datasources"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SensorsDataSourceModel struct {
	Filter []datasources.FilterModel `tfsdk:"filter"`
	Items  []SensorModel             `tfsdk:"items"`
}

type SensorModel struct {
	ID types.String `tfsdk:"id"`
}

func NewSensorsDataSource() datasource.DataSource {
	return &datasources.GenericDataSource[SensorsDataSourceModel]{
		TypeName: "protect_sensors",
		TFSchema: schema.Schema{
			Description: "Data source for listing Sensors.",
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
		ReadFunc: func(ctx context.Context, c *client.Client, model *SensorsDataSourceModel) error {
			items, err := c.Protect.ListSensors(ctx)
			if err != nil {
				return err
			}
			for _, i := range items {
				model.Items = append(model.Items, SensorModel{
					ID: types.StringValue(i.ID),
				})
			}
			return nil
		},
	}
}
