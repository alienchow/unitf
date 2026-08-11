package client

import (
	"context"
)

type DeviceDto struct {
	ID         string `json:"_id,omitempty"`
	Mac        string `json:"mac,omitempty"`
	Name       string `json:"name,omitempty"`
	Adopted    bool   `json:"adopted,omitempty"`
	PortConfig []Port `json:"port_overrides,omitempty"`
}

type Port struct {
	PortIdx  int    `json:"port_idx"`
	PortConf string `json:"portconf_id"` // ID of the port profile
	Name     string `json:"name,omitempty"`
}

func (c *Client) AdoptDevice(ctx context.Context, siteID string, mac string) error {
	path := "/api/s/" + siteID + "/cmd/devmgr"
	payload := map[string]interface{}{
		"cmd": "adopt",
		"mac": mac,
	}
	return c.Network.Request(ctx, "POST", path, payload, nil)
}

func (c *Client) ForgetDevice(ctx context.Context, siteID string, mac string) error {
	path := "/api/s/" + siteID + "/cmd/sitemgr"
	payload := map[string]interface{}{
		"cmd": "forget-dev",
		"mac": mac,
	}
	return c.Network.Request(ctx, "POST", path, payload, nil)
}

func (c *Client) GetDevice(ctx context.Context, siteID string, mac string) (*DeviceDto, error) {
	path := "/api/s/" + siteID + "/stat/device/" + mac
	var resp struct {
		Data []DeviceDto `json:"data"`
	}
	if err := c.Network.Request(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	if len(resp.Data) > 0 {
		return &resp.Data[0], nil
	}
	return nil, nil // Or an error "not found"
}

func (c *Client) UpdateDevice(ctx context.Context, siteID string, deviceID string, req *DeviceDto) (*DeviceDto, error) {
	path := "/api/s/" + siteID + "/upd/device/" + deviceID
	var resp struct {
		Data []DeviceDto `json:"data"`
	}
	if err := c.Network.Request(ctx, "PUT", path, req, &resp); err != nil {
		return nil, err
	}
	if len(resp.Data) > 0 {
		return &resp.Data[0], nil
	}
	return nil, nil
}
