package network

import (
	"context"
	"fmt"
	"strings"

	"github.com/alienchow/unitf/internal/client"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                = &WifiBroadcastResource{}
	_ resource.ResourceWithConfigure   = &WifiBroadcastResource{}
	_ resource.ResourceWithImportState = &WifiBroadcastResource{}
)

type WifiBroadcastResource struct {
	client client.NetworkClient
}

type WifiBroadcastResourceModel struct {
	ID         types.String `tfsdk:"id"`
	SiteID     types.String `tfsdk:"site_id"`
	Name       types.String `tfsdk:"name"`
	Enabled    types.Bool   `tfsdk:"enabled"`
	SSID       types.String `tfsdk:"ssid"`
	Security   types.String `tfsdk:"security"`
	Passphrase types.String `tfsdk:"passphrase"`
	NetworkID  types.String `tfsdk:"network_id"`
	Mode       types.String `tfsdk:"mode"`
}

func NewWifiBroadcastResource() resource.Resource {
	return &WifiBroadcastResource{}
}

func (r *WifiBroadcastResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wifi_broadcast"
}

func (r *WifiBroadcastResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages a UniFi WiFi Broadcast.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "Unique WiFi Broadcast UUID.",
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
				Description: "Name of the WiFi broadcast.",
			},
			"enabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(true),
				Description: "Is the WiFi broadcast enabled?",
			},
			"ssid": schema.StringAttribute{
				Required:    true,
				Description: "SSID of the WiFi broadcast.",
			},
			"security": schema.StringAttribute{
				Required:    true,
				Description: "Security protocol (e.g., WPA2, WPA3).",
			},
			"passphrase": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "WiFi passphrase. Required for WPA2/WPA3.",
			},
			"network_id": schema.StringAttribute{
				Required:    true,
				Description: "ID of the network to bridge this WiFi to.",
			},
			"mode": schema.StringAttribute{
				Optional:    true,
				Description: "Mode of the WiFi broadcast (e.g., STANDARD, IOT).",
			},
		},
	}
}

func (r *WifiBroadcastResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *WifiBroadcastResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan WifiBroadcastResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	dto := &client.WifiBroadcastDto{
		Name:      plan.Name.ValueString(),
		Enabled:   plan.Enabled.ValueBool(),
		SSID:      plan.SSID.ValueString(),
		Security:  plan.Security.ValueString(),
		NetworkID: plan.NetworkID.ValueString(),
	}
	if !plan.Passphrase.IsNull() {
		dto.Passphrase = plan.Passphrase.ValueString()
	}
	if !plan.Mode.IsNull() {
		dto.Mode = plan.Mode.ValueString()
	}

	res, err := r.client.CreateWifiBroadcast(ctx, plan.SiteID.ValueString(), dto)
	if err != nil {
		resp.Diagnostics.AddError("Error Creating UniFi WiFi Broadcast", err.Error())
		return
	}

	plan.ID = types.StringValue(res.ID)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *WifiBroadcastResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state WifiBroadcastResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := r.client.GetWifiBroadcast(ctx, state.SiteID.ValueString(), state.ID.ValueString())
	if err != nil {
		// Mock logic for not found check can go here if IsNotFound is defined
		resp.Diagnostics.AddError("Error Reading UniFi WiFi Broadcast", err.Error())
		return
	}

	state.Name = types.StringValue(res.Name)
	state.Enabled = types.BoolValue(res.Enabled)
	state.SSID = types.StringValue(res.SSID)
	state.Security = types.StringValue(res.Security)
	state.NetworkID = types.StringValue(res.NetworkID)

	if res.Mode != "" {
		state.Mode = types.StringValue(res.Mode)
	} else {
		state.Mode = types.StringNull()
	}

	// We typically don't read back Passphrase because APIs often redact it.
	// But if the API returns it, we can set it. Otherwise rely on state.

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *WifiBroadcastResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan WifiBroadcastResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	dto := &client.WifiBroadcastDto{
		Name:      plan.Name.ValueString(),
		Enabled:   plan.Enabled.ValueBool(),
		SSID:      plan.SSID.ValueString(),
		Security:  plan.Security.ValueString(),
		NetworkID: plan.NetworkID.ValueString(),
	}
	if !plan.Passphrase.IsNull() {
		dto.Passphrase = plan.Passphrase.ValueString()
	}
	if !plan.Mode.IsNull() {
		dto.Mode = plan.Mode.ValueString()
	}

	res, err := r.client.UpdateWifiBroadcast(ctx, plan.SiteID.ValueString(), plan.ID.ValueString(), dto)
	if err != nil {
		resp.Diagnostics.AddError("Error Updating UniFi WiFi Broadcast", err.Error())
		return
	}

	plan.ID = types.StringValue(res.ID)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *WifiBroadcastResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state WifiBroadcastResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteWifiBroadcast(ctx, state.SiteID.ValueString(), state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error Deleting UniFi WiFi Broadcast", err.Error())
		return
	}
}

func (r *WifiBroadcastResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, "/")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		resp.Diagnostics.AddError("Invalid Import Identifier", fmt.Sprintf("Expected 'site_id/wlan_id', got %q", req.ID))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("site_id"), parts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), parts[1])...)
}
