package sirens

import (
	"context"

	"github.com/alienchow/unitf/internal/client"
	"github.com/alienchow/unitf/internal/datasources"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SirensDataSourceModel struct {
	Filter []datasources.FilterModel `tfsdk:"filter"`
	Items  []SirenModel              `tfsdk:"items"`
}

type SirenModel struct {
	ID types.String `tfsdk:"id"`
}

func NewSirensDataSource() datasource.DataSource {
	return &datasources.GenericDataSource[SirensDataSourceModel]{
		TypeName: "protect_sirens",
		TFSchema: schema.Schema{
			Description: "Data source for listing Sirens.",
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
		ReadFunc: func(ctx context.Context, c *client.Client, model *SirensDataSourceModel) error {
			items, err := c.Protect.ListSirens(ctx)
			if err != nil {
				return err
			}
			for _, i := range items {
				model.Items = append(model.Items, SirenModel{
					ID: types.StringValue(i.ID),
				})
			}
			return nil
		},
	}
}
