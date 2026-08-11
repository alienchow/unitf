package client

import (
	"context"
)

func (c *Client) ListCameras(ctx context.Context) ([]CameraDto, error) {
	path := "/v1/cameras"
	var resp []CameraDto
	if err := c.Protect.Request(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) ListNvr(ctx context.Context) ([]struct{ ID string }, error) {
	path := "/v1/nvrs"
	var resp []struct {
		ID string `json:"id"`
	}
	if err := c.Protect.Request(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	var out []struct{ ID string }
	for _, v := range resp {
		out = append(out, struct{ ID string }{ID: v.ID})
	}
	return out, nil
}

func (c *Client) ListSensors(ctx context.Context) ([]struct{ ID string }, error) {
	path := "/v1/sensors"
	var resp []struct {
		ID string `json:"id"`
	}
	if err := c.Protect.Request(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	var out []struct{ ID string }
	for _, v := range resp {
		out = append(out, struct{ ID string }{ID: v.ID})
	}
	return out, nil
}

func (c *Client) ListLights(ctx context.Context) ([]struct{ ID string }, error) {
	path := "/v1/lights"
	var resp []struct {
		ID string `json:"id"`
	}
	if err := c.Protect.Request(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	var out []struct{ ID string }
	for _, v := range resp {
		out = append(out, struct{ ID string }{ID: v.ID})
	}
	return out, nil
}

func (c *Client) ListSpeakers(ctx context.Context) ([]struct{ ID string }, error) {
	path := "/v1/speakers"
	var resp []struct {
		ID string `json:"id"`
	}
	if err := c.Protect.Request(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	var out []struct{ ID string }
	for _, v := range resp {
		out = append(out, struct{ ID string }{ID: v.ID})
	}
	return out, nil
}

func (c *Client) ListLiveviews(ctx context.Context) ([]struct{ ID string }, error) {
	path := "/v1/liveviews"
	var resp []struct {
		ID string `json:"id"`
	}
	if err := c.Protect.Request(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	var out []struct{ ID string }
	for _, v := range resp {
		out = append(out, struct{ ID string }{ID: v.ID})
	}
	return out, nil
}

func (c *Client) ListRelays(ctx context.Context) ([]struct{ ID string }, error) {
	path := "/v1/relays"
	var resp []struct {
		ID string `json:"id"`
	}
	if err := c.Protect.Request(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	var out []struct{ ID string }
	for _, v := range resp {
		out = append(out, struct{ ID string }{ID: v.ID})
	}
	return out, nil
}

func (c *Client) ListSirens(ctx context.Context) ([]struct{ ID string }, error) {
	path := "/v1/sirens"
	var resp []struct {
		ID string `json:"id"`
	}
	if err := c.Protect.Request(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	var out []struct{ ID string }
	for _, v := range resp {
		out = append(out, struct{ ID string }{ID: v.ID})
	}
	return out, nil
}

func (c *Client) ListChimes(ctx context.Context) ([]struct{ ID string }, error) {
	path := "/v1/chimes"
	var resp []struct {
		ID string `json:"id"`
	}
	if err := c.Protect.Request(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	var out []struct{ ID string }
	for _, v := range resp {
		out = append(out, struct{ ID string }{ID: v.ID})
	}
	return out, nil
}

func (c *Client) ListViewers(ctx context.Context) ([]struct{ ID string }, error) {
	path := "/v1/viewers"
	var resp []struct {
		ID string `json:"id"`
	}
	if err := c.Protect.Request(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	var out []struct{ ID string }
	for _, v := range resp {
		out = append(out, struct{ ID string }{ID: v.ID})
	}
	return out, nil
}

func (c *Client) ListFobs(ctx context.Context) ([]struct{ ID string }, error) {
	path := "/v1/fobs"
	var resp []struct {
		ID string `json:"id"`
	}
	if err := c.Protect.Request(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	var out []struct{ ID string }
	for _, v := range resp {
		out = append(out, struct{ ID string }{ID: v.ID})
	}
	return out, nil
}

func (c *Client) ListUsers(ctx context.Context) ([]struct{ ID string }, error) {
	path := "/v1/users"
	var resp []struct {
		ID string `json:"id"`
	}
	if err := c.Protect.Request(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	var out []struct{ ID string }
	for _, v := range resp {
		out = append(out, struct{ ID string }{ID: v.ID})
	}
	return out, nil
}
