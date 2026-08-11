package networks

import (
	"context"

	"github.com/alienchow/unitf/internal/client"
	"github.com/alienchow/unitf/internal/datasources"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type NetworksDataSourceModel struct {
	SiteID types.String              `tfsdk:"site_id"`
	Filter []datasources.FilterModel `tfsdk:"filter"`
	Items  []NetworkModel            `tfsdk:"items"`
}

type NetworkModel struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

func NewNetworksDataSource() datasource.DataSource {
	return &datasources.GenericDataSource[NetworksDataSourceModel]{
		TypeName: "networks",
		TFSchema: schema.Schema{
			Description: "Data source for listing networks.",
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
		ReadFunc: func(ctx context.Context, c *client.Client, model *NetworksDataSourceModel) error {
			items, err := c.Network.ListNetworks(ctx, model.SiteID.ValueString())
			if err != nil {
				return err
			}

			model.Items = make([]NetworkModel, 0, len(items))
			for _, n := range items {
				model.Items = append(model.Items, NetworkModel{
					ID:   types.StringValue(n.ID),
					Name: types.StringValue(n.Name),
				})
			}
			return nil
		},
	}
}
