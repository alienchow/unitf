package client

import (
	"context"
)

type WifiBroadcastDto struct {
	ID         string `json:"id,omitempty"`
	Name       string `json:"name"`
	Enabled    bool   `json:"enabled"`
	SSID       string `json:"ssid"`
	Security   string `json:"security"`
	Passphrase string `json:"passphrase,omitempty"`
	NetworkID  string `json:"networkId"`
	Mode       string `json:"mode,omitempty"` // Standard, IoT, etc.
}

func (c *Client) CreateWifiBroadcast(ctx context.Context, siteID string, req *WifiBroadcastDto) (*WifiBroadcastDto, error) {
	path := "/v1/sites/" + siteID + "/wifi-broadcasts"
	var resp WifiBroadcastDto
	if err := c.Network.Request(ctx, "POST", path, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetWifiBroadcast(ctx context.Context, siteID, wlanID string) (*WifiBroadcastDto, error) {
	path := "/v1/sites/" + siteID + "/wifi-broadcasts/" + wlanID
	var resp WifiBroadcastDto
	if err := c.Network.Request(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) UpdateWifiBroadcast(ctx context.Context, siteID, wlanID string, req *WifiBroadcastDto) (*WifiBroadcastDto, error) {
	path := "/v1/sites/" + siteID + "/wifi-broadcasts/" + wlanID
	var resp WifiBroadcastDto
	if err := c.Network.Request(ctx, "PUT", path, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) DeleteWifiBroadcast(ctx context.Context, siteID, wlanID string) error {
	path := "/v1/sites/" + siteID + "/wifi-broadcasts/" + wlanID
	return c.Network.Request(ctx, "DELETE", path, nil, nil)
}

func (c *Client) ListWifiBroadcasts(ctx context.Context, siteID string) ([]WifiBroadcastDto, error) {
	path := "/v1/sites/" + siteID + "/wifi-broadcasts"
	var resp struct {
		Data []WifiBroadcastDto `json:"data"`
	}
	if err := c.Network.Request(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
