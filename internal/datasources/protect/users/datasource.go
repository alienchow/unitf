package users

import (
	"context"

	"github.com/alienchow/unitf/internal/client"
	"github.com/alienchow/unitf/internal/datasources"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type UsersDataSourceModel struct {
	Filter []datasources.FilterModel `tfsdk:"filter"`
	Items  []UserModel               `tfsdk:"items"`
}

type UserModel struct {
	ID types.String `tfsdk:"id"`
}

func NewUsersDataSource() datasource.DataSource {
	return &datasources.GenericDataSource[UsersDataSourceModel]{
		TypeName: "protect_users",
		TFSchema: schema.Schema{
			Description: "Data source for listing Users.",
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
		ReadFunc: func(ctx context.Context, c *client.Client, model *UsersDataSourceModel) error {
			items, err := c.Protect.ListUsers(ctx)
			if err != nil {
				return err
			}
			for _, i := range items {
				model.Items = append(model.Items, UserModel{
					ID: types.StringValue(i.ID),
				})
			}
			return nil
		},
	}
}
