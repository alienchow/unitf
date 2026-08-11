package fobs

import (
	"context"

	"github.com/alienchow/unitf/internal/client"
	"github.com/alienchow/unitf/internal/datasources"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type FobsDataSourceModel struct {
	Items  []FobModel                `tfsdk:"items"`
	Filter []datasources.FilterModel `tfsdk:"filter"`
}

type FobModel struct {
	ID types.String `tfsdk:"id"`
}

func NewFobsDataSource() datasource.DataSource {
	return &datasources.GenericDataSource[FobsDataSourceModel]{
		TypeName: "protect_fobs",
		TFSchema: schema.Schema{
			Description: "Data source for listing Fobs.",
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
		ReadFunc: func(ctx context.Context, c *client.Client, model *FobsDataSourceModel) error {
			items, err := c.Protect.ListFobs(ctx)
			if err != nil {
				return err
			}

			for _, i := range items {
				model.Items = append(model.Items, FobModel{
					ID: types.StringValue(i.ID),
				})
			}
			return nil
		},
	}
}
