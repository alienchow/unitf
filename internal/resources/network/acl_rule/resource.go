package network

import (
	"context"
	"fmt"
	"strings"

	"github.com/alienchow/unitf/internal/client"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type FirewallEndpointModel struct {
	ZoneID      types.String `tfsdk:"zone_id"`
	NetworkID   types.String `tfsdk:"network_id"`
	Address     types.String `tfsdk:"address"`
	Port        types.String `tfsdk:"port"`
	MacAddress  types.String `tfsdk:"mac_address"`
	MatchListID types.String `tfsdk:"match_list_id"`
}

func mapToEndpointDto(m *FirewallEndpointModel) *client.FirewallEndpointDto {
	if m == nil {
		return nil
	}
	dto := &client.FirewallEndpointDto{}
	if !m.ZoneID.IsNull() {
		dto.ZoneID = m.ZoneID.ValueString()
	}
	if !m.NetworkID.IsNull() {
		dto.NetworkID = m.NetworkID.ValueString()
	}
	if !m.Address.IsNull() {
		dto.Address = m.Address.ValueString()
	}
	if !m.Port.IsNull() {
		dto.Port = m.Port.ValueString()
	}
	if !m.MacAddress.IsNull() {
		dto.MACAddress = m.MacAddress.ValueString()
	}
	if !m.MatchListID.IsNull() {
		dto.MatchListID = m.MatchListID.ValueString()
	}
	return dto
}

func mapToEndpointModel(dto *client.FirewallEndpointDto) *FirewallEndpointModel {
	if dto == nil {
		return nil
	}
	m := &FirewallEndpointModel{}
	if dto.ZoneID != "" {
		m.ZoneID = types.StringValue(dto.ZoneID)
	}
	if dto.NetworkID != "" {
		m.NetworkID = types.StringValue(dto.NetworkID)
	}
	if dto.Address != "" {
		m.Address = types.StringValue(dto.Address)
	}
	if dto.Port != "" {
		m.Port = types.StringValue(dto.Port)
	}
	if dto.MACAddress != "" {
		m.MacAddress = types.StringValue(dto.MACAddress)
	}
	if dto.MatchListID != "" {
		m.MatchListID = types.StringValue(dto.MatchListID)
	}
	return m
}

var (
	_ resource.Resource                = &AclRuleResource{}
	_ resource.ResourceWithConfigure   = &AclRuleResource{}
	_ resource.ResourceWithImportState = &AclRuleResource{}
)

type AclRuleResource struct {
	client client.NetworkClient
}

type AclRuleResourceModel struct {
	ID          types.String           `tfsdk:"id"`
	SiteID      types.String           `tfsdk:"site_id"`
	Name        types.String           `tfsdk:"name"`
	Enabled     types.Bool             `tfsdk:"enabled"`
	Action      types.String           `tfsdk:"action"`
	IPVersion   types.String           `tfsdk:"ip_version"`
	Protocols   []types.String         `tfsdk:"protocols"`
	Source      *FirewallEndpointModel `tfsdk:"source"`
	Destination *FirewallEndpointModel `tfsdk:"destination"`
	Logging     types.Bool             `tfsdk:"logging"`
}

func NewAclRuleResource() resource.Resource {
	return &AclRuleResource{}
}

func (r *AclRuleResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_acl_rule"
}

func (r *AclRuleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages a UniFi ACL Rule.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "Unique ACL Rule UUID.",
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
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the ACL rule.",
			},
			"enabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(true),
				Description: "Is the rule enabled?",
			},
			"action": schema.StringAttribute{
				Required:    true,
				Description: "Action to take: ACCEPT, DROP, REJECT.",
				Validators: []validator.String{
					stringvalidator.OneOf("ACCEPT", "DROP", "REJECT"),
				},
			},
			"ip_version": schema.StringAttribute{
				Required:    true,
				Description: "IP version: IPV4, IPV6, IPV4_AND_IPV6, MAC.",
				Validators: []validator.String{
					stringvalidator.OneOf("IPV4", "IPV6", "IPV4_AND_IPV6", "MAC"),
				},
			},
			"protocols": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"logging": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
				Description: "Enable logging for this rule.",
			},
			"source": schema.SingleNestedAttribute{
				Optional:    true,
				Description: "Source endpoint matching.",
				Attributes: map[string]schema.Attribute{
					"zone_id":       schema.StringAttribute{Optional: true},
					"network_id":    schema.StringAttribute{Optional: true},
					"address":       schema.StringAttribute{Optional: true},
					"port":          schema.StringAttribute{Optional: true},
					"mac_address":   schema.StringAttribute{Optional: true},
					"match_list_id": schema.StringAttribute{Optional: true},
				},
			},
			"destination": schema.SingleNestedAttribute{
				Optional:    true,
				Description: "Destination endpoint matching.",
				Attributes: map[string]schema.Attribute{
					"zone_id":       schema.StringAttribute{Optional: true},
					"network_id":    schema.StringAttribute{Optional: true},
					"address":       schema.StringAttribute{Optional: true},
					"port":          schema.StringAttribute{Optional: true},
					"mac_address":   schema.StringAttribute{Optional: true},
					"match_list_id": schema.StringAttribute{Optional: true},
				},
			},
		},
	}
}

func (r *AclRuleResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *AclRuleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan AclRuleResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	dto := &client.AclRuleDto{
		Name:      plan.Name.ValueString(),
		Enabled:   plan.Enabled.ValueBool(),
		Action:    plan.Action.ValueString(),
		IPVersion: plan.IPVersion.ValueString(),
		Logging:   plan.Logging.ValueBool(),
	}

	for _, p := range plan.Protocols {
		dto.Protocols = append(dto.Protocols, p.ValueString())
	}

	dto.Source = mapToEndpointDto(plan.Source)
	dto.Destination = mapToEndpointDto(plan.Destination)

	res, err := r.client.CreateAclRule(ctx, plan.SiteID.ValueString(), dto)
	if err != nil {
		resp.Diagnostics.AddError("Error Creating UniFi ACL Rule", err.Error())
		return
	}

	plan.ID = types.StringValue(res.ID)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *AclRuleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state AclRuleResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := r.client.GetAclRule(ctx, state.SiteID.ValueString(), state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error Reading UniFi ACL Rule", err.Error())
		return
	}

	state.Name = types.StringValue(res.Name)
	state.Enabled = types.BoolValue(res.Enabled)
	state.Action = types.StringValue(res.Action)
	state.IPVersion = types.StringValue(res.IPVersion)
	state.Logging = types.BoolValue(res.Logging)

	state.Protocols = make([]types.String, len(res.Protocols))
	for i, p := range res.Protocols {
		state.Protocols[i] = types.StringValue(p)
	}

	state.Source = mapToEndpointModel(res.Source)
	state.Destination = mapToEndpointModel(res.Destination)

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *AclRuleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan AclRuleResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	dto := &client.AclRuleDto{
		Name:      plan.Name.ValueString(),
		Enabled:   plan.Enabled.ValueBool(),
		Action:    plan.Action.ValueString(),
		IPVersion: plan.IPVersion.ValueString(),
		Logging:   plan.Logging.ValueBool(),
	}

	for _, p := range plan.Protocols {
		dto.Protocols = append(dto.Protocols, p.ValueString())
	}

	dto.Source = mapToEndpointDto(plan.Source)
	dto.Destination = mapToEndpointDto(plan.Destination)

	res, err := r.client.UpdateAclRule(ctx, plan.SiteID.ValueString(), plan.ID.ValueString(), dto)
	if err != nil {
		resp.Diagnostics.AddError("Error Updating UniFi ACL Rule", err.Error())
		return
	}

	plan.ID = types.StringValue(res.ID)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *AclRuleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state AclRuleResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteAclRule(ctx, state.SiteID.ValueString(), state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error Deleting UniFi ACL Rule", err.Error())
		return
	}
}

func (r *AclRuleResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, "/")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		resp.Diagnostics.AddError("Invalid Import Identifier", fmt.Sprintf("Expected 'site_id/rule_id', got %q", req.ID))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("site_id"), parts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), parts[1])...)
}
