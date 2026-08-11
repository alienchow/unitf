package client

import (
	"context"
)

type FirewallZoneDto struct {
	ID         string   `json:"id,omitempty"`
	Name       string   `json:"name"`
	NetworkIDs []string `json:"networkIds"`
}

func (h *NetworkHandler) CreateFirewallZone(ctx context.Context, siteID string, req *FirewallZoneDto) (*FirewallZoneDto, error) {
	path := "/v1/sites/" + siteID + "/firewall/zones"
	var resp FirewallZoneDto
	if err := h.Request(ctx, "POST", path, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (h *NetworkHandler) GetFirewallZone(ctx context.Context, siteID, zoneID string) (*FirewallZoneDto, error) {
	path := "/v1/sites/" + siteID + "/firewall/zones/" + zoneID
	var resp FirewallZoneDto
	if err := h.Request(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (h *NetworkHandler) UpdateFirewallZone(ctx context.Context, siteID, zoneID string, req *FirewallZoneDto) (*FirewallZoneDto, error) {
	path := "/v1/sites/" + siteID + "/firewall/zones/" + zoneID
	var resp FirewallZoneDto
	if err := h.Request(ctx, "PUT", path, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (h *NetworkHandler) DeleteFirewallZone(ctx context.Context, siteID, zoneID string) error {
	path := "/v1/sites/" + siteID + "/firewall/zones/" + zoneID
	return h.Request(ctx, "DELETE", path, nil, nil)
}

func (h *NetworkHandler) ListFirewallZones(ctx context.Context, siteID string) ([]FirewallZoneDto, error) {
	path := "/v1/sites/" + siteID + "/firewall/zones"
	var resp struct {
		Data []FirewallZoneDto `json:"data"`
	}
	if err := h.Request(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
