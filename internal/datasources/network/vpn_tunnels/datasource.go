package vpn_tunnels

import (
	"context"

	"github.com/alienchow/unitf/internal/client"
	"github.com/alienchow/unitf/internal/datasources"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type VpnTunnelsDataSourceModel struct {
	SiteID types.String              `tfsdk:"site_id"`
	Filter []datasources.FilterModel `tfsdk:"filter"`
	Items  []VpnTunnelModel          `tfsdk:"items"`
}

type VpnTunnelModel struct {
	ID types.String `tfsdk:"id"`
}

func NewVpnTunnelsDataSource() datasource.DataSource {
	return &datasources.GenericDataSource[VpnTunnelsDataSourceModel]{
		TypeName: "network_vpn_tunnels",
		TFSchema: schema.Schema{
			Description: "Data source for listing VpnTunnels.",
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
		ReadFunc: func(ctx context.Context, c *client.Client, model *VpnTunnelsDataSourceModel) error {
			items, err := c.Network.ListVpnTunnels(ctx, model.SiteID.ValueString())
			if err != nil {
				return err
			}

			model.Items = make([]VpnTunnelModel, 0, len(items))
			for _, i := range items {
				model.Items = append(model.Items, VpnTunnelModel{
					ID: types.StringValue(i.ID),
				})
			}
			return nil
		},
	}
}
