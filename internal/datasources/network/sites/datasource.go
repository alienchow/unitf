package sites

import (
	"context"

	"github.com/alienchow/unitf/internal/client"
	"github.com/alienchow/unitf/internal/datasources"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SitesDataSourceModel struct {
	Filter []datasources.FilterModel `tfsdk:"filter"`
	Sites  []SiteModel               `tfsdk:"sites"`
}

type SiteModel struct {
	ID                types.String `tfsdk:"id"`
	Name              types.String `tfsdk:"name"`
	InternalReference types.String `tfsdk:"internal_reference"`
}

func NewSitesDataSource() datasource.DataSource {
	return &datasources.GenericDataSource[SitesDataSourceModel]{
		TypeName: "sites",
		TFSchema: schema.Schema{
			Description: "Data source for listing local UniFi Sites.",
			Attributes: map[string]schema.Attribute{
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
				"sites": schema.ListNestedAttribute{
					Computed:    true,
					Description: "List of sites.",
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"id": schema.StringAttribute{
								Computed:    true,
								Description: "Unique site UUID.",
							},
							"name": schema.StringAttribute{
								Computed:    true,
								Description: "Name of the site.",
							},
							"internal_reference": schema.StringAttribute{
								Computed:    true,
								Description: "Internal reference handle.",
							},
						},
					},
				},
			},
		},
		ReadFunc: func(ctx context.Context, c *client.Client, model *SitesDataSourceModel) error {
			sites, err := c.Network.ListSites(ctx)
			if err != nil {
				return err
			}

			model.Sites = make([]SiteModel, 0, len(sites))
			for _, s := range sites {
				model.Sites = append(model.Sites, SiteModel{
					ID:                types.StringValue(s.ID),
					Name:              types.StringValue(s.Name),
					InternalReference: types.StringValue(s.InternalReference),
				})
			}
			return nil
		},
	}
}
