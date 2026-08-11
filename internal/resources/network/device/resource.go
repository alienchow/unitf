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
	_ resource.Resource                = &DeviceResource{}
	_ resource.ResourceWithConfigure   = &DeviceResource{}
	_ resource.ResourceWithImportState = &DeviceResource{}
)

type DeviceResource struct {
	client client.NetworkClient
}

type DeviceResourceModel struct {
	ID     types.String `tfsdk:"id"`
	SiteID types.String `tfsdk:"site_id"`
	Mac    types.String `tfsdk:"mac"`
	Name   types.String `tfsdk:"name"`
	Adopt  types.Bool   `tfsdk:"adopt"`
}

func NewDeviceResource() resource.Resource {
	return &DeviceResource{}
}

func (r *DeviceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device"
}

func (r *DeviceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages adoption and configuration of a UniFi Device.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "Internal UUID of the device.",
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
			"mac": schema.StringAttribute{
				Required:    true,
				Description: "MAC address of the device.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the device.",
			},
			"adopt": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(true),
				Description: "Whether to adopt the device.",
			},
		},
	}
}

func (r *DeviceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *DeviceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan DeviceResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	siteID := plan.SiteID.ValueString()
	mac := plan.Mac.ValueString()

	if plan.Adopt.ValueBool() {
		err := r.client.AdoptDevice(ctx, siteID, mac)
		if err != nil {
			resp.Diagnostics.AddError("Error Adopting UniFi Device", err.Error())
			return
		}
	}

	res, err := r.client.GetDevice(ctx, siteID, mac)
	if err != nil || res == nil {
		resp.Diagnostics.AddError("Error Fetching Adopting UniFi Device", "Device may not be ready")
		return
	}

	if !plan.Name.IsNull() && !plan.Name.IsUnknown() {
		dto := r.modelToDto(plan)
		updatedRes, err := r.client.UpdateDevice(ctx, siteID, res.ID, dto)
		if err != nil {
			resp.Diagnostics.AddError("Error Updating UniFi Device Name", err.Error())
			return
		}
		if updatedRes != nil {
			res = updatedRes
		}
	}

	state := r.dtoToModel(siteID, mac, res)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *DeviceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state DeviceResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := r.client.GetDevice(ctx, state.SiteID.ValueString(), state.Mac.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error Reading UniFi Device", err.Error())
		return
	}
	if res == nil {
		resp.State.RemoveResource(ctx)
		return
	}

	newState := r.dtoToModel(state.SiteID.ValueString(), state.Mac.ValueString(), res)
	resp.Diagnostics.Append(resp.State.Set(ctx, &newState)...)
}

func (r *DeviceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan DeviceResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	siteID := plan.SiteID.ValueString()
	mac := plan.Mac.ValueString()

	dto := r.modelToDto(plan)
	res, err := r.client.UpdateDevice(ctx, siteID, plan.ID.ValueString(), dto)
	if err != nil {
		resp.Diagnostics.AddError("Error Updating UniFi Device", err.Error())
		return
	}

	state := r.dtoToModel(siteID, mac, res)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *DeviceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state DeviceResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.ForgetDevice(ctx, state.SiteID.ValueString(), state.Mac.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error Forgetting UniFi Device", err.Error())
		return
	}
}

func (r *DeviceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, "/")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		resp.Diagnostics.AddError("Invalid Import Identifier", fmt.Sprintf("Expected 'site_id/mac', got %q", req.ID))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("site_id"), parts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("mac"), parts[1])...)
}

func (r *DeviceResource) dtoToModel(siteID string, mac string, dto *client.DeviceDto) DeviceResourceModel {
	state := DeviceResourceModel{
		ID:     types.StringValue(dto.ID),
		SiteID: types.StringValue(siteID),
		Mac:    types.StringValue(mac),
		Name:   types.StringValue(dto.Name),
		Adopt:  types.BoolValue(dto.Adopted),
	}
	return state
}

func (r *DeviceResource) modelToDto(plan DeviceResourceModel) *client.DeviceDto {
	dto := &client.DeviceDto{}
	if !plan.Name.IsNull() && !plan.Name.IsUnknown() {
		dto.Name = plan.Name.ValueString()
	}
	return dto
}
