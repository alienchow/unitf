package client

import (
	"context"
)

type DnsPolicyDto struct {
	ID      string `json:"id,omitempty"`
	Name    string `json:"name"`
	Enabled bool   `json:"enabled"`
	Type    string `json:"type"` // A, AAAA, CNAME, MX, SRV, TXT, FORWARD_DOMAIN
	Value   string `json:"value"`
	TTL     int    `json:"ttl,omitempty"`
}

func (c *Client) CreateDnsPolicy(ctx context.Context, siteID string, req *DnsPolicyDto) (*DnsPolicyDto, error) {
	path := "/proxy/network/integration/v1/sites/" + siteID + "/dns-policies"
	var resp DnsPolicyDto
	if err := c.DoRequest(ctx, "POST", path, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetDnsPolicy(ctx context.Context, siteID, policyID string) (*DnsPolicyDto, error) {
	path := "/proxy/network/integration/v1/sites/" + siteID + "/dns-policies/" + policyID
	var resp DnsPolicyDto
	if err := c.DoRequest(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) UpdateDnsPolicy(ctx context.Context, siteID, policyID string, req *DnsPolicyDto) (*DnsPolicyDto, error) {
	path := "/proxy/network/integration/v1/sites/" + siteID + "/dns-policies/" + policyID
	var resp DnsPolicyDto
	if err := c.DoRequest(ctx, "PUT", path, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) DeleteDnsPolicy(ctx context.Context, siteID, policyID string) error {
	path := "/proxy/network/integration/v1/sites/" + siteID + "/dns-policies/" + policyID
	return c.DoRequest(ctx, "DELETE", path, nil, nil)
}
