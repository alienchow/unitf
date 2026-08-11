package client

import (
	"context"
)

type AclRuleOrderingDto struct {
	RuleIDs []string `json:"ruleIds"`
}

func (h *NetworkHandler) GetAclRuleOrdering(ctx context.Context, siteID string) (*AclRuleOrderingDto, error) {
	path := "/v1/sites/" + siteID + "/firewall/acl-rules/ordering"
	var resp AclRuleOrderingDto
	if err := h.Request(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (h *NetworkHandler) UpdateAclRuleOrdering(ctx context.Context, siteID string, req *AclRuleOrderingDto) (*AclRuleOrderingDto, error) {
	path := "/v1/sites/" + siteID + "/firewall/acl-rules/ordering"
	var resp AclRuleOrderingDto
	if err := h.Request(ctx, "PUT", path, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
