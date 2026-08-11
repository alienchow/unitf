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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                = &TrafficMatchingListResource{}
	_ resource.ResourceWithConfigure   = &TrafficMatchingListResource{}
	_ resource.ResourceWithImportState = &TrafficMatchingListResource{}
)

type TrafficMatchingListResource struct {
	client client.NetworkClient
}

type TrafficMatchingListResourceModel struct {
	ID        types.String   `tfsdk:"id"`
	SiteID    types.String   `tfsdk:"site_id"`
	Name      types.String   `tfsdk:"name"`
	Type      types.String   `tfsdk:"type"`
	Addresses []types.String `tfsdk:"addresses"`
	Ports     []types.String `tfsdk:"ports"`
}

func NewTrafficMatchingListResource() resource.Resource {
	return &TrafficMatchingListResource{}
}

func (r *TrafficMatchingListResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_traffic_matching_list"
}

func (r *TrafficMatchingListResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages a UniFi Traffic Matching List.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "Unique List UUID.",
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
				Description: "Name of the list.",
			},
			"type": schema.StringAttribute{
				Required:    true,
				Description: "List type: IPV4, IPV6, PORT, IPV4_ADDRESSES, IPV6_ADDRESSES, PORTS.",
				Validators: []validator.String{
					stringvalidator.OneOf("IPV4", "IPV6", "PORT", "IPV4_ADDRESSES", "IPV6_ADDRESSES", "PORTS"),
				},
			},
			"addresses": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "List of IP addresses.",
			},
			"ports": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "List of ports.",
			},
		},
	}
}

func (r *TrafficMatchingListResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *TrafficMatchingListResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan TrafficMatchingListResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	dto := r.modelToDto(plan)

	res, err := r.client.CreateTrafficMatchingList(ctx, plan.SiteID.ValueString(), dto)
	if err != nil {
		resp.Diagnostics.AddError("Error Creating UniFi Traffic Matching List", err.Error())
		return
	}

	state := r.dtoToModel(plan.SiteID.ValueString(), res.ID, res)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *TrafficMatchingListResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state TrafficMatchingListResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := r.client.GetTrafficMatchingList(ctx, state.SiteID.ValueString(), state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Error Reading UniFi Traffic Matching List", err.Error())
		return
	}

	newState := r.dtoToModel(state.SiteID.ValueString(), state.ID.ValueString(), res)
	resp.Diagnostics.Append(resp.State.Set(ctx, &newState)...)
}

func (r *TrafficMatchingListResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan TrafficMatchingListResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	dto := r.modelToDto(plan)

	res, err := r.client.UpdateTrafficMatchingList(ctx, plan.SiteID.ValueString(), plan.ID.ValueString(), dto)
	if err != nil {
		resp.Diagnostics.AddError("Error Updating UniFi Traffic Matching List", err.Error())
		return
	}

	state := r.dtoToModel(plan.SiteID.ValueString(), res.ID, res)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *TrafficMatchingListResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state TrafficMatchingListResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteTrafficMatchingList(ctx, state.SiteID.ValueString(), state.ID.ValueString())
	if err != nil && !client.IsNotFound(err) {
		resp.Diagnostics.AddError("Error Deleting UniFi Traffic Matching List", err.Error())
		return
	}
}

func (r *TrafficMatchingListResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, "/")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		resp.Diagnostics.AddError("Invalid Import Identifier", fmt.Sprintf("Expected 'site_id/list_id', got %q", req.ID))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("site_id"), parts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), parts[1])...)
}

func (r *TrafficMatchingListResource) dtoToModel(siteID string, id string, res *client.TrafficMatchingListDto) TrafficMatchingListResourceModel {
	state := TrafficMatchingListResourceModel{
		ID:     types.StringValue(id),
		SiteID: types.StringValue(siteID),
		Name:   types.StringValue(res.Name),
		Type:   types.StringValue(res.Type),
	}

	state.Addresses = make([]types.String, len(res.Addresses))
	for i, a := range res.Addresses {
		state.Addresses[i] = types.StringValue(a)
	}

	state.Ports = make([]types.String, len(res.Ports))
	for i, p := range res.Ports {
		state.Ports[i] = types.StringValue(p)
	}

	return state
}

func (r *TrafficMatchingListResource) modelToDto(plan TrafficMatchingListResourceModel) *client.TrafficMatchingListDto {
	dto := &client.TrafficMatchingListDto{}

	if !plan.Name.IsNull() && !plan.Name.IsUnknown() {
		dto.Name = plan.Name.ValueString()
	}
	if !plan.Type.IsNull() && !plan.Type.IsUnknown() {
		dto.Type = plan.Type.ValueString()
	}

	for _, a := range plan.Addresses {
		if !a.IsNull() && !a.IsUnknown() {
			dto.Addresses = append(dto.Addresses, a.ValueString())
		}
	}
	for _, p := range plan.Ports {
		if !p.IsNull() && !p.IsUnknown() {
			dto.Ports = append(dto.Ports, p.ValueString())
		}
	}
	return dto
}
