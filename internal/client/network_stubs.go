package client

import (
	"context"
	"fmt"
)

func (h *NetworkHandler) listIDs(ctx context.Context, path string) ([]struct{ ID string }, error) {
	var res []struct{ ID string }
	err := h.Request(ctx, "GET", path, nil, &res)
	return res, err
}

func (h *NetworkHandler) ListDevices(ctx context.Context, siteID string) ([]struct{ ID string }, error) {
	return h.listIDs(ctx, fmt.Sprintf("/v1/sites/%s/devices", siteID))
}

func (h *NetworkHandler) ListDeviceStatistics(ctx context.Context, siteID string) ([]struct{ ID string }, error) {
	return h.listIDs(ctx, fmt.Sprintf("/v1/sites/%s/devices/statistics/latest", siteID))
}

func (h *NetworkHandler) ListClients(ctx context.Context, siteID string) ([]struct{ ID string }, error) {
	return h.listIDs(ctx, fmt.Sprintf("/v1/sites/%s/clients", siteID))
}

func (h *NetworkHandler) ListCountries(ctx context.Context, siteID string) ([]struct{ ID string }, error) {
	return h.listIDs(ctx, "/v1/countries")
}

func (h *NetworkHandler) ListDpiApplications(ctx context.Context, siteID string) ([]struct{ ID string }, error) {
	return h.listIDs(ctx, "/v1/dpi/applications")
}

func (h *NetworkHandler) ListDpiCategories(ctx context.Context, siteID string) ([]struct{ ID string }, error) {
	return h.listIDs(ctx, "/v1/dpi/categories")
}

func (h *NetworkHandler) ListDeviceTags(ctx context.Context, siteID string) ([]struct{ ID string }, error) {
	return h.listIDs(ctx, fmt.Sprintf("/v1/sites/%s/device-tags", siteID))
}

func (h *NetworkHandler) ListRadiusProfiles(ctx context.Context, siteID string) ([]struct{ ID string }, error) {
	return h.listIDs(ctx, fmt.Sprintf("/v1/sites/%s/radius/profiles", siteID))
}

func (h *NetworkHandler) ListVpnServers(ctx context.Context, siteID string) ([]struct{ ID string }, error) {
	return h.listIDs(ctx, fmt.Sprintf("/v1/sites/%s/vpn/servers", siteID))
}

func (h *NetworkHandler) ListVpnTunnels(ctx context.Context, siteID string) ([]struct{ ID string }, error) {
	return h.listIDs(ctx, fmt.Sprintf("/v1/sites/%s/vpn/site-to-site-tunnels", siteID))
}

func (h *NetworkHandler) ListLags(ctx context.Context, siteID string) ([]struct{ ID string }, error) {
	return h.listIDs(ctx, fmt.Sprintf("/v1/sites/%s/switching/lags", siteID))
}

func (h *NetworkHandler) ListMcLagDomains(ctx context.Context, siteID string) ([]struct{ ID string }, error) {
	return h.listIDs(ctx, fmt.Sprintf("/v1/sites/%s/switching/mc-lag-domains", siteID))
}

func (h *NetworkHandler) ListSwitchStacks(ctx context.Context, siteID string) ([]struct{ ID string }, error) {
	return h.listIDs(ctx, fmt.Sprintf("/v1/sites/%s/switching/switch-stacks", siteID))
}
