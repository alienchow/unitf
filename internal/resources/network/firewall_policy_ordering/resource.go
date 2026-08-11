package network

import (
	"context"
	"fmt"
	"strings"

	"github.com/alienchow/unitf/internal/client"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                = &FirewallPolicyOrderingResource{}
	_ resource.ResourceWithConfigure   = &FirewallPolicyOrderingResource{}
	_ resource.ResourceWithImportState = &FirewallPolicyOrderingResource{}
)

type FirewallPolicyOrderingResource struct {
	client client.NetworkClient
}

type FirewallPolicyOrderingResourceModel struct {
	ID                  types.String   `tfsdk:"id"`
	SiteID              types.String   `tfsdk:"site_id"`
	FromZoneID          types.String   `tfsdk:"from_zone_id"`
	ToZoneID            types.String   `tfsdk:"to_zone_id"`
	BeforeSystemDefined []types.String `tfsdk:"before_system_defined"`
	AfterSystemDefined  []types.String `tfsdk:"after_system_defined"`
}

func NewFirewallPolicyOrderingResource() resource.Resource {
	return &FirewallPolicyOrderingResource{}
}

func (r *FirewallPolicyOrderingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_firewall_policy_ordering"
}

func (r *FirewallPolicyOrderingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages UniFi Firewall Policy Ordering for a pair of zones.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "Internal ID of the ordering resource.",
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
			"from_zone_id": schema.StringAttribute{
				Required:    true,
				Description: "ID of the source firewall zone.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"to_zone_id": schema.StringAttribute{
				Required:    true,
				Description: "ID of the destination firewall zone.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"before_system_defined": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "List of policy IDs to apply before system-defined rules.",
			},
			"after_system_defined": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "List of policy IDs to apply after system-defined rules.",
			},
		},
	}
}

func (r *FirewallPolicyOrderingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Resource Configure Type", "Expected *client.Client")
		return
	}
	r.client = c.Network
}

func (r *FirewallPolicyOrderingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan FirewallPolicyOrderingResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	dto := r.modelToDto(plan)

	res, err := r.client.UpdateFirewallPolicyOrdering(ctx, plan.SiteID.ValueString(), plan.FromZoneID.ValueString(), plan.ToZoneID.ValueString(), dto)
	if err != nil {
		resp.Diagnostics.AddError("Error Creating/Updating UniFi Firewall Policy Ordering", err.Error())
		return
	}

	state := r.dtoToModel(plan.SiteID.ValueString(), plan.FromZoneID.ValueString(), plan.ToZoneID.ValueString(), res)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *FirewallPolicyOrderingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state FirewallPolicyOrderingResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := r.client.GetFirewallPolicyOrdering(ctx, state.SiteID.ValueString(), state.FromZoneID.ValueString(), state.ToZoneID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error Reading UniFi Firewall Policy Ordering", err.Error())
		return
	}

	newState := r.dtoToModel(state.SiteID.ValueString(), state.FromZoneID.ValueString(), state.ToZoneID.ValueString(), res)
	resp.Diagnostics.Append(resp.State.Set(ctx, &newState)...)
}

func (r *FirewallPolicyOrderingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan FirewallPolicyOrderingResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	dto := r.modelToDto(plan)

	res, err := r.client.UpdateFirewallPolicyOrdering(ctx, plan.SiteID.ValueString(), plan.FromZoneID.ValueString(), plan.ToZoneID.ValueString(), dto)
	if err != nil {
		resp.Diagnostics.AddError("Error Updating UniFi Firewall Policy Ordering", err.Error())
		return
	}

	state := r.dtoToModel(plan.SiteID.ValueString(), plan.FromZoneID.ValueString(), plan.ToZoneID.ValueString(), res)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *FirewallPolicyOrderingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state FirewallPolicyOrderingResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// For ordering, deletion typically implies clearing the order list
	dto := &client.FirewallPolicyOrderingDto{
		BeforeSystemDefined: []string{},
		AfterSystemDefined:  []string{},
	}

	_, err := r.client.UpdateFirewallPolicyOrdering(ctx, state.SiteID.ValueString(), state.FromZoneID.ValueString(), state.ToZoneID.ValueString(), dto)
	if err != nil && !client.IsNotFound(err) {
		resp.Diagnostics.AddError("Error Deleting/Clearing UniFi Firewall Policy Ordering", err.Error())
		return
	}
}

func (r *FirewallPolicyOrderingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, "/")
	if len(parts) != 3 || parts[0] == "" || parts[1] == "" || parts[2] == "" {
		resp.Diagnostics.AddError("Invalid Import Identifier", fmt.Sprintf("Expected 'site_id/from_zone_id/to_zone_id', got %q", req.ID))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("site_id"), parts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("from_zone_id"), parts[1])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("to_zone_id"), parts[2])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}

func (r *FirewallPolicyOrderingResource) dtoToModel(siteID, fromZoneID, toZoneID string, res *client.FirewallPolicyOrderingDto) FirewallPolicyOrderingResourceModel {
	state := FirewallPolicyOrderingResourceModel{
		ID:         types.StringValue(fmt.Sprintf("%s/%s/%s", siteID, fromZoneID, toZoneID)),
		SiteID:     types.StringValue(siteID),
		FromZoneID: types.StringValue(fromZoneID),
		ToZoneID:   types.StringValue(toZoneID),
	}

	state.BeforeSystemDefined = make([]types.String, len(res.BeforeSystemDefined))
	for i, id := range res.BeforeSystemDefined {
		state.BeforeSystemDefined[i] = types.StringValue(id)
	}

	state.AfterSystemDefined = make([]types.String, len(res.AfterSystemDefined))
	for i, id := range res.AfterSystemDefined {
		state.AfterSystemDefined[i] = types.StringValue(id)
	}

	return state
}

func (r *FirewallPolicyOrderingResource) modelToDto(model FirewallPolicyOrderingResourceModel) *client.FirewallPolicyOrderingDto {
	dto := &client.FirewallPolicyOrderingDto{}
	for _, id := range model.BeforeSystemDefined {
		if !id.IsNull() && !id.IsUnknown() {
			dto.BeforeSystemDefined = append(dto.BeforeSystemDefined, id.ValueString())
		}
	}
	for _, id := range model.AfterSystemDefined {
		if !id.IsNull() && !id.IsUnknown() {
			dto.AfterSystemDefined = append(dto.AfterSystemDefined, id.ValueString())
		}
	}
	return dto
}
