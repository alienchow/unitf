package cameras

import (
	"context"

	"github.com/alienchow/unitf/internal/client"
	"github.com/alienchow/unitf/internal/datasources"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CamerasDataSourceModel struct {
	Items  []CameraModel             `tfsdk:"items"`
	Filter []datasources.FilterModel `tfsdk:"filter"`
}

type CameraModel struct {
	ID types.String `tfsdk:"id"`
}

func NewCamerasDataSource() datasource.DataSource {
	return &datasources.GenericDataSource[CamerasDataSourceModel]{
		TypeName: "protect_cameras",
		TFSchema: schema.Schema{
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
		ReadFunc: func(ctx context.Context, c *client.Client, model *CamerasDataSourceModel) error {
			items, err := c.Protect.ListCameras(ctx)
			if err != nil {
				return err
			}

			for _, i := range items {
				model.Items = append(model.Items, CameraModel{
					ID: types.StringValue(i.ID),
				})
			}
			return nil
		},
	}
}
