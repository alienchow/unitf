package provider

import (
	"context"
	"os"

	"github.com/alienchow/unitf/internal/client"
	clients_ds "github.com/alienchow/unitf/internal/datasources/network/clients"
	countries_ds "github.com/alienchow/unitf/internal/datasources/network/countries"
	device_statistics_ds "github.com/alienchow/unitf/internal/datasources/network/device_statistics"
	device_tags_ds "github.com/alienchow/unitf/internal/datasources/network/device_tags"
	devices_ds "github.com/alienchow/unitf/internal/datasources/network/devices"
	dpi_applications_ds "github.com/alienchow/unitf/internal/datasources/network/dpi_applications"
	dpi_categories_ds "github.com/alienchow/unitf/internal/datasources/network/dpi_categories"
	lags_ds "github.com/alienchow/unitf/internal/datasources/network/lags"
	mc_lag_domains_ds "github.com/alienchow/unitf/internal/datasources/network/mc_lag_domains"
	radius_profiles_ds "github.com/alienchow/unitf/internal/datasources/network/radius_profiles"
	sites_ds "github.com/alienchow/unitf/internal/datasources/network/sites"
	traffic_matching_lists_ds "github.com/alienchow/unitf/internal/datasources/network/traffic_matching_lists"
	dns_policies_ds "github.com/alienchow/unitf/internal/datasources/network/dns_policies"
	acl_rules_ds "github.com/alienchow/unitf/internal/datasources/network/acl_rules"
	wifi_broadcasts_ds "github.com/alienchow/unitf/internal/datasources/network/wifi_broadcasts"
	firewall_policies_ds "github.com/alienchow/unitf/internal/datasources/network/firewall_policies"
	firewall_zones_ds "github.com/alienchow/unitf/internal/datasources/network/firewall_zones"
	networks_ds "github.com/alienchow/unitf/internal/datasources/network/networks"
	switch_stacks_ds "github.com/alienchow/unitf/internal/datasources/network/switch_stacks"
	vpn_servers_ds "github.com/alienchow/unitf/internal/datasources/network/vpn_servers"
	vpn_tunnels_ds "github.com/alienchow/unitf/internal/datasources/network/vpn_tunnels"
	wans_ds "github.com/alienchow/unitf/internal/datasources/network/wans"
	cameras_ds "github.com/alienchow/unitf/internal/datasources/protect/cameras"
	chimes_ds "github.com/alienchow/unitf/internal/datasources/protect/chimes"
	fobs_ds "github.com/alienchow/unitf/internal/datasources/protect/fobs"
	lights_ds "github.com/alienchow/unitf/internal/datasources/protect/lights"
	liveviews_ds "github.com/alienchow/unitf/internal/datasources/protect/liveviews"
	nvr_ds "github.com/alienchow/unitf/internal/datasources/protect/nvr"
	relays_ds "github.com/alienchow/unitf/internal/datasources/protect/relays"
	sensors_ds "github.com/alienchow/unitf/internal/datasources/protect/sensors"
	sirens_ds "github.com/alienchow/unitf/internal/datasources/protect/sirens"
	speakers_ds "github.com/alienchow/unitf/internal/datasources/protect/speakers"
	users_ds "github.com/alienchow/unitf/internal/datasources/protect/users"
	viewers_ds "github.com/alienchow/unitf/internal/datasources/protect/viewers"
	acl_rule_res "github.com/alienchow/unitf/internal/resources/network/acl_rule"
	acl_rule_ordering_res "github.com/alienchow/unitf/internal/resources/network/acl_rule_ordering"
	device_res "github.com/alienchow/unitf/internal/resources/network/device"
	dns_policy_res "github.com/alienchow/unitf/internal/resources/network/dns_policy"
	firewall_policy_res "github.com/alienchow/unitf/internal/resources/network/firewall_policy"
	firewall_policy_ordering_res "github.com/alienchow/unitf/internal/resources/network/firewall_policy_ordering"
	firewall_zone_res "github.com/alienchow/unitf/internal/resources/network/firewall_zone"
	hotspot_voucher_res "github.com/alienchow/unitf/internal/resources/network/hotspot_voucher"
	network_res "github.com/alienchow/unitf/internal/resources/network/network"
	traffic_matching_list_res "github.com/alienchow/unitf/internal/resources/network/traffic_matching_list"
	wifi_broadcast_res "github.com/alienchow/unitf/internal/resources/network/wifi_broadcast"
	protect_camera_res "github.com/alienchow/unitf/internal/resources/protect/camera"
	protect_chime_res "github.com/alienchow/unitf/internal/resources/protect/chime"
	protect_light_res "github.com/alienchow/unitf/internal/resources/protect/light"
	protect_liveview_res "github.com/alienchow/unitf/internal/resources/protect/liveview"
	protect_relay_res "github.com/alienchow/unitf/internal/resources/protect/relay"
	protect_sensor_res "github.com/alienchow/unitf/internal/resources/protect/sensor"
	protect_siren_res "github.com/alienchow/unitf/internal/resources/protect/siren"
	protect_viewer_res "github.com/alienchow/unitf/internal/resources/protect/viewer"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ provider.Provider = &UniFiProvider{}

type UniFiProvider struct {
	version string
}

type UniFiProviderModel struct {
	Host     types.String `tfsdk:"host"`
	APIKey   types.String `tfsdk:"api_key"`
	SiteID   types.String `tfsdk:"site_id"`
	Insecure types.Bool   `tfsdk:"insecure"`
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &UniFiProvider{version: version}
	}
}

func (p *UniFiProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "unifi"
	resp.Version = p.version
}

func (p *UniFiProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "OpenTofu-native provider for UniFi Network and UniFi Protect.",
		Attributes: map[string]schema.Attribute{
			"host": schema.StringAttribute{
				Optional:    true,
				Description: "URL or IP address of the UniFi Console/UDM/NVR (e.g. https://192.168.1.1). Env: UNIFI_HOST.",
			},
			"api_key": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "API Key generated from the UniFi application settings. Env: UNIFI_API_KEY.",
			},
			"site_id": schema.StringAttribute{
				Optional:    true,
				Description: "Default UniFi site ID (e.g. default). Env: UNIFI_SITE_ID.",
			},
			"insecure": schema.BoolAttribute{
				Optional:    true,
				Description: "Skip TLS verification for self-signed certificates. Env: UNIFI_INSECURE.",
			},
		},
	}
}

func (p *UniFiProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config UniFiProviderModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	host := os.Getenv("UNIFI_HOST")
	if !config.Host.IsNull() {
		host = config.Host.ValueString()
	}

	apiKey := os.Getenv("UNIFI_API_KEY")
	if !config.APIKey.IsNull() {
		apiKey = config.APIKey.ValueString()
	}

	insecure := false
	if os.Getenv("UNIFI_INSECURE") == "true" || os.Getenv("UNIFI_INSECURE") == "1" {
		insecure = true
	}
	if !config.Insecure.IsNull() {
		insecure = config.Insecure.ValueBool()
	}

	if host == "" {
		resp.Diagnostics.AddError("Missing Host", "UniFi host must be specified in provider configuration or UNIFI_HOST env var.")
		return
	}

	if apiKey == "" {
		resp.Diagnostics.AddError("Missing API Key", "UniFi API Key must be specified in provider configuration or UNIFI_API_KEY env var.")
		return
	}

	apiClient, err := client.NewClient(host, apiKey, insecure)
	if err != nil {
		resp.Diagnostics.AddError("Failed to Initialize UniFi Client", err.Error())
		return
	}

	resp.DataSourceData = apiClient
	resp.ResourceData = apiClient
}

func (p *UniFiProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		network_res.NewNetworkResource,
		firewall_zone_res.NewFirewallZoneResource,
		firewall_policy_res.NewFirewallPolicyResource,
		firewall_policy_ordering_res.NewFirewallPolicyOrderingResource,
		wifi_broadcast_res.NewWifiBroadcastResource,
		acl_rule_res.NewAclRuleResource,
		acl_rule_ordering_res.NewAclRuleOrderingResource,
		device_res.NewDeviceResource,
		dns_policy_res.NewDnsPolicyResource,
		hotspot_voucher_res.NewHotspotVoucherResource,
		traffic_matching_list_res.NewTrafficMatchingListResource,
		protect_camera_res.NewProtectCameraResource,
		protect_sensor_res.NewProtectSensorResource,
		protect_light_res.NewProtectLightResource,
		protect_relay_res.NewProtectRelayResource,
		protect_siren_res.NewProtectSirenResource,
		protect_chime_res.NewProtectChimeResource,
		protect_viewer_res.NewProtectViewerResource,
		protect_liveview_res.NewProtectLiveviewResource,
	}
}

func (p *UniFiProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		sites_ds.NewSitesDataSource,
		traffic_matching_lists_ds.NewTrafficMatchingListsDataSource,
		dns_policies_ds.NewDnsPoliciesDataSource,
		acl_rules_ds.NewAclRulesDataSource,
		wifi_broadcasts_ds.NewWifiBroadcastsDataSource,
		firewall_policies_ds.NewFirewallPoliciesDataSource,
		firewall_zones_ds.NewFirewallZonesDataSource,
		networks_ds.NewNetworksDataSource,
		wans_ds.NewWansDataSource,
		devices_ds.NewDevicesDataSource,
		device_statistics_ds.NewDeviceStatisticsDataSource,
		clients_ds.NewClientsDataSource,
		countries_ds.NewCountriesDataSource,
		dpi_applications_ds.NewDpiApplicationsDataSource,
		dpi_categories_ds.NewDpiCategoriesDataSource,
		device_tags_ds.NewDeviceTagsDataSource,
		radius_profiles_ds.NewRadiusProfilesDataSource,
		vpn_servers_ds.NewVpnServersDataSource,
		vpn_tunnels_ds.NewVpnTunnelsDataSource,
		lags_ds.NewLagsDataSource,
		mc_lag_domains_ds.NewMcLagDomainsDataSource,
		switch_stacks_ds.NewSwitchStacksDataSource,
		cameras_ds.NewCamerasDataSource,
		nvr_ds.NewNvrDataSource,
		sensors_ds.NewSensorsDataSource,
		lights_ds.NewLightsDataSource,
		speakers_ds.NewSpeakersDataSource,
		liveviews_ds.NewLiveviewsDataSource,
		relays_ds.NewRelaysDataSource,
		sirens_ds.NewSirensDataSource,
		chimes_ds.NewChimesDataSource,
		viewers_ds.NewViewersDataSource,
		fobs_ds.NewFobsDataSource,
		users_ds.NewUsersDataSource,
	}
}
