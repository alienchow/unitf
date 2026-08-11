package client_test

import (
	"context"
	"embed"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alienchow/unitf/internal/client"
)

//go:embed testdata/*
var clientTestDataFS embed.FS

func loadClientTestData(t *testing.T, filename string) []byte {
	t.Helper()
	data, err := clientTestDataFS.ReadFile("testdata/" + filename)
	if err != nil {
		t.Fatalf("Failed to read embedded testdata %s: %v", filename, err)
	}
	return data
}

func TestClient_CreateDnsPolicy_Types(t *testing.T) {
	tests := []struct {
		name          string
		testdataFile  string
		req           *client.DnsPolicyDto
		expectedID    string
		expectedField string
	}{
		{
			name:         "A_RECORD",
			testdataFile: "dns_policy_response.json",
			req: &client.DnsPolicyDto{
				Domain:      "unifi.internal",
				Enabled:     true,
				Type:        "A_RECORD",
				IPv4Address: "192.168.1.1",
				TTLSeconds:  300,
			},
			expectedID:    "d367f7dd-a077-44d4-81e6-5327388d3561",
			expectedField: "192.168.1.1",
		},
		{
			name:         "AAAA_RECORD",
			testdataFile: "dns_policy_aaaa.json",
			req: &client.DnsPolicyDto{
				Domain:      "ipv6.internal",
				Enabled:     true,
				Type:        "AAAA_RECORD",
				IPv6Address: "2001:db8::1",
				TTLSeconds:  600,
			},
			expectedID:    "aaaa-1234",
			expectedField: "2001:db8::1",
		},
		{
			name:         "CNAME_RECORD",
			testdataFile: "dns_policy_cname.json",
			req: &client.DnsPolicyDto{
				Domain:       "alias.internal",
				Enabled:      true,
				Type:         "CNAME_RECORD",
				TargetDomain: "unifi.internal",
				TTLSeconds:   300,
			},
			expectedID:    "cname-1234",
			expectedField: "unifi.internal",
		},
		{
			name:         "TXT_RECORD",
			testdataFile: "dns_policy_txt.json",
			req: &client.DnsPolicyDto{
				Domain:  "info.internal",
				Enabled: true,
				Type:    "TXT_RECORD",
				Text:    "v=spf1 -all",
			},
			expectedID:    "txt-1234",
			expectedField: "v=spf1 -all",
		},
		{
			name:         "FORWARD_DOMAIN",
			testdataFile: "dns_policy_forward.json",
			req: &client.DnsPolicyDto{
				Domain:    "forward.internal",
				Enabled:   true,
				Type:      "FORWARD_DOMAIN",
				IPAddress: "1.1.1.1",
			},
			expectedID:    "forward-1234",
			expectedField: "1.1.1.1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockResp := loadClientTestData(t, tt.testdataFile)
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != "POST" {
					t.Errorf("Expected POST, got %s", r.Method)
				}
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write(mockResp)
			}))
			defer ts.Close()

			c, err := client.NewClient(ts.URL, "test-key", true)
			if err != nil {
				t.Fatalf("Failed to create client: %v", err)
			}

			res, err := c.CreateDnsPolicy(context.Background(), "site123", tt.req)
			if err != nil {
				t.Fatalf("CreateDnsPolicy failed for %s: %v", tt.name, err)
			}

			if res.ID != tt.expectedID {
				t.Errorf("Expected ID %s, got %s", tt.expectedID, res.ID)
			}
		})
	}
}

func TestClient_GetDnsPolicy_Success(t *testing.T) {
	mockResp := loadClientTestData(t, "dns_policy_response.json")

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

func TestClient_GetDnsPolicy_NotFound(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"error": "not found"}`))
	}))
	defer ts.Close()

	c, err := client.NewClient(ts.URL, "test-key", true)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	_, err = c.GetDnsPolicy(context.Background(), "site123", "nonexistent")
	if err == nil {
		t.Fatal("Expected error for 404 response, got nil")
	}
	if !client.IsNotFound(err) {
		t.Errorf("Expected IsNotFound(err) to be true")
	}
}

func TestClient_CreateDnsPolicy_Error400(t *testing.T) {
	errResp := loadClientTestData(t, "api_error_400.json")

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(errResp)
	}))
	defer ts.Close()

	c, err := client.NewClient(ts.URL, "test-key", true)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	_, err = c.CreateDnsPolicy(context.Background(), "site123", &client.DnsPolicyDto{})
	if err == nil {
		t.Fatal("Expected error for 400 response, got nil")
	}
}

func TestClient_CreateDnsPolicy_Error500(t *testing.T) {
	errResp := loadClientTestData(t, "api_error_500.json")

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(errResp)
	}))
	defer ts.Close()

	c, err := client.NewClient(ts.URL, "test-key", true)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	_, err = c.CreateDnsPolicy(context.Background(), "site123", &client.DnsPolicyDto{})
	if err == nil {
		t.Fatal("Expected error for 500 server error, got nil")
	}
}

func TestClient_GetDnsPolicy_MalformedJSON(t *testing.T) {
	malformedResp := loadClientTestData(t, "malformed_json.json")

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(malformedResp)
	}))
	defer ts.Close()

	c, err := client.NewClient(ts.URL, "test-key", true)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	_, err = c.GetDnsPolicy(context.Background(), "site123", "policy123")
	if err == nil {
		t.Fatal("Expected JSON decode error for malformed JSON response, got nil")
	}
}

func TestClient_UpdateDnsPolicy(t *testing.T) {
	mockResp := loadClientTestData(t, "dns_policy_response.json")

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			t.Errorf("Expected PUT method, got %s", r.Method)
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
	mockResp := loadClientTestData(t, "dns_policies_list_response.json")

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET method, got %s", r.Method)
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
