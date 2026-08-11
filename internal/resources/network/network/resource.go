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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                = &NetworkResource{}
	_ resource.ResourceWithConfigure   = &NetworkResource{}
	_ resource.ResourceWithImportState = &NetworkResource{}
)

type NetworkResource struct {
	client client.NetworkClient
}

type NetworkResourceModel struct {
	ID             types.String                `tfsdk:"id"`
	SiteID         types.String                `tfsdk:"site_id"`
	Name           types.String                `tfsdk:"name"`
	Type           types.String                `tfsdk:"type"`
	GatewayManaged *GatewayManagedNetworkModel `tfsdk:"gateway_managed"`
	SwitchManaged  *SwitchManagedNetworkModel  `tfsdk:"switch_managed"`
	Unmanaged      *UnmanagedNetworkModel      `tfsdk:"unmanaged"`
}

type GatewayManagedNetworkModel struct {
	VlanID       types.Int64        `tfsdk:"vlan_id"`
	Purpose      types.String       `tfsdk:"purpose"`
	MulticastDNS types.Bool         `tfsdk:"multicast_dns"`
	DHCPGuarding *DHCPGuardingModel `tfsdk:"dhcp_guarding"`
	IPv4         *GatewayIPv4Model  `tfsdk:"ipv4"`
}

type SwitchManagedNetworkModel struct {
	VlanID       types.Int64        `tfsdk:"vlan_id"`
	DHCPGuarding *DHCPGuardingModel `tfsdk:"dhcp_guarding"`
}

type UnmanagedNetworkModel struct {
	VlanID types.Int64 `tfsdk:"vlan_id"`
}

type DHCPGuardingModel struct {
	Enabled   types.Bool     `tfsdk:"enabled"`
	ServerIPs []types.String `tfsdk:"server_ips"`
}

type GatewayIPv4Model struct {
	Enabled    types.Bool       `tfsdk:"enabled"`
	SubnetMask types.String     `tfsdk:"subnet_mask"`
	DHCPServer *DHCPServerModel `tfsdk:"dhcp_server"`
}

type DHCPServerModel struct {
	RangeStart     types.String   `tfsdk:"range_start"`
	RangeStop      types.String   `tfsdk:"range_stop"`
	LeaseTimeSec   types.Int64    `tfsdk:"lease_time_sec"`
	GatewayAddress types.String   `tfsdk:"gateway_address"`
	DNSServers     []types.String `tfsdk:"dns_servers"`
}

func NewNetworkResource() resource.Resource {
	return &NetworkResource{}
}

func (r *NetworkResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network"
}

func (r *NetworkResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages a UniFi Network.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "Unique Network UUID.",
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
				Description: "Name of the network.",
			},
			"type": schema.StringAttribute{
				Required:    true,
				Description: "Type of network: GATEWAY_MANAGED, SWITCH_MANAGED, UNMANAGED.",
				Validators: []validator.String{
					stringvalidator.OneOf("GATEWAY_MANAGED", "SWITCH_MANAGED", "UNMANAGED"),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"gateway_managed": schema.SingleNestedAttribute{
				Optional:    true,
				Description: "Configuration for GATEWAY_MANAGED networks.",
				Attributes: map[string]schema.Attribute{
					"vlan_id": schema.Int64Attribute{
						Optional:    true,
						Description: "VLAN ID (1-4094).",
					},
					"purpose": schema.StringAttribute{
						Optional:    true,
						Description: "Network purpose: CORPORATE, GUEST, etc.",
					},
					"multicast_dns": schema.BoolAttribute{
						Optional:    true,
						Description: "Enable mDNS forwarding.",
					},
					"dhcp_guarding": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{Required: true},
							"server_ips": schema.ListAttribute{
								ElementType: types.StringType,
								Optional:    true,
							},
						},
					},
					"ipv4": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"enabled":     schema.BoolAttribute{Required: true},
							"subnet_mask": schema.StringAttribute{Required: true},
							"dhcp_server": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"range_start":     schema.StringAttribute{Optional: true},
									"range_stop":      schema.StringAttribute{Optional: true},
									"lease_time_sec":  schema.Int64Attribute{Optional: true},
									"gateway_address": schema.StringAttribute{Optional: true},
									"dns_servers": schema.ListAttribute{
										ElementType: types.StringType,
										Optional:    true,
									},
								},
							},
						},
					},
				},
			},
			"switch_managed": schema.SingleNestedAttribute{
				Optional:    true,
				Description: "Configuration for SWITCH_MANAGED networks.",
				Attributes: map[string]schema.Attribute{
					"vlan_id": schema.Int64Attribute{Required: true},
					"dhcp_guarding": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{Required: true},
							"server_ips": schema.ListAttribute{
								ElementType: types.StringType,
								Optional:    true,
							},
						},
					},
				},
			},
			"unmanaged": schema.SingleNestedAttribute{
				Optional:    true,
				Description: "Configuration for UNMANAGED networks.",
				Attributes: map[string]schema.Attribute{
					"vlan_id": schema.Int64Attribute{Optional: true},
				},
			},
		},
	}
}

func (r *NetworkResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *NetworkResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan NetworkResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	dto := buildNetworkDto(plan)

	res, err := r.client.CreateNetwork(ctx, plan.SiteID.ValueString(), dto)
	if err != nil {
		resp.Diagnostics.AddError("Error Creating UniFi Network", err.Error())
		return
	}

	plan.ID = types.StringValue(res.ID)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *NetworkResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state NetworkResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := r.client.GetNetwork(ctx, state.SiteID.ValueString(), state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error Reading UniFi Network", err.Error())
		return
	}

	populateStateFromDto(&state, res)

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *NetworkResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan NetworkResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	dto := buildNetworkDto(plan)

	res, err := r.client.UpdateNetwork(ctx, plan.SiteID.ValueString(), plan.ID.ValueString(), dto)
	if err != nil {
		resp.Diagnostics.AddError("Error Updating UniFi Network", err.Error())
		return
	}

	plan.ID = types.StringValue(res.ID)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *NetworkResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state NetworkResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteNetwork(ctx, state.SiteID.ValueString(), state.ID.ValueString())
	if err != nil && !client.IsNotFound(err) {
		resp.Diagnostics.AddError("Error Deleting UniFi Network", err.Error())
		return
	}
}

func (r *NetworkResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, "/")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		resp.Diagnostics.AddError("Invalid Import Identifier", fmt.Sprintf("Expected 'site_id/network_id', got %q", req.ID))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("site_id"), parts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), parts[1])...)
}

func buildNetworkDto(plan NetworkResourceModel) *client.NetworkDto {
	dto := &client.NetworkDto{
		Name: plan.Name.ValueString(),
		Type: plan.Type.ValueString(),
	}

	switch plan.Type.ValueString() {
	case "GATEWAY_MANAGED":
		if plan.GatewayManaged != nil {
			if !plan.GatewayManaged.VlanID.IsNull() {
				vlan := int(plan.GatewayManaged.VlanID.ValueInt64())
				dto.VlanID = &vlan
			}
			if !plan.GatewayManaged.Purpose.IsNull() {
				dto.Purpose = plan.GatewayManaged.Purpose.ValueString()
			}
			if !plan.GatewayManaged.MulticastDNS.IsNull() {
				mdns := plan.GatewayManaged.MulticastDNS.ValueBool()
				dto.MulticastDNS = &mdns
			}
			if plan.GatewayManaged.DHCPGuarding != nil {
				dto.DHCPGuarding = &client.NetworkDHCPGuarding{
					Enabled: plan.GatewayManaged.DHCPGuarding.Enabled.ValueBool(),
				}
				for _, ip := range plan.GatewayManaged.DHCPGuarding.ServerIPs {
					if !ip.IsNull() {
						dto.DHCPGuarding.ServerIPs = append(dto.DHCPGuarding.ServerIPs, ip.ValueString())
					}
				}
			}
			if plan.GatewayManaged.IPv4 != nil {
				dto.IPv4 = &client.GatewayIPv4Config{
					Enabled:    plan.GatewayManaged.IPv4.Enabled.ValueBool(),
					SubnetMask: plan.GatewayManaged.IPv4.SubnetMask.ValueString(),
				}
				if plan.GatewayManaged.IPv4.DHCPServer != nil {
					dto.IPv4.DHCP = &client.IPv4DHCPDto{
						Mode:       "DHCP_SERVER",
						RangeStart: plan.GatewayManaged.IPv4.DHCPServer.RangeStart.ValueString(),
						RangeStop:  plan.GatewayManaged.IPv4.DHCPServer.RangeStop.ValueString(),
					}
					if !plan.GatewayManaged.IPv4.DHCPServer.LeaseTimeSec.IsNull() {
						lease := plan.GatewayManaged.IPv4.DHCPServer.LeaseTimeSec.ValueInt64()
						dto.IPv4.DHCP.LeaseTimeSec = &lease
					}
					if !plan.GatewayManaged.IPv4.DHCPServer.GatewayAddress.IsNull() {
						dto.IPv4.DHCP.GatewayAddress = plan.GatewayManaged.IPv4.DHCPServer.GatewayAddress.ValueString()
					}
					for _, ip := range plan.GatewayManaged.IPv4.DHCPServer.DNSServers {
						if !ip.IsNull() {
							dto.IPv4.DHCP.DNSServers = append(dto.IPv4.DHCP.DNSServers, ip.ValueString())
						}
					}
				}
			}
		}
	case "SWITCH_MANAGED":
		if plan.SwitchManaged != nil {
			if !plan.SwitchManaged.VlanID.IsNull() {
				vlan := int(plan.SwitchManaged.VlanID.ValueInt64())
				dto.VlanID = &vlan
			}
			if plan.SwitchManaged.DHCPGuarding != nil {
				dto.DHCPGuarding = &client.NetworkDHCPGuarding{
					Enabled: plan.SwitchManaged.DHCPGuarding.Enabled.ValueBool(),
				}
				for _, ip := range plan.SwitchManaged.DHCPGuarding.ServerIPs {
					if !ip.IsNull() {
						dto.DHCPGuarding.ServerIPs = append(dto.DHCPGuarding.ServerIPs, ip.ValueString())
					}
				}
			}
		}
	case "UNMANAGED":
		if plan.Unmanaged != nil {
			if !plan.Unmanaged.VlanID.IsNull() {
				vlan := int(plan.Unmanaged.VlanID.ValueInt64())
				dto.VlanID = &vlan
			}
		}
	}

	return dto
}

func populateStateFromDto(state *NetworkResourceModel, res *client.NetworkDto) {
	state.Name = types.StringValue(res.Name)
	if res.Type == "" {
		state.Type = types.StringNull()
	} else {
		state.Type = types.StringValue(res.Type)
	}

	state.GatewayManaged = nil
	state.SwitchManaged = nil
	state.Unmanaged = nil

	switch res.Type {
	case "GATEWAY_MANAGED":
		gm := &GatewayManagedNetworkModel{
			VlanID:       types.Int64Null(),
			Purpose:      types.StringNull(),
			MulticastDNS: types.BoolNull(),
		}
		if res.VlanID != nil {
			gm.VlanID = types.Int64Value(int64(*res.VlanID))
		}
		if res.Purpose != "" {
			gm.Purpose = types.StringValue(res.Purpose)
		}
		if res.MulticastDNS != nil {
			gm.MulticastDNS = types.BoolValue(*res.MulticastDNS)
		}
		if res.DHCPGuarding != nil {
			dg := &DHCPGuardingModel{
				Enabled: types.BoolValue(res.DHCPGuarding.Enabled),
			}
			for _, ip := range res.DHCPGuarding.ServerIPs {
				dg.ServerIPs = append(dg.ServerIPs, types.StringValue(ip))
			}
			gm.DHCPGuarding = dg
		}
		if res.IPv4 != nil {
			ipv4 := &GatewayIPv4Model{
				Enabled:    types.BoolValue(res.IPv4.Enabled),
				SubnetMask: types.StringValue(res.IPv4.SubnetMask),
			}
			if res.IPv4.DHCP != nil && res.IPv4.DHCP.Mode == "DHCP_SERVER" {
				dhcp := &DHCPServerModel{
					RangeStart:     types.StringNull(),
					RangeStop:      types.StringNull(),
					LeaseTimeSec:   types.Int64Null(),
					GatewayAddress: types.StringNull(),
				}
				if res.IPv4.DHCP.RangeStart != "" {
					dhcp.RangeStart = types.StringValue(res.IPv4.DHCP.RangeStart)
				}
				if res.IPv4.DHCP.RangeStop != "" {
					dhcp.RangeStop = types.StringValue(res.IPv4.DHCP.RangeStop)
				}
				if res.IPv4.DHCP.LeaseTimeSec != nil {
					dhcp.LeaseTimeSec = types.Int64Value(*res.IPv4.DHCP.LeaseTimeSec)
				}
				if res.IPv4.DHCP.GatewayAddress != "" {
					dhcp.GatewayAddress = types.StringValue(res.IPv4.DHCP.GatewayAddress)
				}
				for _, dns := range res.IPv4.DHCP.DNSServers {
					dhcp.DNSServers = append(dhcp.DNSServers, types.StringValue(dns))
				}
				ipv4.DHCPServer = dhcp
			}
			gm.IPv4 = ipv4
		}
		state.GatewayManaged = gm

	case "SWITCH_MANAGED":
		sm := &SwitchManagedNetworkModel{
			VlanID: types.Int64Null(),
		}
		if res.VlanID != nil {
			sm.VlanID = types.Int64Value(int64(*res.VlanID))
		}
		if res.DHCPGuarding != nil {
			dg := &DHCPGuardingModel{
				Enabled: types.BoolValue(res.DHCPGuarding.Enabled),
			}
			for _, ip := range res.DHCPGuarding.ServerIPs {
				dg.ServerIPs = append(dg.ServerIPs, types.StringValue(ip))
			}
			sm.DHCPGuarding = dg
		}
		state.SwitchManaged = sm

	case "UNMANAGED":
		um := &UnmanagedNetworkModel{
			VlanID: types.Int64Null(),
		}
		if res.VlanID != nil {
			um.VlanID = types.Int64Value(int64(*res.VlanID))
		}
		state.Unmanaged = um
	}
}
