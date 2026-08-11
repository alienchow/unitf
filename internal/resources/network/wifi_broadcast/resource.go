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
				Computed:    true,
				Sensitive:   true,
				Description: "WiFi passphrase. Required for WPA2/WPA3.",
			},
			"network_id": schema.StringAttribute{
				Required:    true,
				Description: "ID of the network to bridge this WiFi to.",
			},
			"mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Mode of the WiFi broadcast (e.g., STANDARD, IOT).",
			},
		},
	}
}

func (r *WifiBroadcastResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *WifiBroadcastResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan WifiBroadcastResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	dto := r.modelToDto(plan)

	res, err := r.client.CreateWifiBroadcast(ctx, plan.SiteID.ValueString(), dto)
	if err != nil {
		resp.Diagnostics.AddError("Error Creating UniFi WiFi Broadcast", err.Error())
		return
	}

	state := r.dtoToModel(plan.SiteID.ValueString(), res.ID, res, plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *WifiBroadcastResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state WifiBroadcastResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := r.client.GetWifiBroadcast(ctx, state.SiteID.ValueString(), state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error Reading UniFi WiFi Broadcast", err.Error())
		return
	}

	newState := r.dtoToModel(state.SiteID.ValueString(), state.ID.ValueString(), res, state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &newState)...)
}

func (r *WifiBroadcastResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan WifiBroadcastResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	dto := r.modelToDto(plan)

	res, err := r.client.UpdateWifiBroadcast(ctx, plan.SiteID.ValueString(), plan.ID.ValueString(), dto)
	if err != nil {
		resp.Diagnostics.AddError("Error Updating UniFi WiFi Broadcast", err.Error())
		return
	}

	state := r.dtoToModel(plan.SiteID.ValueString(), res.ID, res, plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *WifiBroadcastResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state WifiBroadcastResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteWifiBroadcast(ctx, state.SiteID.ValueString(), state.ID.ValueString())
	if err != nil && !client.IsNotFound(err) {
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

func (r *WifiBroadcastResource) dtoToModel(siteID string, id string, res *client.WifiBroadcastDto, plan WifiBroadcastResourceModel) WifiBroadcastResourceModel {
	state := WifiBroadcastResourceModel{
		ID:        types.StringValue(id),
		SiteID:    types.StringValue(siteID),
		Name:      types.StringValue(res.Name),
		Enabled:   types.BoolValue(res.Enabled),
		SSID:      types.StringValue(res.SSID),
		Security:  types.StringValue(res.Security),
		NetworkID: types.StringValue(res.NetworkID),
	}

	if res.Mode != "" {
		state.Mode = types.StringValue(res.Mode)
	} else {
		state.Mode = types.StringNull()
	}

	if res.Passphrase != "" {
		state.Passphrase = types.StringValue(res.Passphrase)
	} else {
		state.Passphrase = plan.Passphrase
	}

	return state
}

func (r *WifiBroadcastResource) modelToDto(plan WifiBroadcastResourceModel) *client.WifiBroadcastDto {
	dto := &client.WifiBroadcastDto{}

	if !plan.Name.IsNull() && !plan.Name.IsUnknown() {
		dto.Name = plan.Name.ValueString()
	}
	if !plan.Enabled.IsNull() && !plan.Enabled.IsUnknown() {
		dto.Enabled = plan.Enabled.ValueBool()
	}
	if !plan.SSID.IsNull() && !plan.SSID.IsUnknown() {
		dto.SSID = plan.SSID.ValueString()
	}
	if !plan.Security.IsNull() && !plan.Security.IsUnknown() {
		dto.Security = plan.Security.ValueString()
	}
	if !plan.NetworkID.IsNull() && !plan.NetworkID.IsUnknown() {
		dto.NetworkID = plan.NetworkID.ValueString()
	}
	if !plan.Passphrase.IsNull() && !plan.Passphrase.IsUnknown() {
		dto.Passphrase = plan.Passphrase.ValueString()
	}
	if !plan.Mode.IsNull() && !plan.Mode.IsUnknown() {
		dto.Mode = plan.Mode.ValueString()
	}
	return dto
}
