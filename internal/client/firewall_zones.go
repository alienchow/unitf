package client

import (
	"context"
)

type FirewallZoneDto struct {
	ID         string   `json:"id,omitempty"`
	Name       string   `json:"name"`
	NetworkIDs []string `json:"networkIds"`
}

func (c *Client) CreateFirewallZone(ctx context.Context, siteID string, req *FirewallZoneDto) (*FirewallZoneDto, error) {
	path := "/v1/sites/" + siteID + "/firewall/zones"
	var resp FirewallZoneDto
	if err := c.DoRequest(ctx, "POST", path, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetFirewallZone(ctx context.Context, siteID, zoneID string) (*FirewallZoneDto, error) {
	path := "/v1/sites/" + siteID + "/firewall/zones/" + zoneID
	var resp FirewallZoneDto
	if err := c.DoRequest(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) UpdateFirewallZone(ctx context.Context, siteID, zoneID string, req *FirewallZoneDto) (*FirewallZoneDto, error) {
	path := "/v1/sites/" + siteID + "/firewall/zones/" + zoneID
	var resp FirewallZoneDto
	if err := c.DoRequest(ctx, "PUT", path, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) DeleteFirewallZone(ctx context.Context, siteID, zoneID string) error {
	path := "/v1/sites/" + siteID + "/firewall/zones/" + zoneID
	return c.DoRequest(ctx, "DELETE", path, nil, nil)
}

func (c *Client) ListFirewallZones(ctx context.Context, siteID string) ([]FirewallZoneDto, error) {
	path := "/v1/sites/" + siteID + "/firewall/zones"
	var resp struct {
		Data []FirewallZoneDto `json:"data"`
	}
	if err := c.DoRequest(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
