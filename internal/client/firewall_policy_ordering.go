package client

import (
	"context"
)

type FirewallPolicyOrderingDto struct {
	BeforeSystemDefined []string `json:"beforeSystemDefined,omitempty"`
	AfterSystemDefined  []string `json:"afterSystemDefined,omitempty"`
}

func (c *Client) GetFirewallPolicyOrdering(ctx context.Context, siteID, fromZoneID, toZoneID string) (*FirewallPolicyOrderingDto, error) {
	path := "/v1/sites/" + siteID + "/firewall/zones/" + fromZoneID + "/policy-ordering/" + toZoneID
	var resp FirewallPolicyOrderingDto
	if err := c.DoRequest(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) UpdateFirewallPolicyOrdering(ctx context.Context, siteID, fromZoneID, toZoneID string, req *FirewallPolicyOrderingDto) (*FirewallPolicyOrderingDto, error) {
	path := "/v1/sites/" + siteID + "/firewall/zones/" + fromZoneID + "/policy-ordering/" + toZoneID
	var resp FirewallPolicyOrderingDto
	if err := c.DoRequest(ctx, "PUT", path, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
