package client_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/alienchow/unitf/internal/client"
)

func loadTestData(t *testing.T, filename string) []byte {
	t.Helper()
	data, err := os.ReadFile(filepath.Join("testdata", filename))
	if err != nil {
		t.Fatalf("Failed to read testdata %s: %v", filename, err)
	}
	return data
}

func TestClient_CreateDnsPolicy(t *testing.T) {
	mockResp := loadTestData(t, "dns_policy_response.json")

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected POST method, got %s", r.Method)
		}
		if r.URL.Path != "/proxy/network/integration/v1/sites/site123/dns/policies" {
			t.Errorf("Unexpected path: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(mockResp)
	}))
	defer ts.Close()

	c, err := client.NewClient(ts.URL, "test-key", true)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	req := &client.DnsPolicyDto{
		Domain:      "unifi.internal",
		Enabled:     true,
		Type:        "A_RECORD",
		IPv4Address: "192.168.1.1",
		TTLSeconds:  300,
	}

	res, err := c.CreateDnsPolicy(context.Background(), "site123", req)
	if err != nil {
		t.Fatalf("CreateDnsPolicy returned error: %v", err)
	}
	if res.ID != "d367f7dd-a077-44d4-81e6-5327388d3561" {
		t.Errorf("Unexpected ID: %s", res.ID)
	}
	if res.Domain != "unifi.internal" {
		t.Errorf("Unexpected Domain: %s", res.Domain)
	}
}

func TestClient_GetDnsPolicy(t *testing.T) {
	mockResp := loadTestData(t, "dns_policy_response.json")

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET method, got %s", r.Method)
		}
		if r.URL.Path != "/proxy/network/integration/v1/sites/site123/dns/policies/policy123" {
			t.Errorf("Unexpected path: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(mockResp)
	}))
	defer ts.Close()

	c, err := client.NewClient(ts.URL, "test-key", true)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	res, err := c.GetDnsPolicy(context.Background(), "site123", "policy123")
	if err != nil {
		t.Fatalf("GetDnsPolicy returned error: %v", err)
	}
	if res.IPv4Address != "192.168.1.1" {
		t.Errorf("Unexpected IPv4Address: %s", res.IPv4Address)
	}
}

func TestClient_UpdateDnsPolicy(t *testing.T) {
	mockResp := loadTestData(t, "dns_policy_response.json")

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			t.Errorf("Expected PUT method, got %s", r.Method)
		}
		if r.URL.Path != "/proxy/network/integration/v1/sites/site123/dns/policies/policy123" {
			t.Errorf("Unexpected path: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(mockResp)
	}))
	defer ts.Close()

	c, err := client.NewClient(ts.URL, "test-key", true)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	req := &client.DnsPolicyDto{
		Domain:      "unifi.internal",
		Enabled:     true,
		Type:        "A_RECORD",
		IPv4Address: "192.168.1.1",
		TTLSeconds:  300,
	}

	res, err := c.UpdateDnsPolicy(context.Background(), "site123", "policy123", req)
	if err != nil {
		t.Fatalf("UpdateDnsPolicy returned error: %v", err)
	}
	if res.Type != "A_RECORD" {
		t.Errorf("Unexpected Type: %s", res.Type)
	}
}

func TestClient_DeleteDnsPolicy(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			t.Errorf("Expected DELETE method, got %s", r.Method)
		}
		if r.URL.Path != "/proxy/network/integration/v1/sites/site123/dns/policies/policy123" {
			t.Errorf("Unexpected path: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer ts.Close()

	c, err := client.NewClient(ts.URL, "test-key", true)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	err = c.DeleteDnsPolicy(context.Background(), "site123", "policy123")
	if err != nil {
		t.Fatalf("DeleteDnsPolicy returned error: %v", err)
	}
}

func TestClient_ListDnsPolicies(t *testing.T) {
	mockResp := loadTestData(t, "dns_policies_list_response.json")

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET method, got %s", r.Method)
		}
		if r.URL.Path != "/proxy/network/integration/v1/sites/site123/dns/policies" {
			t.Errorf("Unexpected path: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(mockResp)
	}))
	defer ts.Close()

	c, err := client.NewClient(ts.URL, "test-key", true)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	items, err := c.ListDnsPolicies(context.Background(), "site123")
	if err != nil {
		t.Fatalf("ListDnsPolicies returned error: %v", err)
	}
	if len(items) != 1 {
		t.Fatalf("Expected 1 item, got %d", len(items))
	}
	if items[0].Domain != "unifi.internal" {
		t.Errorf("Unexpected domain: %s", items[0].Domain)
	}
}
