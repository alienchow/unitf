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
	_ resource.Resource                = &ProtectSensorResource{}
	_ resource.ResourceWithConfigure   = &ProtectSensorResource{}
	_ resource.ResourceWithImportState = &ProtectSensorResource{}
)

type ProtectSensorResource struct {
	client client.ProtectClient
}

type ProtectSensorResourceModel struct {
	ID        types.String `tfsdk:"id"`
	Name      types.String `tfsdk:"name"`
	Alarm     types.Bool   `tfsdk:"alarm"`
	TempLimit types.Int64  `tfsdk:"temp_limit"`
}

func NewProtectSensorResource() resource.Resource {
	return &ProtectSensorResource{}
}

func (r *ProtectSensorResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_protect_sensor"
}

func (r *ProtectSensorResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages a UniFi Protect Sensor.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required:    true,
				Description: "Sensor UUID (Sensors are adopted, we just configure them).",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the sensor.",
			},
			"alarm": schema.BoolAttribute{
				Optional:    true,
				Description: "Is alarm enabled?",
			},
			"temp_limit": schema.Int64Attribute{
				Optional:    true,
				Description: "Temperature alarm threshold.",
			},
		},
	}
}

func (r *ProtectSensorResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ProtectSensorResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ProtectSensorResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	dto := &client.SensorDto{
		Name: plan.Name.ValueString(),
	}
	if !plan.Alarm.IsNull() {
		dto.Alarm = plan.Alarm.ValueBool()
	}
	if !plan.TempLimit.IsNull() {
		dto.TempLimit = int(plan.TempLimit.ValueInt64())
	}

	_, err := r.client.UpdateSensor(ctx, plan.ID.ValueString(), dto)
	if err != nil {
		resp.Diagnostics.AddError("Error Configuring UniFi Protect Sensor", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ProtectSensorResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ProtectSensorResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := r.client.GetSensor(ctx, state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error Reading UniFi Protect Sensor", err.Error())
		return
	}

	state.Name = types.StringValue(res.Name)
	state.Alarm = types.BoolValue(res.Alarm)
	state.TempLimit = types.Int64Value(int64(res.TempLimit))

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *ProtectSensorResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ProtectSensorResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	dto := &client.SensorDto{
		Name: plan.Name.ValueString(),
	}
	if !plan.Alarm.IsNull() {
		dto.Alarm = plan.Alarm.ValueBool()
	}
	if !plan.TempLimit.IsNull() {
		dto.TempLimit = int(plan.TempLimit.ValueInt64())
	}

	_, err := r.client.UpdateSensor(ctx, plan.ID.ValueString(), dto)
	if err != nil {
		resp.Diagnostics.AddError("Error Updating UniFi Protect Sensor", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ProtectSensorResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Just remove from state.
}

func (r *ProtectSensorResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
