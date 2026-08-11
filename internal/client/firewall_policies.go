package client

import (
	"context"
)

type FirewallPolicyDto struct {
	ID              string               `json:"id,omitempty"`
	Name            string               `json:"name"`
	Enabled         bool                 `json:"enabled"`
	Action          *FirewallActionDto   `json:"action"`
	IPProtocolScope *IPProtocolScopeDto  `json:"ipProtocolScope,omitempty"`
	Source          *FirewallEndpointDto `json:"source,omitempty"`
	Destination     *FirewallEndpointDto `json:"destination,omitempty"`
	Logging         bool                 `json:"logging,omitempty"`
	Schedule        *ScheduleDto         `json:"schedule,omitempty"`
}

type FirewallActionDto struct {
	Accept *struct{} `json:"accept,omitempty"`
	Block  *struct{} `json:"block,omitempty"`
	Drop   *struct{} `json:"drop,omitempty"`
	Reject *struct{} `json:"reject,omitempty"`
}

type IPProtocolScopeDto struct {
	IPVersion string   `json:"ipVersion"`
	Protocols []string `json:"protocols,omitempty"`
}

type FirewallEndpointDto struct {
	ZoneID      string `json:"zoneId,omitempty"`
	NetworkID   string `json:"networkId,omitempty"`
	Address     string `json:"address,omitempty"`
	Port        string `json:"port,omitempty"`
	MACAddress  string `json:"macAddress,omitempty"`
	MatchListID string `json:"matchListId,omitempty"`
}

type ScheduleDto struct {
	Mode      string        `json:"mode"`
	TimeRange *TimeRangeDto `json:"timeRange,omitempty"`
}

type TimeRangeDto struct {
	StartTime string   `json:"startTime"`
	EndTime   string   `json:"endTime"`
	Days      []string `json:"days"`
}

func (c *Client) CreateFirewallPolicy(ctx context.Context, siteID string, req *FirewallPolicyDto) (*FirewallPolicyDto, error) {
	path := "/proxy/network/integration/v1/sites/" + siteID + "/firewall/policies"
	var resp FirewallPolicyDto
	if err := c.DoRequest(ctx, "POST", path, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetFirewallPolicy(ctx context.Context, siteID, policyID string) (*FirewallPolicyDto, error) {
	path := "/proxy/network/integration/v1/sites/" + siteID + "/firewall/policies/" + policyID
	var resp FirewallPolicyDto
	if err := c.DoRequest(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) UpdateFirewallPolicy(ctx context.Context, siteID, policyID string, req *FirewallPolicyDto) (*FirewallPolicyDto, error) {
	path := "/proxy/network/integration/v1/sites/" + siteID + "/firewall/policies/" + policyID
	var resp FirewallPolicyDto
	if err := c.DoRequest(ctx, "PUT", path, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) DeleteFirewallPolicy(ctx context.Context, siteID, policyID string) error {
	path := "/proxy/network/integration/v1/sites/" + siteID + "/firewall/policies/" + policyID
	return c.DoRequest(ctx, "DELETE", path, nil, nil)
}

func (c *Client) ListFirewallPolicies(ctx context.Context, siteID string) ([]FirewallPolicyDto, error) {
	path := "/proxy/network/integration/v1/sites/" + siteID + "/firewall/policies"
	var resp struct {
		Data []FirewallPolicyDto `json:"data"`
	}
	if err := c.DoRequest(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
