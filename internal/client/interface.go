package client

import (
	"context"
)

// NetworkClient defines the interface for Network API operations.
type NetworkClient interface {
	CreateNetwork(ctx context.Context, siteID string, req *NetworkDto) (*NetworkDto, error)
	GetNetwork(ctx context.Context, siteID, networkID string) (*NetworkDto, error)
	UpdateNetwork(ctx context.Context, siteID, networkID string, req *NetworkDto) (*NetworkDto, error)
	DeleteNetwork(ctx context.Context, siteID, networkID string) error

	CreateFirewallZone(ctx context.Context, siteID string, req *FirewallZoneDto) (*FirewallZoneDto, error)
	GetFirewallZone(ctx context.Context, siteID, zoneID string) (*FirewallZoneDto, error)
	UpdateFirewallZone(ctx context.Context, siteID, zoneID string, req *FirewallZoneDto) (*FirewallZoneDto, error)
	DeleteFirewallZone(ctx context.Context, siteID, zoneID string) error

	CreateFirewallPolicy(ctx context.Context, siteID string, req *FirewallPolicyDto) (*FirewallPolicyDto, error)
	GetFirewallPolicy(ctx context.Context, siteID, policyID string) (*FirewallPolicyDto, error)
	UpdateFirewallPolicy(ctx context.Context, siteID, policyID string, req *FirewallPolicyDto) (*FirewallPolicyDto, error)
	DeleteFirewallPolicy(ctx context.Context, siteID, policyID string) error

	GetFirewallPolicyOrdering(ctx context.Context, siteID, fromZoneID, toZoneID string) (*FirewallPolicyOrderingDto, error)
	UpdateFirewallPolicyOrdering(ctx context.Context, siteID, fromZoneID, toZoneID string, req *FirewallPolicyOrderingDto) (*FirewallPolicyOrderingDto, error)

	ListSites(ctx context.Context) ([]SiteOverview, error)

	CreateWifiBroadcast(ctx context.Context, siteID string, req *WifiBroadcastDto) (*WifiBroadcastDto, error)
	GetWifiBroadcast(ctx context.Context, siteID, wlanID string) (*WifiBroadcastDto, error)
	UpdateWifiBroadcast(ctx context.Context, siteID, wlanID string, req *WifiBroadcastDto) (*WifiBroadcastDto, error)
	DeleteWifiBroadcast(ctx context.Context, siteID, wlanID string) error

	CreateAclRule(ctx context.Context, siteID string, req *AclRuleDto) (*AclRuleDto, error)
	GetAclRule(ctx context.Context, siteID, ruleID string) (*AclRuleDto, error)
	UpdateAclRule(ctx context.Context, siteID, ruleID string, req *AclRuleDto) (*AclRuleDto, error)
	DeleteAclRule(ctx context.Context, siteID, ruleID string) error

	GetAclRuleOrdering(ctx context.Context, siteID string) (*AclRuleOrderingDto, error)
	UpdateAclRuleOrdering(ctx context.Context, siteID string, req *AclRuleOrderingDto) (*AclRuleOrderingDto, error)

	CreateDnsPolicy(ctx context.Context, siteID string, req *DnsPolicyDto) (*DnsPolicyDto, error)
	GetDnsPolicy(ctx context.Context, siteID, policyID string) (*DnsPolicyDto, error)
	UpdateDnsPolicy(ctx context.Context, siteID, policyID string, req *DnsPolicyDto) (*DnsPolicyDto, error)
	DeleteDnsPolicy(ctx context.Context, siteID, policyID string) error

	CreateTrafficMatchingList(ctx context.Context, siteID string, req *TrafficMatchingListDto) (*TrafficMatchingListDto, error)
	GetTrafficMatchingList(ctx context.Context, siteID, listID string) (*TrafficMatchingListDto, error)
	UpdateTrafficMatchingList(ctx context.Context, siteID, listID string, req *TrafficMatchingListDto) (*TrafficMatchingListDto, error)
	DeleteTrafficMatchingList(ctx context.Context, siteID, listID string) error

	AdoptDevice(ctx context.Context, siteID string, mac string) error
	ForgetDevice(ctx context.Context, siteID string, mac string) error
	GetDevice(ctx context.Context, siteID string, mac string) (*DeviceDto, error)
	UpdateDevice(ctx context.Context, siteID string, deviceID string, req *DeviceDto) (*DeviceDto, error)

	CreateHotspotVoucher(ctx context.Context, siteID string, req *HotspotVoucherDto) error
	GetHotspotVoucher(ctx context.Context, siteID string, id string) (*HotspotVoucherDto, error)
	DeleteHotspotVoucher(ctx context.Context, siteID string, id string) error

	ListWans(ctx context.Context, siteID string) ([]WanOverview, error)
	ListDevices(ctx context.Context, siteID string) ([]struct{ ID string }, error)
	ListDeviceStatistics(ctx context.Context, siteID string) ([]struct{ ID string }, error)
	ListClients(ctx context.Context, siteID string) ([]struct{ ID string }, error)
	ListCountries(ctx context.Context, siteID string) ([]struct{ ID string }, error)
	ListDpiApplications(ctx context.Context, siteID string) ([]struct{ ID string }, error)
	ListDpiCategories(ctx context.Context, siteID string) ([]struct{ ID string }, error)
	ListDeviceTags(ctx context.Context, siteID string) ([]struct{ ID string }, error)
	ListRadiusProfiles(ctx context.Context, siteID string) ([]struct{ ID string }, error)
	ListVpnServers(ctx context.Context, siteID string) ([]struct{ ID string }, error)
	ListVpnTunnels(ctx context.Context, siteID string) ([]struct{ ID string }, error)
	ListLags(ctx context.Context, siteID string) ([]struct{ ID string }, error)
	ListMcLagDomains(ctx context.Context, siteID string) ([]struct{ ID string }, error)
	ListSwitchStacks(ctx context.Context, siteID string) ([]struct{ ID string }, error)
}

// Client implements NetworkClient
var _ NetworkClient = (*Client)(nil)
