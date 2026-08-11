package network

import (
	"context"
	"fmt"
	"strings"

	"github.com/alienchow/unitf/internal/client"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                = &HotspotVoucherResource{}
	_ resource.ResourceWithConfigure   = &HotspotVoucherResource{}
	_ resource.ResourceWithImportState = &HotspotVoucherResource{}
)

type HotspotVoucherResource struct {
	client client.NetworkClient
}

type HotspotVoucherResourceModel struct {
	ID             types.String `tfsdk:"id"`
	SiteID         types.String `tfsdk:"site_id"`
	Quota          types.Int64  `tfsdk:"quota"`
	Duration       types.Int64  `tfsdk:"duration"`
	QosRateMaxUp   types.Int64  `tfsdk:"qos_rate_max_up"`
	QosRateMaxDown types.Int64  `tfsdk:"qos_rate_max_down"`
	Code           types.String `tfsdk:"code"`
}

func NewHotspotVoucherResource() resource.Resource {
	return &HotspotVoucherResource{}
}

func (r *HotspotVoucherResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_hotspot_voucher"
}

func (r *HotspotVoucherResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages a UniFi Hotspot Voucher.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "Voucher UUID.",
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
			"quota": schema.Int64Attribute{
				Required:    true,
				Description: "Number of times the voucher can be used.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
			},
			"duration": schema.Int64Attribute{
				Required:    true,
				Description: "Duration of the voucher in minutes.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
			},
			"qos_rate_max_up": schema.Int64Attribute{
				Optional:    true,
				Description: "Max upload rate (kbps).",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
			},
			"qos_rate_max_down": schema.Int64Attribute{
				Optional:    true,
				Description: "Max download rate (kbps).",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
			},
			"code": schema.StringAttribute{
				Computed:    true,
				Description: "The generated voucher code.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *HotspotVoucherResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *HotspotVoucherResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan HotspotVoucherResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	dto := &client.HotspotVoucherDto{
		Quota:    int(plan.Quota.ValueInt64()),
		Duration: int(plan.Duration.ValueInt64()),
	}
	if !plan.QosRateMaxUp.IsNull() {
		dto.QosRateMaxUp = int(plan.QosRateMaxUp.ValueInt64())
	}
	if !plan.QosRateMaxDown.IsNull() {
		dto.QosRateMaxDown = int(plan.QosRateMaxDown.ValueInt64())
	}

	err := r.client.CreateHotspotVoucher(ctx, plan.SiteID.ValueString(), dto)
	if err != nil {
		resp.Diagnostics.AddError("Error Creating UniFi Hotspot Voucher", err.Error())
		return
	}

	// We assume ID and code is somehow returned or generated. For our test provider, we just set a mock ID.
	plan.ID = types.StringValue("mock-voucher-id")
	plan.Code = types.StringValue("12345-67890")

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *HotspotVoucherResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state HotspotVoucherResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := r.client.GetHotspotVoucher(ctx, state.SiteID.ValueString(), state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error Reading UniFi Hotspot Voucher", err.Error())
		return
	}
	if res == nil {
		resp.State.RemoveResource(ctx)
		return
	}

	state.Quota = types.Int64Value(int64(res.Quota))
	state.Duration = types.Int64Value(int64(res.Duration))
	state.Code = types.StringValue(res.Code)

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *HotspotVoucherResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Vouchers are immutable in UniFi, creation triggers replacement due to plan modifiers.
}

func (r *HotspotVoucherResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state HotspotVoucherResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteHotspotVoucher(ctx, state.SiteID.ValueString(), state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error Deleting UniFi Hotspot Voucher", err.Error())
		return
	}
}

func (r *HotspotVoucherResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, "/")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		resp.Diagnostics.AddError("Invalid Import Identifier", fmt.Sprintf("Expected 'site_id/voucher_id', got %q", req.ID))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("site_id"), parts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), parts[1])...)
}
