package client

import (
	"context"
	"errors"
)

func (c *Client) ListDevices(ctx context.Context, siteID string) ([]struct{ ID string }, error) {
	return nil, errors.New("connection refused")
}

func (c *Client) ListDeviceStatistics(ctx context.Context, siteID string) ([]struct{ ID string }, error) {
	return nil, errors.New("connection refused")
}

func (c *Client) ListClients(ctx context.Context, siteID string) ([]struct{ ID string }, error) {
	return nil, errors.New("connection refused")
}

func (c *Client) ListCountries(ctx context.Context, siteID string) ([]struct{ ID string }, error) {
	return nil, errors.New("connection refused")
}

func (c *Client) ListDpiApplications(ctx context.Context, siteID string) ([]struct{ ID string }, error) {
	return nil, errors.New("connection refused")
}

func (c *Client) ListDpiCategories(ctx context.Context, siteID string) ([]struct{ ID string }, error) {
	return nil, errors.New("connection refused")
}

func (c *Client) ListDeviceTags(ctx context.Context, siteID string) ([]struct{ ID string }, error) {
	return nil, errors.New("connection refused")
}

func (c *Client) ListRadiusProfiles(ctx context.Context, siteID string) ([]struct{ ID string }, error) {
	return nil, errors.New("connection refused")
}

func (c *Client) ListVpnServers(ctx context.Context, siteID string) ([]struct{ ID string }, error) {
	return nil, errors.New("connection refused")
}

func (c *Client) ListVpnTunnels(ctx context.Context, siteID string) ([]struct{ ID string }, error) {
	return nil, errors.New("connection refused")
}

func (c *Client) ListLags(ctx context.Context, siteID string) ([]struct{ ID string }, error) {
	return nil, errors.New("connection refused")
}

func (c *Client) ListMcLagDomains(ctx context.Context, siteID string) ([]struct{ ID string }, error) {
	return nil, errors.New("connection refused")
}

func (c *Client) ListSwitchStacks(ctx context.Context, siteID string) ([]struct{ ID string }, error) {
	return nil, errors.New("connection refused")
}
