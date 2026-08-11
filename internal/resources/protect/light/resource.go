package protect

import (
	"context"

	"github.com/alienchow/unitf/internal/client"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                = &ProtectLightResource{}
	_ resource.ResourceWithConfigure   = &ProtectLightResource{}
	_ resource.ResourceWithImportState = &ProtectLightResource{}
)

type ProtectLightResource struct {
	client client.ProtectClient
}

type ProtectLightResourceModel struct {
	ID                 types.String `tfsdk:"id"`
	Name               types.String `tfsdk:"name"`
	IsIndicatorEnabled types.Bool   `tfsdk:"is_indicator_enabled"`
	LedLevel           types.Int64  `tfsdk:"led_level"`
}

func NewProtectLightResource() resource.Resource {
	return &ProtectLightResource{}
}

func (r *ProtectLightResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_protect_light"
}

func (r *ProtectLightResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages a UniFi Protect Light.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required:    true,
				Description: "Light UUID.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the light.",
			},
			"is_indicator_enabled": schema.BoolAttribute{
				Optional:    true,
				Description: "Enable status indicator.",
			},
			"led_level": schema.Int64Attribute{
				Optional:    true,
				Description: "Brightness level.",
			},
		},
	}
}

func (r *ProtectLightResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(client.ProtectClient)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Resource Configure Type", "Expected client.ProtectClient")
		return
	}
	r.client = c
}

func (r *ProtectLightResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ProtectLightResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	dto := &client.LightDto{
		Name:                plan.Name.ValueString(),
		LightDeviceSettings: &client.LightSettings{},
	}
	if !plan.IsIndicatorEnabled.IsNull() {
		dto.LightDeviceSettings.IsIndicatorEnabled = plan.IsIndicatorEnabled.ValueBool()
	}
	if !plan.LedLevel.IsNull() {
		dto.LightDeviceSettings.LedLevel = int(plan.LedLevel.ValueInt64())
	}

	_, err := r.client.UpdateLight(ctx, plan.ID.ValueString(), dto)
	if err != nil {
		resp.Diagnostics.AddError("Error Configuring UniFi Protect Light", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ProtectLightResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ProtectLightResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := r.client.GetLight(ctx, state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error Reading UniFi Protect Light", err.Error())
		return
	}

	state.Name = types.StringValue(res.Name)
	if res.LightDeviceSettings != nil {
		state.IsIndicatorEnabled = types.BoolValue(res.LightDeviceSettings.IsIndicatorEnabled)
		state.LedLevel = types.Int64Value(int64(res.LightDeviceSettings.LedLevel))
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *ProtectLightResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ProtectLightResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	dto := &client.LightDto{
		Name:                plan.Name.ValueString(),
		LightDeviceSettings: &client.LightSettings{},
	}
	if !plan.IsIndicatorEnabled.IsNull() {
		dto.LightDeviceSettings.IsIndicatorEnabled = plan.IsIndicatorEnabled.ValueBool()
	}
	if !plan.LedLevel.IsNull() {
		dto.LightDeviceSettings.LedLevel = int(plan.LedLevel.ValueInt64())
	}

	_, err := r.client.UpdateLight(ctx, plan.ID.ValueString(), dto)
	if err != nil {
		resp.Diagnostics.AddError("Error Updating UniFi Protect Light", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ProtectLightResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Only remove from state
}

func (r *ProtectLightResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
