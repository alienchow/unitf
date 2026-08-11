package client_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alienchow/unitf/internal/client"
)

func TestClient_Networks_MassivePayload(t *testing.T) {
	mockData := loadClientTestData(t, "networks_massive.json")
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(mockData)
	}))
	defer ts.Close()

	c, err := client.NewClient(ts.URL, "test-key", true)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()
	siteID := "site123"

	nets, err := c.ListNetworks(ctx, siteID)
	if err != nil {
		t.Fatalf("ListNetworks failed: %v", err)
	}
	if len(nets) != 20 {
		t.Fatalf("Expected 20 networks, got %d", len(nets))
	}
	if nets[0].Name != "VLAN-101-Enterprise-Segment-1" {
		t.Errorf("Unexpected first network name: %s", nets[0].Name)
	}

	// Test single CRUD operations with massive DTO fields
	_, err = c.GetNetwork(ctx, siteID, nets[0].ID)
	if err != nil {
		t.Errorf("GetNetwork failed: %v", err)
	}
	_, err = c.CreateNetwork(ctx, siteID, &nets[0])
	if err != nil {
		t.Errorf("CreateNetwork failed: %v", err)
	}
	_, err = c.UpdateNetwork(ctx, siteID, nets[0].ID, &nets[0])
	if err != nil {
		t.Errorf("UpdateNetwork failed: %v", err)
	}
	err = c.DeleteNetwork(ctx, siteID, nets[0].ID)
	if err != nil {
		t.Errorf("DeleteNetwork failed: %v", err)
	}
}

func TestClient_FirewallPolicies_MassivePayload(t *testing.T) {
	mockData := loadClientTestData(t, "firewall_policies_massive.json")
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(mockData)
	}))
	defer ts.Close()

	c, err := client.NewClient(ts.URL, "test-key", true)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()
	siteID := "site123"

	pols, err := c.ListFirewallPolicies(ctx, siteID)
	if err != nil {
		t.Fatalf("ListFirewallPolicies failed: %v", err)
	}
	if len(pols) != 30 {
		t.Fatalf("Expected 30 firewall policies, got %d", len(pols))
	}

	_, err = c.GetFirewallPolicy(ctx, siteID, pols[0].ID)
	if err != nil {
		t.Errorf("GetFirewallPolicy failed: %v", err)
	}
	_, err = c.CreateFirewallPolicy(ctx, siteID, &pols[0])
	if err != nil {
		t.Errorf("CreateFirewallPolicy failed: %v", err)
	}
	_, err = c.UpdateFirewallPolicy(ctx, siteID, pols[0].ID, &pols[0])
	if err != nil {
		t.Errorf("UpdateFirewallPolicy failed: %v", err)
	}
	err = c.DeleteFirewallPolicy(ctx, siteID, pols[0].ID)
	if err != nil {
		t.Errorf("DeleteFirewallPolicy failed: %v", err)
	}
}

func TestClient_FirewallZones_MassivePayload(t *testing.T) {
	mockData := loadClientTestData(t, "firewall_zones_massive.json")
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(mockData)
	}))
	defer ts.Close()

	c, err := client.NewClient(ts.URL, "test-key", true)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()
	siteID := "site123"

	zones, err := c.ListFirewallZones(ctx, siteID)
	if err != nil {
		t.Fatalf("ListFirewallZones failed: %v", err)
	}
	if len(zones) != 15 {
		t.Fatalf("Expected 15 zones, got %d", len(zones))
	}

	_, err = c.GetFirewallZone(ctx, siteID, zones[0].ID)
	if err != nil {
		t.Errorf("GetFirewallZone failed: %v", err)
	}
	_, err = c.CreateFirewallZone(ctx, siteID, &zones[0])
	if err != nil {
		t.Errorf("CreateFirewallZone failed: %v", err)
	}
	_, err = c.UpdateFirewallZone(ctx, siteID, zones[0].ID, &zones[0])
	if err != nil {
		t.Errorf("UpdateFirewallZone failed: %v", err)
	}
	err = c.DeleteFirewallZone(ctx, siteID, zones[0].ID)
	if err != nil {
		t.Errorf("DeleteFirewallZone failed: %v", err)
	}
}

func TestClient_AclRules_MassivePayload(t *testing.T) {
	mockData := loadClientTestData(t, "acl_rules_massive.json")
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(mockData)
	}))
	defer ts.Close()

	c, err := client.NewClient(ts.URL, "test-key", true)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()
	siteID := "site123"

	acls, err := c.ListAclRules(ctx, siteID)
	if err != nil {
		t.Fatalf("ListAclRules failed: %v", err)
	}
	if len(acls) != 20 {
		t.Fatalf("Expected 20 ACL rules, got %d", len(acls))
	}

	_, err = c.GetAclRule(ctx, siteID, acls[0].ID)
	if err != nil {
		t.Errorf("GetAclRule failed: %v", err)
	}
	_, err = c.CreateAclRule(ctx, siteID, &acls[0])
	if err != nil {
		t.Errorf("CreateAclRule failed: %v", err)
	}
	_, err = c.UpdateAclRule(ctx, siteID, acls[0].ID, &acls[0])
	if err != nil {
		t.Errorf("UpdateAclRule failed: %v", err)
	}
	err = c.DeleteAclRule(ctx, siteID, acls[0].ID)
	if err != nil {
		t.Errorf("DeleteAclRule failed: %v", err)
	}
}

func TestClient_Devices_MassivePayload(t *testing.T) {
	mockData := loadClientTestData(t, "devices_massive.json")
	tsList := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(mockData)
	}))
	defer tsList.Close()

	cList, err := client.NewClient(tsList.URL, "test-key", true)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()
	siteID := "site123"

	devs, err := cList.ListDevices(ctx, siteID)
	if err != nil {
		t.Fatalf("ListDevices failed: %v", err)
	}
	if len(devs) != 20 {
		t.Fatalf("Expected 20 devices, got %d", len(devs))
	}

	tsItem := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"data": [{"id": "dev-1", "mac": "74:83:c2:01:11:22"}]}`))
	}))
	defer tsItem.Close()

	cItem, err := client.NewClient(tsItem.URL, "test-key", true)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	_, err = cItem.GetDevice(ctx, siteID, "74:83:c2:01:11:22")
	if err != nil {
		t.Errorf("GetDevice failed: %v", err)
	}
	_, err = cItem.UpdateDevice(ctx, siteID, devs[0].ID, &client.DeviceDto{Name: "Updated-USW"})
	if err != nil {
		t.Errorf("UpdateDevice failed: %v", err)
	}
	err = cItem.AdoptDevice(ctx, siteID, "74:83:c2:01:11:22")
	if err != nil {
		t.Errorf("AdoptDevice failed: %v", err)
	}
	err = cItem.ForgetDevice(ctx, siteID, "74:83:c2:01:11:22")
	if err != nil {
		t.Errorf("ForgetDevice failed: %v", err)
	}
}

func TestClient_TrafficMatching_MassivePayload(t *testing.T) {
	mockData := loadClientTestData(t, "traffic_matching_lists_massive.json")
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(mockData)
	}))
	defer ts.Close()

	c, err := client.NewClient(ts.URL, "test-key", true)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()
	siteID := "site123"

	tmls, err := c.ListTrafficMatchingLists(ctx, siteID)
	if err != nil {
		t.Fatalf("ListTrafficMatchingLists failed: %v", err)
	}
	if len(tmls) != 15 {
		t.Fatalf("Expected 15 traffic matching lists, got %d", len(tmls))
	}

	_, err = c.GetTrafficMatchingList(ctx, siteID, tmls[0].ID)
	if err != nil {
		t.Errorf("GetTrafficMatchingList failed: %v", err)
	}
	_, err = c.CreateTrafficMatchingList(ctx, siteID, &tmls[0])
	if err != nil {
		t.Errorf("CreateTrafficMatchingList failed: %v", err)
	}
	_, err = c.UpdateTrafficMatchingList(ctx, siteID, tmls[0].ID, &tmls[0])
	if err != nil {
		t.Errorf("UpdateTrafficMatchingList failed: %v", err)
	}
	err = c.DeleteTrafficMatchingList(ctx, siteID, tmls[0].ID)
	if err != nil {
		t.Errorf("DeleteTrafficMatchingList failed: %v", err)
	}
}

func TestClient_WifiBroadcasts_MassivePayload(t *testing.T) {
	mockData := loadClientTestData(t, "wifi_broadcasts_massive.json")
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(mockData)
	}))
	defer ts.Close()

	c, err := client.NewClient(ts.URL, "test-key", true)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()
	siteID := "site123"

	wifis, err := c.ListWifiBroadcasts(ctx, siteID)
	if err != nil {
		t.Fatalf("ListWifiBroadcasts failed: %v", err)
	}
	if len(wifis) != 10 {
		t.Fatalf("Expected 10 wifi broadcasts, got %d", len(wifis))
	}

	_, err = c.GetWifiBroadcast(ctx, siteID, wifis[0].ID)
	if err != nil {
		t.Errorf("GetWifiBroadcast failed: %v", err)
	}
	_, err = c.CreateWifiBroadcast(ctx, siteID, &wifis[0])
	if err != nil {
		t.Errorf("CreateWifiBroadcast failed: %v", err)
	}
	_, err = c.UpdateWifiBroadcast(ctx, siteID, wifis[0].ID, &wifis[0])
	if err != nil {
		t.Errorf("UpdateWifiBroadcast failed: %v", err)
	}
	err = c.DeleteWifiBroadcast(ctx, siteID, wifis[0].ID)
	if err != nil {
		t.Errorf("DeleteWifiBroadcast failed: %v", err)
	}
}

func TestClient_NetworkStubs(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`[{"id": "stub-1"}]`))
	}))
	defer ts.Close()

	c, err := client.NewClient(ts.URL, "test-key", true)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()
	siteID := "site123"

	_, _ = c.ListClients(ctx, siteID)
	_, _ = c.ListCountries(ctx, siteID)
	_, _ = c.ListDeviceStatistics(ctx, siteID)
	_, _ = c.ListDpiApplications(ctx, siteID)
	_, _ = c.ListDpiCategories(ctx, siteID)
	_, _ = c.ListLags(ctx, siteID)
	_, _ = c.ListMcLagDomains(ctx, siteID)
	_, _ = c.ListRadiusProfiles(ctx, siteID)
	_, _ = c.ListSwitchStacks(ctx, siteID)
	_, _ = c.ListVpnServers(ctx, siteID)
	_, _ = c.ListVpnTunnels(ctx, siteID)
	_, _ = c.ListWans(ctx, siteID)
}
