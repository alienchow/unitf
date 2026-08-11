package lags

import (
	"context"

	"github.com/alienchow/unitf/internal/client"
	"github.com/alienchow/unitf/internal/datasources"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type LagsDataSourceModel struct {
	SiteID types.String              `tfsdk:"site_id"`
	Filter []datasources.FilterModel `tfsdk:"filter"`
	Items  []LagModel                `tfsdk:"items"`
}

type LagModel struct {
	ID types.String `tfsdk:"id"`
}

func NewLagsDataSource() datasource.DataSource {
	return &datasources.GenericDataSource[LagsDataSourceModel]{
		TypeName: "network_lags",
		TFSchema: schema.Schema{
			Description: "Data source for listing Lags.",
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
								Description: "Unique ID.",
							},
						},
					},
				},
			},
		},
		ReadFunc: func(ctx context.Context, c *client.Client, model *LagsDataSourceModel) error {
			items, err := c.Network.ListLags(ctx, model.SiteID.ValueString())
			if err != nil {
				return err
			}

			model.Items = make([]LagModel, 0, len(items))
			for _, i := range items {
				model.Items = append(model.Items, LagModel{
					ID: types.StringValue(i.ID),
				})
			}
			return nil
		},
	}
}
