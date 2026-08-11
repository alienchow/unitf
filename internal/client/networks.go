package client

import "context"

type NetworkDto struct {
	ID           string                 `json:"id,omitempty"`
	Name         string                 `json:"name"`
	Type         string                 `json:"type"` // GATEWAY_MANAGED, SWITCH_MANAGED, UNMANAGED
	VlanID       *int                   `json:"vlanId,omitempty"`
	Purpose      string                 `json:"purpose,omitempty"`
	DHCPGuarding *NetworkDHCPGuarding   `json:"dhcpGuarding,omitempty"`
	IPv4         *GatewayIPv4Config     `json:"ipv4,omitempty"`
	IPv6         *NetworkIPv6Config     `json:"ipv6,omitempty"`
	MulticastDNS *bool                  `json:"multicastDns,omitempty"`
	NATOutbound  *WANNATOutboundConfig  `json:"natOutbound,omitempty"`
	Metadata     map[string]interface{} `json:"metadata,omitempty"`
}

type NetworkDHCPGuarding struct {
	Enabled   bool     `json:"enabled"`
	ServerIPs []string `json:"serverIps,omitempty"`
}

type GatewayIPv4Config struct {
	Enabled    bool         `json:"enabled"`
	SubnetMask string       `json:"subnetMask"`
	DHCP       *IPv4DHCPDto `json:"dhcp"`
}

type IPv4DHCPDto struct {
	Mode           string        `json:"mode"` // DHCP_SERVER, DHCP_RELAY, NONE
	DNSServers     []string      `json:"dnsServers,omitempty"`
	GatewayAddress string        `json:"gatewayAddress,omitempty"`
	LeaseTimeSec   *int64        `json:"leaseTimeSec,omitempty"`
	RangeStart     string        `json:"rangeStart,omitempty"`
	RangeStop      string        `json:"rangeStop,omitempty"`
	RelayIP        string        `json:"relayIp,omitempty"`
	PXE            *PXEConfigDto `json:"pxe,omitempty"`
}

type PXEConfigDto struct {
	Enabled  bool   `json:"enabled"`
	ServerIP string `json:"serverIp,omitempty"`
	FileName string `json:"fileName,omitempty"`
}

type NetworkIPv6Config struct {
	Type                 string                  `json:"type"` // PREFIX_DELEGATION, NONE, STATIC
	PrefixDelegationSize *int                    `json:"prefixDelegationSize,omitempty"`
	RouterAdvertisement  *RouterAdvertisementDto `json:"routerAdvertisement,omitempty"`
}

type RouterAdvertisementDto struct {
	Mode                 string   `json:"mode"` // MANAGED, STATELESS
	DNSServers           []string `json:"dnsServers,omitempty"`
	PreferredLifetimeSec *int     `json:"preferredLifetimeSec,omitempty"`
	ValidLifetimeSec     *int     `json:"validLifetimeSec,omitempty"`
}

type WANNATOutboundConfig struct {
	Mode  string `json:"mode"` // AUTO, STATIC
	WanIP string `json:"wanIp,omitempty"`
}

func (c *Client) CreateNetwork(ctx context.Context, siteID string, req *NetworkDto) (*NetworkDto, error) {
	path := "/proxy/network/integration/v1/sites/" + siteID + "/networks"
	var resp NetworkDto
	if err := c.DoRequest(ctx, "POST", path, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetNetwork(ctx context.Context, siteID, networkID string) (*NetworkDto, error) {
	path := "/proxy/network/integration/v1/sites/" + siteID + "/networks/" + networkID
	var resp NetworkDto
	if err := c.DoRequest(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) UpdateNetwork(ctx context.Context, siteID, networkID string, req *NetworkDto) (*NetworkDto, error) {
	path := "/proxy/network/integration/v1/sites/" + siteID + "/networks/" + networkID
	var resp NetworkDto
	if err := c.DoRequest(ctx, "PUT", path, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) DeleteNetwork(ctx context.Context, siteID, networkID string) error {
	path := "/proxy/network/integration/v1/sites/" + siteID + "/networks/" + networkID
	return c.DoRequest(ctx, "DELETE", path, nil, nil)
}
