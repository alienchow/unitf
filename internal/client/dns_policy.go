package client

import (
	"context"
)

type DnsPolicyDto struct {
	ID           string `json:"id,omitempty"`
	Domain       string `json:"domain"`
	Enabled      bool   `json:"enabled"`
	Type         string `json:"type"` // A_RECORD, AAAA_RECORD, CNAME_RECORD, TXT_RECORD, FORWARD_DOMAIN
	TTLSeconds   int    `json:"ttlSeconds,omitempty"`
	IPv4Address  string `json:"ipv4Address,omitempty"`
	IPv6Address  string `json:"ipv6Address,omitempty"`
	TargetDomain string `json:"targetDomain,omitempty"`
	Text         string `json:"text,omitempty"`
	IPAddress    string `json:"ipAddress,omitempty"`
}

func (h *NetworkHandler) CreateDnsPolicy(ctx context.Context, siteID string, req *DnsPolicyDto) (*DnsPolicyDto, error) {
	path := "/v1/sites/" + siteID + "/dns/policies"
	var resp DnsPolicyDto
	if err := h.Request(ctx, "POST", path, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (h *NetworkHandler) GetDnsPolicy(ctx context.Context, siteID, policyID string) (*DnsPolicyDto, error) {
	path := "/v1/sites/" + siteID + "/dns/policies/" + policyID
	var resp DnsPolicyDto
	if err := h.Request(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (h *NetworkHandler) UpdateDnsPolicy(ctx context.Context, siteID, policyID string, req *DnsPolicyDto) (*DnsPolicyDto, error) {
	path := "/v1/sites/" + siteID + "/dns/policies/" + policyID
	var resp DnsPolicyDto
	if err := h.Request(ctx, "PUT", path, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (h *NetworkHandler) DeleteDnsPolicy(ctx context.Context, siteID, policyID string) error {
	path := "/v1/sites/" + siteID + "/dns/policies/" + policyID
	return h.Request(ctx, "DELETE", path, nil, nil)
}

func (h *NetworkHandler) ListDnsPolicies(ctx context.Context, siteID string) ([]DnsPolicyDto, error) {
	path := "/v1/sites/" + siteID + "/dns/policies"
	var resp struct {
		Data []DnsPolicyDto `json:"data"`
	}
	if err := h.Request(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
