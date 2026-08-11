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
	_ resource.Resource                = &ProtectChimeResource{}
	_ resource.ResourceWithConfigure   = &ProtectChimeResource{}
	_ resource.ResourceWithImportState = &ProtectChimeResource{}
)

type ProtectChimeResource struct {
	client client.ProtectClient
}

type ProtectChimeResourceModel struct {
	ID     types.String `tfsdk:"id"`
	Name   types.String `tfsdk:"name"`
	Volume types.Int64  `tfsdk:"volume"`
}

func NewProtectChimeResource() resource.Resource {
	return &ProtectChimeResource{}
}

func (r *ProtectChimeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_protect_chime"
}

func (r *ProtectChimeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages a UniFi Protect Chime.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required:    true,
				Description: "Chime UUID.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the chime.",
			},
			"volume": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Volume level.",
			},
		},
	}
}

func (r *ProtectChimeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ProtectChimeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ProtectChimeResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	dto := r.modelToDto(plan)

	res, err := r.client.UpdateChime(ctx, plan.ID.ValueString(), dto)
	if err != nil {
		resp.Diagnostics.AddError("Error Configuring UniFi Protect Chime", err.Error())
		return
	}

	state := r.dtoToModel(plan.ID.ValueString(), res)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *ProtectChimeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ProtectChimeResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := r.client.GetChime(ctx, state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error Reading UniFi Protect Chime", err.Error())
		return
	}

	newState := r.dtoToModel(state.ID.ValueString(), res)
	resp.Diagnostics.Append(resp.State.Set(ctx, &newState)...)
}

func (r *ProtectChimeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ProtectChimeResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	dto := r.modelToDto(plan)

	res, err := r.client.UpdateChime(ctx, plan.ID.ValueString(), dto)
	if err != nil {
		resp.Diagnostics.AddError("Error Updating UniFi Protect Chime", err.Error())
		return
	}

	state := r.dtoToModel(plan.ID.ValueString(), res)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *ProtectChimeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Only remove from state
}

func (r *ProtectChimeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ProtectChimeResource) dtoToModel(id string, dto *client.ChimeDto) ProtectChimeResourceModel {
	return ProtectChimeResourceModel{
		ID:     types.StringValue(id),
		Name:   types.StringValue(dto.Name),
		Volume: types.Int64Value(int64(dto.Volume)),
	}
}

func (r *ProtectChimeResource) modelToDto(model ProtectChimeResourceModel) *client.ChimeDto {
	dto := &client.ChimeDto{}
	if !model.Name.IsNull() && !model.Name.IsUnknown() {
		dto.Name = model.Name.ValueString()
	}
	if !model.Volume.IsNull() && !model.Volume.IsUnknown() {
		dto.Volume = int(model.Volume.ValueInt64())
	}
	return dto
}
