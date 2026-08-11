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
	_ resource.Resource                = &ProtectViewerResource{}
	_ resource.ResourceWithConfigure   = &ProtectViewerResource{}
	_ resource.ResourceWithImportState = &ProtectViewerResource{}
)

type ProtectViewerResource struct {
	client client.ProtectClient
}

type ProtectViewerResourceModel struct {
	ID         types.String `tfsdk:"id"`
	Name       types.String `tfsdk:"name"`
	LiveviewID types.String `tfsdk:"liveview_id"`
}

func NewProtectViewerResource() resource.Resource {
	return &ProtectViewerResource{}
}

func (r *ProtectViewerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_protect_viewer"
}

func (r *ProtectViewerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages a UniFi Protect Viewer (Viewport).",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required:    true,
				Description: "Viewer UUID.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the viewer.",
			},
			"liveview_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Liveview ID to display.",
			},
		},
	}
}

func (r *ProtectViewerResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Resource Configure Type", "Expected *client.Client")
		return
	}
	r.client = c.Protect
}

func (r *ProtectViewerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ProtectViewerResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	dto := r.modelToDto(plan)

	res, err := r.client.UpdateViewer(ctx, plan.ID.ValueString(), dto)
	if err != nil {
		resp.Diagnostics.AddError("Error Configuring UniFi Protect Viewer", err.Error())
		return
	}

	state := r.dtoToModel(plan.ID.ValueString(), res)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *ProtectViewerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ProtectViewerResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := r.client.GetViewer(ctx, state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error Reading UniFi Protect Viewer", err.Error())
		return
	}

	newState := r.dtoToModel(state.ID.ValueString(), res)
	resp.Diagnostics.Append(resp.State.Set(ctx, &newState)...)
}

func (r *ProtectViewerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ProtectViewerResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	dto := r.modelToDto(plan)

	res, err := r.client.UpdateViewer(ctx, plan.ID.ValueString(), dto)
	if err != nil {
		resp.Diagnostics.AddError("Error Updating UniFi Protect Viewer", err.Error())
		return
	}

	state := r.dtoToModel(plan.ID.ValueString(), res)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *ProtectViewerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Only remove from state
}

func (r *ProtectViewerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ProtectViewerResource) dtoToModel(id string, dto *client.ViewerDto) ProtectViewerResourceModel {
	return ProtectViewerResourceModel{
		ID:         types.StringValue(id),
		Name:       types.StringValue(dto.Name),
		LiveviewID: types.StringValue(dto.LiveviewID),
	}
}

func (r *ProtectViewerResource) modelToDto(model ProtectViewerResourceModel) *client.ViewerDto {
	dto := &client.ViewerDto{}
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		dto.Name = model.Name.ValueString()
	}
	if !model.LiveviewID.IsNull() && !model.LiveviewID.IsUnknown() {
		dto.LiveviewID = model.LiveviewID.ValueString()
	}
	return dto
}
