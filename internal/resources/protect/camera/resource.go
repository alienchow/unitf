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
	_ resource.Resource                = &ProtectCameraResource{}
	_ resource.ResourceWithConfigure   = &ProtectCameraResource{}
	_ resource.ResourceWithImportState = &ProtectCameraResource{}
)

type ProtectCameraResource struct {
	client client.ProtectClient
}

type ProtectCameraResourceModel struct {
	ID               types.String             `tfsdk:"id"`
	Name             types.String             `tfsdk:"name"`
	VideoMode        types.String             `tfsdk:"video_mode"`
	RecordEverything types.Bool               `tfsdk:"record_everything"`
	OSDSettings      *ProtectCameraOSDModel   `tfsdk:"osd_settings"`
	SmartDetect      *ProtectCameraSmartModel `tfsdk:"smart_detect_settings"`
}

type ProtectCameraOSDModel struct {
	IsNameEnabled types.Bool `tfsdk:"is_name_enabled"`
	IsDateEnabled types.Bool `tfsdk:"is_date_enabled"`
}

type ProtectCameraSmartModel struct {
	ObjectTypes []types.String `tfsdk:"object_types"`
}

func NewProtectCameraResource() resource.Resource {
	return &ProtectCameraResource{}
}

func (r *ProtectCameraResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_protect_camera"
}

func (r *ProtectCameraResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages a UniFi Protect Camera.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required:    true,
				Description: "Camera UUID (Cameras are adopted, we just configure them).",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the camera.",
			},
			"video_mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Video mode (e.g. default, highFps).",
			},
			"record_everything": schema.BoolAttribute{
				Optional: true,
				Computed: true,
			},
			"osd_settings": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"is_name_enabled": schema.BoolAttribute{Optional: true, Computed: true},
					"is_date_enabled": schema.BoolAttribute{Optional: true, Computed: true},
				},
			},
			"smart_detect_settings": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"object_types": schema.ListAttribute{
						ElementType: types.StringType,
						Optional:    true,
						Computed:    true,
					},
				},
			},
		},
	}
}

func (r *ProtectCameraResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	// Ensure that our provider also gives us a type that implements ProtectClient
	c, ok := req.ProviderData.(client.ProtectClient)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Resource Configure Type", "Expected client.ProtectClient")
		return
	}
	r.client = c
}

func (r *ProtectCameraResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Cameras are physical devices adopted to the NVR. We don't "Create" them via API POST usually.
	// The resource ID is provided. We just treat Create as an Update to configure it.
	var plan ProtectCameraResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	dto := r.modelToDto(plan)
	_, err := r.client.UpdateCamera(ctx, plan.ID.ValueString(), dto)
	if err != nil {
		resp.Diagnostics.AddError("Error Configuring UniFi Protect Camera", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ProtectCameraResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ProtectCameraResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := r.client.GetCamera(ctx, state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error Reading UniFi Protect Camera", err.Error())
		return
	}

	state.Name = types.StringValue(res.Name)
	state.VideoMode = types.StringValue(res.VideoMode)
	state.RecordEverything = types.BoolValue(res.RecordEverything)

	if res.OSDSettings != nil {
		state.OSDSettings = &ProtectCameraOSDModel{
			IsNameEnabled: types.BoolValue(res.OSDSettings.IsNameEnabled),
			IsDateEnabled: types.BoolValue(res.OSDSettings.IsDateEnabled),
		}
	} else {
		state.OSDSettings = nil
	}

	if res.SmartDetect != nil && len(res.SmartDetect.ObjectTypes) > 0 {
		typesList := make([]types.String, len(res.SmartDetect.ObjectTypes))
		for i, v := range res.SmartDetect.ObjectTypes {
			typesList[i] = types.StringValue(v)
		}
		state.SmartDetect = &ProtectCameraSmartModel{
			ObjectTypes: typesList,
		}
	} else {
		state.SmartDetect = nil
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *ProtectCameraResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ProtectCameraResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	dto := r.modelToDto(plan)
	_, err := r.client.UpdateCamera(ctx, plan.ID.ValueString(), dto)
	if err != nil {
		resp.Diagnostics.AddError("Error Updating UniFi Protect Camera", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ProtectCameraResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Just remove from state. We don't unadopt the physical camera.
}

func (r *ProtectCameraResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ProtectCameraResource) modelToDto(model ProtectCameraResourceModel) *client.CameraDto {
	dto := &client.CameraDto{}

	if !model.Name.IsNull() {
		dto.Name = model.Name.ValueString()
	}
	if !model.VideoMode.IsNull() {
		dto.VideoMode = model.VideoMode.ValueString()
	}
	if !model.RecordEverything.IsNull() {
		dto.RecordEverything = model.RecordEverything.ValueBool()
	}

	if model.OSDSettings != nil {
		dto.OSDSettings = &client.CameraOSDSettings{
			IsNameEnabled: model.OSDSettings.IsNameEnabled.ValueBool(),
			IsDateEnabled: model.OSDSettings.IsDateEnabled.ValueBool(),
		}
	}

	if model.SmartDetect != nil {
		var objs []string
		for _, o := range model.SmartDetect.ObjectTypes {
			objs = append(objs, o.ValueString())
		}
		dto.SmartDetect = &client.CameraSmartDetect{
			ObjectTypes: objs,
		}
	}

	return dto
}
