package acl_rules

import (
	"context"

	"github.com/alienchow/unitf/internal/client"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &AclRulesDataSource{}
	_ datasource.DataSourceWithConfigure = &AclRulesDataSource{}
)

type AclRulesDataSource struct {
	client client.NetworkClient
}

type AclRulesDataSourceModel struct {
	SiteID types.String   `tfsdk:"site_id"`
	Items  []AclruleModel `tfsdk:"items"`
}

type AclruleModel struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

func NewAclRulesDataSource() datasource.DataSource {
	return &AclRulesDataSource{}
}

func (d *AclRulesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_acl_rules"
}

func (d *AclRulesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Data source for listing acl_rules.",
		Attributes: map[string]schema.Attribute{
			"site_id": schema.StringAttribute{
				Required:    true,
				Description: "Site ID.",
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
	}
}

func (d *AclRulesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *AclRulesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state AclRulesDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	items, err := d.client.ListAclRules(ctx, state.SiteID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error Reading acl_rules", err.Error())
		return
	}

	state.Items = make([]AclruleModel, 0, len(items))
	for _, n := range items {
		state.Items = append(state.Items, AclruleModel{
			ID:   types.StringValue(n.ID),
			Name: types.StringValue(n.Name),
		})
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}
