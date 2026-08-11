package nvr

import (
	"context"

	"github.com/alienchow/unitf/internal/client"
	"github.com/alienchow/unitf/internal/datasources"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NvrDataSourceModel struct {
	Items  []NvrModel                `tfsdk:"items"`
	Filter []datasources.FilterModel `tfsdk:"filter"`
}

type NvrModel struct {
	ID types.String `tfsdk:"id"`
}

func NewNvrDataSource() datasource.DataSource {
	return &datasources.GenericDataSource[NvrDataSourceModel]{
		TypeName: "protect_nvr",
		TFSchema: schema.Schema{
			Description: "Data source for listing Nvr.",
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
		ReadFunc: func(ctx context.Context, c *client.Client, model *NvrDataSourceModel) error {
			items, err := c.Protect.ListNvr(ctx)
			if err != nil {
				return err
			}

			for _, i := range items {
				model.Items = append(model.Items, NvrModel{
					ID: types.StringValue(i.ID),
				})
			}
			return nil
		},
	}
}
