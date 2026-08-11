package client

import (
	"context"
	"fmt"
)

func (c *Client) ListDevices(ctx context.Context, siteID string) ([]struct{ ID string }, error) {
	var res []struct{ ID string }
	err := c.DoRequest(ctx, "GET", fmt.Sprintf("/proxy/network/integration/v1/sites/%s/devices", siteID), nil, &res)
	return res, err
}

func (c *Client) ListDeviceStatistics(ctx context.Context, siteID string) ([]struct{ ID string }, error) {
	var res []struct{ ID string }
	err := c.DoRequest(ctx, "GET", fmt.Sprintf("/proxy/network/integration/v1/sites/%s/devices/statistics/latest", siteID), nil, &res)
	return res, err
}

func (c *Client) ListClients(ctx context.Context, siteID string) ([]struct{ ID string }, error) {
	var res []struct{ ID string }
	err := c.DoRequest(ctx, "GET", fmt.Sprintf("/proxy/network/integration/v1/sites/%s/clients", siteID), nil, &res)
	return res, err
}

func (c *Client) ListCountries(ctx context.Context, siteID string) ([]struct{ ID string }, error) {
	var res []struct{ ID string }
	err := c.DoRequest(ctx, "GET", "/proxy/network/integration/v1/countries", nil, &res)
	return res, err
}

func (c *Client) ListDpiApplications(ctx context.Context, siteID string) ([]struct{ ID string }, error) {
	var res []struct{ ID string }
	err := c.DoRequest(ctx, "GET", "/proxy/network/integration/v1/dpi/applications", nil, &res)
	return res, err
}

func (c *Client) ListDpiCategories(ctx context.Context, siteID string) ([]struct{ ID string }, error) {
	var res []struct{ ID string }
	err := c.DoRequest(ctx, "GET", "/proxy/network/integration/v1/dpi/categories", nil, &res)
	return res, err
}

func (c *Client) ListDeviceTags(ctx context.Context, siteID string) ([]struct{ ID string }, error) {
	var res []struct{ ID string }
	err := c.DoRequest(ctx, "GET", fmt.Sprintf("/proxy/network/integration/v1/sites/%s/device-tags", siteID), nil, &res)
	return res, err
}

func (c *Client) ListRadiusProfiles(ctx context.Context, siteID string) ([]struct{ ID string }, error) {
	var res []struct{ ID string }
	err := c.DoRequest(ctx, "GET", fmt.Sprintf("/proxy/network/integration/v1/sites/%s/radius/profiles", siteID), nil, &res)
	return res, err
}

func (c *Client) ListVpnServers(ctx context.Context, siteID string) ([]struct{ ID string }, error) {
	var res []struct{ ID string }
	err := c.DoRequest(ctx, "GET", fmt.Sprintf("/proxy/network/integration/v1/sites/%s/vpn/servers", siteID), nil, &res)
	return res, err
}

func (c *Client) ListVpnTunnels(ctx context.Context, siteID string) ([]struct{ ID string }, error) {
	var res []struct{ ID string }
	err := c.DoRequest(ctx, "GET", fmt.Sprintf("/proxy/network/integration/v1/sites/%s/vpn/site-to-site-tunnels", siteID), nil, &res)
	return res, err
}

func (c *Client) ListLags(ctx context.Context, siteID string) ([]struct{ ID string }, error) {
	var res []struct{ ID string }
	err := c.DoRequest(ctx, "GET", fmt.Sprintf("/proxy/network/integration/v1/sites/%s/switching/lags", siteID), nil, &res)
	return res, err
}

func (c *Client) ListMcLagDomains(ctx context.Context, siteID string) ([]struct{ ID string }, error) {
	var res []struct{ ID string }
	err := c.DoRequest(ctx, "GET", fmt.Sprintf("/proxy/network/integration/v1/sites/%s/switching/mc-lag-domains", siteID), nil, &res)
	return res, err
}

func (c *Client) ListSwitchStacks(ctx context.Context, siteID string) ([]struct{ ID string }, error) {
	var res []struct{ ID string }
	err := c.DoRequest(ctx, "GET", fmt.Sprintf("/proxy/network/integration/v1/sites/%s/switching/switch-stacks", siteID), nil, &res)
	return res, err
}
