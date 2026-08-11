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
	_ resource.Resource                = &ProtectSirenResource{}
	_ resource.ResourceWithConfigure   = &ProtectSirenResource{}
	_ resource.ResourceWithImportState = &ProtectSirenResource{}
)

type ProtectSirenResource struct {
	client client.ProtectClient
}

type ProtectSirenResourceModel struct {
	ID     types.String `tfsdk:"id"`
	Name   types.String `tfsdk:"name"`
	Volume types.Int64  `tfsdk:"volume"`
}

func NewProtectSirenResource() resource.Resource {
	return &ProtectSirenResource{}
}

func (r *ProtectSirenResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_protect_siren"
}

func (r *ProtectSirenResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages a UniFi Protect Siren.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required:    true,
				Description: "Siren UUID.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the siren.",
			},
			"volume": schema.Int64Attribute{
				Optional:    true,
				Description: "Volume level.",
			},
		},
	}
}

func (r *ProtectSirenResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ProtectSirenResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ProtectSirenResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	dto := &client.SirenDto{
		Name: plan.Name.ValueString(),
	}
	if !plan.Volume.IsNull() {
		dto.Volume = int(plan.Volume.ValueInt64())
	}

	_, err := r.client.UpdateSiren(ctx, plan.ID.ValueString(), dto)
	if err != nil {
		resp.Diagnostics.AddError("Error Configuring UniFi Protect Siren", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ProtectSirenResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ProtectSirenResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := r.client.GetSiren(ctx, state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error Reading UniFi Protect Siren", err.Error())
		return
	}

	state.Name = types.StringValue(res.Name)
	state.Volume = types.Int64Value(int64(res.Volume))

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *ProtectSirenResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ProtectSirenResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	dto := &client.SirenDto{
		Name: plan.Name.ValueString(),
	}
	if !plan.Volume.IsNull() {
		dto.Volume = int(plan.Volume.ValueInt64())
	}

	_, err := r.client.UpdateSiren(ctx, plan.ID.ValueString(), dto)
	if err != nil {
		resp.Diagnostics.AddError("Error Updating UniFi Protect Siren", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ProtectSirenResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Only remove from state
}

func (r *ProtectSirenResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
