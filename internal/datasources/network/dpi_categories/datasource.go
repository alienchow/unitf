package dpi_categories

import (
	"context"

	"github.com/alienchow/unitf/internal/client"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &DpiCategoriesDataSource{}
	_ datasource.DataSourceWithConfigure = &DpiCategoriesDataSource{}
)

type DpiCategoriesDataSource struct {
	client client.NetworkClient
}

type DpiCategoriesDataSourceModel struct {
	Items  []DpiCategoryModel `tfsdk:"items"`
	SiteID types.String       `tfsdk:"site_id"`
}

type DpiCategoryModel struct {
	ID types.String `tfsdk:"id"`
}

func NewDpiCategoriesDataSource() datasource.DataSource {
	return &DpiCategoriesDataSource{}
}

func (d *DpiCategoriesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dpi_categories"
}

func (d *DpiCategoriesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Data source for listing DpiCategories.",
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
			"site_id": schema.StringAttribute{
				Required:    true,
				Description: "Site ID.",
			},
		},
	}
}

func (d *DpiCategoriesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(client.NetworkClient)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Data Source Configure Type", "Expected client.NetworkClient")
		return
	}
	d.client = c
}

func (d *DpiCategoriesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state DpiCategoriesDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	items, err := d.client.ListDpiCategories(ctx, state.SiteID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error Reading DpiCategories", err.Error())
		return
	}

	for _, i := range items {
		state.Items = append(state.Items, DpiCategoryModel{
			ID: types.StringValue(i.ID),
		})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}
