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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                = &DnsPolicyResource{}
	_ resource.ResourceWithConfigure   = &DnsPolicyResource{}
	_ resource.ResourceWithImportState = &DnsPolicyResource{}
)

type DnsPolicyResource struct {
	client client.NetworkClient
}

type DnsPolicyResourceModel struct {
	ID      types.String `tfsdk:"id"`
	SiteID  types.String `tfsdk:"site_id"`
	Name    types.String `tfsdk:"name"`
	Enabled types.Bool   `tfsdk:"enabled"`
	Type    types.String `tfsdk:"type"`
	Value   types.String `tfsdk:"value"`
	TTL     types.Int64  `tfsdk:"ttl"`
}

func NewDnsPolicyResource() resource.Resource {
	return &DnsPolicyResource{}
}

func (r *DnsPolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dns_policy"
}

func (r *DnsPolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages a UniFi DNS Policy.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "Unique DNS Policy UUID.",
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
				Description: "Name of the DNS policy.",
			},
			"enabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(true),
				Description: "Is the DNS policy enabled?",
			},
			"type": schema.StringAttribute{
				Required:    true,
				Description: "Record type: A_RECORD, AAAA_RECORD, CNAME_RECORD, MX_RECORD, SRV_RECORD, TXT_RECORD, FORWARD_DOMAIN",
				Validators: []validator.String{
					stringvalidator.OneOf("A_RECORD", "AAAA_RECORD", "CNAME_RECORD", "MX_RECORD", "SRV_RECORD", "TXT_RECORD", "FORWARD_DOMAIN"),
				},
			},
			"value": schema.StringAttribute{
				Required:    true,
				Description: "Record value or forward target.",
			},
			"ttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(300),
				Description: "Time to Live in seconds.",
			},
		},
	}
}

func (r *DnsPolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *DnsPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan DnsPolicyResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	dto := &client.DnsPolicyDto{
		Domain:  plan.Name.ValueString(),
		Enabled: plan.Enabled.ValueBool(),
		Type:    plan.Type.ValueString(),
	}
	val := plan.Value.ValueString()
	switch dto.Type {
	case "A_RECORD":
		dto.IPv4Address = val
	case "AAAA_RECORD":
		dto.IPv6Address = val
	case "CNAME_RECORD":
		dto.TargetDomain = val
	case "TXT_RECORD":
		dto.Text = val
	case "FORWARD_DOMAIN":
		dto.IPAddress = val
	}
	if !plan.TTL.IsNull() {
		dto.TTLSeconds = int(plan.TTL.ValueInt64())
	}

	res, err := r.client.CreateDnsPolicy(ctx, plan.SiteID.ValueString(), dto)
	if err != nil {
		resp.Diagnostics.AddError("Error Creating UniFi DNS Policy", err.Error())
		return
	}

	plan.ID = types.StringValue(res.ID)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *DnsPolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state DnsPolicyResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := r.client.GetDnsPolicy(ctx, state.SiteID.ValueString(), state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error Reading UniFi DNS Policy", err.Error())
		return
	}

	state.Name = types.StringValue(res.Domain)
	state.Enabled = types.BoolValue(res.Enabled)
	state.Type = types.StringValue(res.Type)
	var val string
	switch res.Type {
	case "A_RECORD":
		val = res.IPv4Address
	case "AAAA_RECORD":
		val = res.IPv6Address
	case "CNAME_RECORD":
		val = res.TargetDomain
	case "TXT_RECORD":
		val = res.Text
	case "FORWARD_DOMAIN":
		val = res.IPAddress
	}
	state.Value = types.StringValue(val)
	if res.TTLSeconds > 0 {
		state.TTL = types.Int64Value(int64(res.TTLSeconds))
	} else {
		state.TTL = types.Int64Null()
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *DnsPolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan DnsPolicyResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	dto := &client.DnsPolicyDto{
		Domain:  plan.Name.ValueString(),
		Enabled: plan.Enabled.ValueBool(),
		Type:    plan.Type.ValueString(),
	}
	val := plan.Value.ValueString()
	switch dto.Type {
	case "A_RECORD":
		dto.IPv4Address = val
	case "AAAA_RECORD":
		dto.IPv6Address = val
	case "CNAME_RECORD":
		dto.TargetDomain = val
	case "TXT_RECORD":
		dto.Text = val
	case "FORWARD_DOMAIN":
		dto.IPAddress = val
	}
	if !plan.TTL.IsNull() {
		dto.TTLSeconds = int(plan.TTL.ValueInt64())
	}

	res, err := r.client.UpdateDnsPolicy(ctx, plan.SiteID.ValueString(), plan.ID.ValueString(), dto)
	if err != nil {
		resp.Diagnostics.AddError("Error Updating UniFi DNS Policy", err.Error())
		return
	}

	plan.ID = types.StringValue(res.ID)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *DnsPolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state DnsPolicyResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteDnsPolicy(ctx, state.SiteID.ValueString(), state.ID.ValueString())
	if err != nil && !client.IsNotFound(err) {
		resp.Diagnostics.AddError("Error Deleting UniFi DNS Policy", err.Error())
		return
	}
}

func (r *DnsPolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, "/")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		resp.Diagnostics.AddError("Invalid Import Identifier", fmt.Sprintf("Expected 'site_id/policy_id', got %q", req.ID))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("site_id"), parts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), parts[1])...)
}
