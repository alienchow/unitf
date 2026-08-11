package client

import (
	"context"
)

type TrafficMatchingListDto struct {
	ID        string   `json:"id,omitempty"`
	Name      string   `json:"name"`
	Addresses []string `json:"addresses,omitempty"`
	Ports     []string `json:"ports,omitempty"`
	Type      string   `json:"type"` // IPV4, IPV6, PORT
}

func (c *Client) CreateTrafficMatchingList(ctx context.Context, siteID string, req *TrafficMatchingListDto) (*TrafficMatchingListDto, error) {
	path := "/proxy/network/integration/v1/sites/" + siteID + "/traffic-matching-lists"
	var resp TrafficMatchingListDto
	if err := c.DoRequest(ctx, "POST", path, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetTrafficMatchingList(ctx context.Context, siteID, listID string) (*TrafficMatchingListDto, error) {
	path := "/proxy/network/integration/v1/sites/" + siteID + "/traffic-matching-lists/" + listID
	var resp TrafficMatchingListDto
	if err := c.DoRequest(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) UpdateTrafficMatchingList(ctx context.Context, siteID, listID string, req *TrafficMatchingListDto) (*TrafficMatchingListDto, error) {
	path := "/proxy/network/integration/v1/sites/" + siteID + "/traffic-matching-lists/" + listID
	var resp TrafficMatchingListDto
	if err := c.DoRequest(ctx, "PUT", path, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) DeleteTrafficMatchingList(ctx context.Context, siteID, listID string) error {
	path := "/proxy/network/integration/v1/sites/" + siteID + "/traffic-matching-lists/" + listID
	return c.DoRequest(ctx, "DELETE", path, nil, nil)
}

func (c *Client) ListTrafficMatchingLists(ctx context.Context, siteID string) ([]TrafficMatchingListDto, error) {
	path := "/proxy/network/integration/v1/sites/" + siteID + "/traffic-matching-lists"
	var resp struct {
		Data []TrafficMatchingListDto `json:"data"`
	}
	if err := c.DoRequest(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
