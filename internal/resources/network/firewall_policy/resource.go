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

var (
	_ resource.Resource                = &FirewallPolicyResource{}
	_ resource.ResourceWithConfigure   = &FirewallPolicyResource{}
	_ resource.ResourceWithImportState = &FirewallPolicyResource{}
)

type FirewallPolicyResource struct {
	client client.NetworkClient
}

type FirewallPolicyResourceModel struct {
	ID              types.String           `tfsdk:"id"`
	SiteID          types.String           `tfsdk:"site_id"`
	Name            types.String           `tfsdk:"name"`
	Enabled         types.Bool             `tfsdk:"enabled"`
	Action          *FirewallActionModel   `tfsdk:"action"`
	IPProtocolScope *IPProtocolScopeModel  `tfsdk:"ip_protocol_scope"`
	Source          *FirewallEndpointModel `tfsdk:"source"`
	Destination     *FirewallEndpointModel `tfsdk:"destination"`
	Logging         types.Bool             `tfsdk:"logging"`
}

type FirewallActionModel struct {
	Accept *struct{} `tfsdk:"accept"`
	Block  *struct{} `tfsdk:"block"`
	Drop   *struct{} `tfsdk:"drop"`
	Reject *struct{} `tfsdk:"reject"`
}

type IPProtocolScopeModel struct {
	IPVersion types.String   `tfsdk:"ip_version"`
	Protocols []types.String `tfsdk:"protocols"`
}

type FirewallEndpointModel struct {
	ZoneID      types.String `tfsdk:"zone_id"`
	NetworkID   types.String `tfsdk:"network_id"`
	Address     types.String `tfsdk:"address"`
	Port        types.String `tfsdk:"port"`
	MACAddress  types.String `tfsdk:"mac_address"`
	MatchListID types.String `tfsdk:"match_list_id"`
}

func NewFirewallPolicyResource() resource.Resource {
	return &FirewallPolicyResource{}
}

func (r *FirewallPolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_firewall_policy"
}

func (r *FirewallPolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages a UniFi Firewall Policy.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "Unique Firewall Policy UUID.",
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
				Description: "Name of the firewall policy.",
			},
			"enabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(true),
				Description: "Is the policy enabled?",
			},
			"logging": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
				Description: "Enable logging for this policy.",
			},
			"action": schema.SingleNestedAttribute{
				Required:    true,
				Description: "Action to take. Exactly one block must be defined.",
				Attributes: map[string]schema.Attribute{
					"accept": schema.SingleNestedAttribute{
						Optional:   true,
						Attributes: map[string]schema.Attribute{},
					},
					"block": schema.SingleNestedAttribute{
						Optional:   true,
						Attributes: map[string]schema.Attribute{},
					},
					"drop": schema.SingleNestedAttribute{
						Optional:   true,
						Attributes: map[string]schema.Attribute{},
					},
					"reject": schema.SingleNestedAttribute{
						Optional:   true,
						Attributes: map[string]schema.Attribute{},
					},
				},
			},
			"ip_protocol_scope": schema.SingleNestedAttribute{
				Optional:    true,
				Description: "IP version and protocol matching.",
				Attributes: map[string]schema.Attribute{
					"ip_version": schema.StringAttribute{
						Required: true,
						Validators: []validator.String{
							stringvalidator.OneOf("IPV4", "IPV6", "IPV4_AND_IPV6"),
						},
					},
					"protocols": schema.ListAttribute{
						ElementType: types.StringType,
						Optional:    true,
						Computed:    true,
					},
				},
			},
			"source": schema.SingleNestedAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Source endpoint matching.",
				Attributes: map[string]schema.Attribute{
					"zone_id":       schema.StringAttribute{Optional: true, Computed: true},
					"network_id":    schema.StringAttribute{Optional: true, Computed: true},
					"address":       schema.StringAttribute{Optional: true, Computed: true},
					"port":          schema.StringAttribute{Optional: true, Computed: true},
					"mac_address":   schema.StringAttribute{Optional: true, Computed: true},
					"match_list_id": schema.StringAttribute{Optional: true, Computed: true},
				},
			},
			"destination": schema.SingleNestedAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Destination endpoint matching.",
				Attributes: map[string]schema.Attribute{
					"zone_id":       schema.StringAttribute{Optional: true, Computed: true},
					"network_id":    schema.StringAttribute{Optional: true, Computed: true},
					"address":       schema.StringAttribute{Optional: true, Computed: true},
					"port":          schema.StringAttribute{Optional: true, Computed: true},
					"mac_address":   schema.StringAttribute{Optional: true, Computed: true},
					"match_list_id": schema.StringAttribute{Optional: true, Computed: true},
				},
			},
		},
	}
}

func (r *FirewallPolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func mapToEndpointDto(model *FirewallEndpointModel) *client.FirewallEndpointDto {
	if model == nil {
		return nil
	}
	dto := &client.FirewallEndpointDto{}
	if !model.ZoneID.IsNull() && !model.ZoneID.IsUnknown() {
		dto.ZoneID = model.ZoneID.ValueString()
	}
	if !model.NetworkID.IsNull() && !model.NetworkID.IsUnknown() {
		dto.NetworkID = model.NetworkID.ValueString()
	}
	if !model.Address.IsNull() && !model.Address.IsUnknown() {
		dto.Address = model.Address.ValueString()
	}
	if !model.Port.IsNull() && !model.Port.IsUnknown() {
		dto.Port = model.Port.ValueString()
	}
	if !model.MACAddress.IsNull() && !model.MACAddress.IsUnknown() {
		dto.MACAddress = model.MACAddress.ValueString()
	}
	if !model.MatchListID.IsNull() && !model.MatchListID.IsUnknown() {
		dto.MatchListID = model.MatchListID.ValueString()
	}
	return dto
}

func mapToEndpointModel(dto *client.FirewallEndpointDto) *FirewallEndpointModel {
	if dto == nil {
		return nil
	}
	// Simplified to check if empty
	if dto.ZoneID == "" && dto.NetworkID == "" && dto.Address == "" && dto.Port == "" && dto.MACAddress == "" && dto.MatchListID == "" {
		return nil
	}

	model := &FirewallEndpointModel{}
	if dto.ZoneID != "" {
		model.ZoneID = types.StringValue(dto.ZoneID)
	}
	if dto.NetworkID != "" {
		model.NetworkID = types.StringValue(dto.NetworkID)
	}
	if dto.Address != "" {
		model.Address = types.StringValue(dto.Address)
	}
	if dto.Port != "" {
		model.Port = types.StringValue(dto.Port)
	}
	if dto.MACAddress != "" {
		model.MACAddress = types.StringValue(dto.MACAddress)
	}
	if dto.MatchListID != "" {
		model.MatchListID = types.StringValue(dto.MatchListID)
	}
	return model
}

func (r *FirewallPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan FirewallPolicyResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	dto := r.modelToDto(plan)

	res, err := r.client.CreateFirewallPolicy(ctx, plan.SiteID.ValueString(), dto)
	if err != nil {
		resp.Diagnostics.AddError("Error Creating UniFi Firewall Policy", err.Error())
		return
	}

	state := r.dtoToModel(plan.SiteID.ValueString(), res)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *FirewallPolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state FirewallPolicyResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := r.client.GetFirewallPolicy(ctx, state.SiteID.ValueString(), state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error Reading UniFi Firewall Policy", err.Error())
		return
	}

	newState := r.dtoToModel(state.SiteID.ValueString(), res)
	resp.Diagnostics.Append(resp.State.Set(ctx, &newState)...)
}

func (r *FirewallPolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan FirewallPolicyResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	dto := r.modelToDto(plan)

	res, err := r.client.UpdateFirewallPolicy(ctx, plan.SiteID.ValueString(), plan.ID.ValueString(), dto)
	if err != nil {
		resp.Diagnostics.AddError("Error Updating UniFi Firewall Policy", err.Error())
		return
	}

	state := r.dtoToModel(plan.SiteID.ValueString(), res)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *FirewallPolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state FirewallPolicyResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteFirewallPolicy(ctx, state.SiteID.ValueString(), state.ID.ValueString())
	if err != nil && !client.IsNotFound(err) {
		resp.Diagnostics.AddError("Error Deleting UniFi Firewall Policy", err.Error())
		return
	}
}

func (r *FirewallPolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, "/")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		resp.Diagnostics.AddError("Invalid Import Identifier", fmt.Sprintf("Expected 'site_id/policy_id', got %q", req.ID))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("site_id"), parts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), parts[1])...)
}

func (r *FirewallPolicyResource) dtoToModel(siteID string, res *client.FirewallPolicyDto) FirewallPolicyResourceModel {
	state := FirewallPolicyResourceModel{
		ID:      types.StringValue(res.ID),
		SiteID:  types.StringValue(siteID),
		Name:    types.StringValue(res.Name),
		Enabled: types.BoolValue(res.Enabled),
		Logging: types.BoolValue(res.Logging),
	}

	if res.Action != nil {
		state.Action = &FirewallActionModel{}
		if res.Action.Accept != nil {
			state.Action.Accept = &struct{}{}
		}
		if res.Action.Block != nil {
			state.Action.Block = &struct{}{}
		}
		if res.Action.Drop != nil {
			state.Action.Drop = &struct{}{}
		}
		if res.Action.Reject != nil {
			state.Action.Reject = &struct{}{}
		}
	}

	if res.IPProtocolScope != nil {
		state.IPProtocolScope = &IPProtocolScopeModel{
			IPVersion: types.StringValue(res.IPProtocolScope.IPVersion),
		}
		for _, p := range res.IPProtocolScope.Protocols {
			state.IPProtocolScope.Protocols = append(state.IPProtocolScope.Protocols, types.StringValue(p))
		}
	}

	state.Source = mapToEndpointModel(res.Source)
	state.Destination = mapToEndpointModel(res.Destination)

	return state
}

func (r *FirewallPolicyResource) modelToDto(model FirewallPolicyResourceModel) *client.FirewallPolicyDto {
	dto := &client.FirewallPolicyDto{}
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		dto.Name = model.Name.ValueString()
	}
	if !model.Enabled.IsNull() && !model.Enabled.IsUnknown() {
		dto.Enabled = model.Enabled.ValueBool()
	}
	if !model.Logging.IsNull() && !model.Logging.IsUnknown() {
		dto.Logging = model.Logging.ValueBool()
	}

	if model.Action != nil {
		dto.Action = &client.FirewallActionDto{}
		if model.Action.Accept != nil {
			dto.Action.Accept = &struct{}{}
		}
		if model.Action.Block != nil {
			dto.Action.Block = &struct{}{}
		}
		if model.Action.Drop != nil {
			dto.Action.Drop = &struct{}{}
		}
		if model.Action.Reject != nil {
			dto.Action.Reject = &struct{}{}
		}
	}

	if model.IPProtocolScope != nil {
		dto.IPProtocolScope = &client.IPProtocolScopeDto{}
		if !model.IPProtocolScope.IPVersion.IsNull() && !model.IPProtocolScope.IPVersion.IsUnknown() {
			dto.IPProtocolScope.IPVersion = model.IPProtocolScope.IPVersion.ValueString()
		}
		for _, p := range model.IPProtocolScope.Protocols {
			if !p.IsNull() && !p.IsUnknown() {
				dto.IPProtocolScope.Protocols = append(dto.IPProtocolScope.Protocols, p.ValueString())
			}
		}
	}

	dto.Source = mapToEndpointDto(model.Source)
	dto.Destination = mapToEndpointDto(model.Destination)

	return dto
}
