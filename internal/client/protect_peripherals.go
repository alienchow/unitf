package client

import (
	"context"
)

type LightDto struct {
	ID                  string         `json:"id,omitempty"`
	Name                string         `json:"name,omitempty"`
	LightDeviceSettings *LightSettings `json:"lightDeviceSettings,omitempty"`
}

type LightSettings struct {
	IsIndicatorEnabled bool `json:"isIndicatorEnabled"`
	LedLevel           int  `json:"ledLevel"`
}

type RelayDto struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type SirenDto struct {
	ID     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Volume int    `json:"volume,omitempty"`
}

type ChimeDto struct {
	ID     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Volume int    `json:"volume,omitempty"`
}

type ViewerDto struct {
	ID         string `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	LiveviewID string `json:"liveview,omitempty"`
}

type LiveviewDto struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func (c *Client) GetLight(ctx context.Context, id string) (*LightDto, error) {
	path := "/proxy/protect/integration/v1/lights/" + id
	var resp LightDto
	if err := c.DoRequest(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) UpdateLight(ctx context.Context, id string, req *LightDto) (*LightDto, error) {
	path := "/proxy/protect/integration/v1/lights/" + id
	var resp LightDto
	if err := c.DoRequest(ctx, "PATCH", path, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetRelay(ctx context.Context, id string) (*RelayDto, error) {
	path := "/proxy/protect/integration/v1/relays/" + id
	var resp RelayDto
	if err := c.DoRequest(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) UpdateRelay(ctx context.Context, id string, req *RelayDto) (*RelayDto, error) {
	path := "/proxy/protect/integration/v1/relays/" + id
	var resp RelayDto
	if err := c.DoRequest(ctx, "PATCH", path, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetSiren(ctx context.Context, id string) (*SirenDto, error) {
	path := "/proxy/protect/integration/v1/sirens/" + id
	var resp SirenDto
	if err := c.DoRequest(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) UpdateSiren(ctx context.Context, id string, req *SirenDto) (*SirenDto, error) {
	path := "/proxy/protect/integration/v1/sirens/" + id
	var resp SirenDto
	if err := c.DoRequest(ctx, "PATCH", path, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetChime(ctx context.Context, id string) (*ChimeDto, error) {
	path := "/proxy/protect/integration/v1/chimes/" + id
	var resp ChimeDto
	if err := c.DoRequest(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) UpdateChime(ctx context.Context, id string, req *ChimeDto) (*ChimeDto, error) {
	path := "/proxy/protect/integration/v1/chimes/" + id
	var resp ChimeDto
	if err := c.DoRequest(ctx, "PATCH", path, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetViewer(ctx context.Context, id string) (*ViewerDto, error) {
	path := "/proxy/protect/integration/v1/viewers/" + id
	var resp ViewerDto
	if err := c.DoRequest(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) UpdateViewer(ctx context.Context, id string, req *ViewerDto) (*ViewerDto, error) {
	path := "/proxy/protect/integration/v1/viewers/" + id
	var resp ViewerDto
	if err := c.DoRequest(ctx, "PATCH", path, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) CreateLiveview(ctx context.Context, req *LiveviewDto) (*LiveviewDto, error) {
	path := "/proxy/protect/integration/v1/liveviews"
	var resp LiveviewDto
	if err := c.DoRequest(ctx, "POST", path, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetLiveview(ctx context.Context, id string) (*LiveviewDto, error) {
	path := "/proxy/protect/integration/v1/liveviews/" + id
	var resp LiveviewDto
	if err := c.DoRequest(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) UpdateLiveview(ctx context.Context, id string, req *LiveviewDto) (*LiveviewDto, error) {
	path := "/proxy/protect/integration/v1/liveviews/" + id
	var resp LiveviewDto
	if err := c.DoRequest(ctx, "PATCH", path, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) DeleteLiveview(ctx context.Context, id string) error {
	path := "/proxy/protect/integration/v1/liveviews/" + id
	return c.DoRequest(ctx, "DELETE", path, nil, nil)
}
