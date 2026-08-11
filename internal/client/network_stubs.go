package client

import (
	"context"
	"fmt"
)

func (c *Client) ListDevices(ctx context.Context, siteID string) ([]struct{ ID string }, error) {
	return nil, fmt.Errorf("data source not yet implemented: devices")
}

func (c *Client) ListDeviceStatistics(ctx context.Context, siteID string) ([]struct{ ID string }, error) {
	return nil, fmt.Errorf("data source not yet implemented: device_statistics")
}

func (c *Client) ListClients(ctx context.Context, siteID string) ([]struct{ ID string }, error) {
	return nil, fmt.Errorf("data source not yet implemented: clients")
}

func (c *Client) ListCountries(ctx context.Context, siteID string) ([]struct{ ID string }, error) {
	return nil, fmt.Errorf("data source not yet implemented: countries")
}

func (c *Client) ListDpiApplications(ctx context.Context, siteID string) ([]struct{ ID string }, error) {
	return nil, fmt.Errorf("data source not yet implemented: dpi_applications")
}

func (c *Client) ListDpiCategories(ctx context.Context, siteID string) ([]struct{ ID string }, error) {
	return nil, fmt.Errorf("data source not yet implemented: dpi_categories")
}

func (c *Client) ListDeviceTags(ctx context.Context, siteID string) ([]struct{ ID string }, error) {
	return nil, fmt.Errorf("data source not yet implemented: device_tags")
}

func (c *Client) ListRadiusProfiles(ctx context.Context, siteID string) ([]struct{ ID string }, error) {
	return nil, fmt.Errorf("data source not yet implemented: radius_profiles")
}

func (c *Client) ListVpnServers(ctx context.Context, siteID string) ([]struct{ ID string }, error) {
	return nil, fmt.Errorf("data source not yet implemented: vpn_servers")
}

func (c *Client) ListVpnTunnels(ctx context.Context, siteID string) ([]struct{ ID string }, error) {
	return nil, fmt.Errorf("data source not yet implemented: vpn_tunnels")
}

func (c *Client) ListLags(ctx context.Context, siteID string) ([]struct{ ID string }, error) {
	return nil, fmt.Errorf("data source not yet implemented: lags")
}

func (c *Client) ListMcLagDomains(ctx context.Context, siteID string) ([]struct{ ID string }, error) {
	return nil, fmt.Errorf("data source not yet implemented: mc_lag_domains")
}

func (c *Client) ListSwitchStacks(ctx context.Context, siteID string) ([]struct{ ID string }, error) {
	return nil, fmt.Errorf("data source not yet implemented: switch_stacks")
}
