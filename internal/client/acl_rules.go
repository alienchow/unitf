package client

import (
	"context"
)

type AclRuleDto struct {
	ID          string               `json:"id,omitempty"`
	Name        string               `json:"name"`
	Enabled     bool                 `json:"enabled"`
	Action      string               `json:"action"`    // ACCEPT, DROP, REJECT
	IPVersion   string               `json:"ipVersion"` // IPV4, IPV6, IPV4_AND_IPV6, MAC
	Protocols   []string             `json:"protocols,omitempty"`
	Source      *FirewallEndpointDto `json:"source,omitempty"`
	Destination *FirewallEndpointDto `json:"destination,omitempty"`
	Logging     bool                 `json:"logging,omitempty"`
}

func (h *NetworkHandler) CreateAclRule(ctx context.Context, siteID string, req *AclRuleDto) (*AclRuleDto, error) {
	path := "/v1/sites/" + siteID + "/firewall/acl-rules"
	var resp AclRuleDto
	if err := h.Request(ctx, "POST", path, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (h *NetworkHandler) GetAclRule(ctx context.Context, siteID, ruleID string) (*AclRuleDto, error) {
	path := "/v1/sites/" + siteID + "/firewall/acl-rules/" + ruleID
	var resp AclRuleDto
	if err := h.Request(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (h *NetworkHandler) UpdateAclRule(ctx context.Context, siteID, ruleID string, req *AclRuleDto) (*AclRuleDto, error) {
	path := "/v1/sites/" + siteID + "/firewall/acl-rules/" + ruleID
	var resp AclRuleDto
	if err := h.Request(ctx, "PUT", path, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (h *NetworkHandler) DeleteAclRule(ctx context.Context, siteID, ruleID string) error {
	path := "/v1/sites/" + siteID + "/firewall/acl-rules/" + ruleID
	return h.Request(ctx, "DELETE", path, nil, nil)
}

func (h *NetworkHandler) ListAclRules(ctx context.Context, siteID string) ([]AclRuleDto, error) {
	path := "/v1/sites/" + siteID + "/firewall/acl-rules"
	var resp struct {
		Data []AclRuleDto `json:"data"`
	}
	if err := h.Request(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
