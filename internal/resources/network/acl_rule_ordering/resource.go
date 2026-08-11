package network

import (
	"context"

	"github.com/alienchow/unitf/internal/client"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource              = &AclRuleOrderingResource{}
	_ resource.ResourceWithConfigure = &AclRuleOrderingResource{}
)

type AclRuleOrderingResource struct {
	client client.NetworkClient
}

type AclRuleOrderingResourceModel struct {
	ID      types.String   `tfsdk:"id"` // Format: site_id
	SiteID  types.String   `tfsdk:"site_id"`
	RuleIDs []types.String `tfsdk:"rule_ids"`
}

func NewAclRuleOrderingResource() resource.Resource {
	return &AclRuleOrderingResource{}
}

func (r *AclRuleOrderingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_acl_rule_ordering"
}

func (r *AclRuleOrderingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages the execution order of UniFi ACL Rules.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "Site ID used as UUID.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"site_id": schema.StringAttribute{
				Required:    true,
				Description: "ID of the site.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"rule_ids": schema.ListAttribute{
				Required:    true,
				ElementType: types.StringType,
				Description: "Ordered list of ACL rule IDs.",
			},
		},
	}
}

func (r *AclRuleOrderingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(client.NetworkClient)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Resource Configure Type", "Expected client.NetworkClient")
		return
	}
	r.client = c
}

func (r *AclRuleOrderingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan AclRuleOrderingResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	dto := &client.AclRuleOrderingDto{}
	for _, id := range plan.RuleIDs {
		dto.RuleIDs = append(dto.RuleIDs, id.ValueString())
	}

	_, err := r.client.UpdateAclRuleOrdering(ctx, plan.SiteID.ValueString(), dto)
	if err != nil {
		resp.Diagnostics.AddError("Error Updating UniFi ACL Rule Ordering", err.Error())
		return
	}

	plan.ID = plan.SiteID
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *AclRuleOrderingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state AclRuleOrderingResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := r.client.GetAclRuleOrdering(ctx, state.SiteID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error Reading UniFi ACL Rule Ordering", err.Error())
		return
	}

	state.RuleIDs = make([]types.String, len(res.RuleIDs))
	for i, id := range res.RuleIDs {
		state.RuleIDs[i] = types.StringValue(id)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *AclRuleOrderingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan AclRuleOrderingResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	dto := &client.AclRuleOrderingDto{}
	for _, id := range plan.RuleIDs {
		dto.RuleIDs = append(dto.RuleIDs, id.ValueString())
	}

	_, err := r.client.UpdateAclRuleOrdering(ctx, plan.SiteID.ValueString(), dto)
	if err != nil {
		resp.Diagnostics.AddError("Error Updating UniFi ACL Rule Ordering", err.Error())
		return
	}

	plan.ID = plan.SiteID
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *AclRuleOrderingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Deletion logic for ordering could just mean resetting to empty or we just remove from state.
	// We'll just clear it out.
	var state AclRuleOrderingResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	dto := &client.AclRuleOrderingDto{RuleIDs: []string{}}
	_, err := r.client.UpdateAclRuleOrdering(ctx, state.SiteID.ValueString(), dto)
	if err != nil {
		resp.Diagnostics.AddError("Error Resetting UniFi ACL Rule Ordering", err.Error())
		return
	}
}
