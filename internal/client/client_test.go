package client_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alienchow/unitf/internal/client"
)

func TestClient_DoRequest_Success(t *testing.T) {
	// Create a mock HTTP server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Validate headers
		if r.Header.Get("X-API-Key") != "test-api-key" {
			t.Errorf("Expected X-API-Key test-api-key, got %s", r.Header.Get("X-API-Key"))
		}

		// Return a mock response
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"name": "test-network", "purpose": "corporate"}`))
	}))
	defer ts.Close()

	// Initialize the client pointing to the mock server
	c, err := client.NewClient(ts.URL, "test-api-key", "default", true)
	if err != nil {
		t.Fatalf("Expected no error creating client, got: %v", err)
	}

	// Define a mock response struct
	var resp struct {
		Name    string `json:"name"`
		Purpose string `json:"purpose"`
	}

	// Perform the request
	err = c.DoRequest(context.Background(), "GET", "/test-path", nil, &resp)
	if err != nil {
		t.Fatalf("Expected no error from DoRequest, got: %v", err)
	}

	// Validate the parsed response
	if resp.Name != "test-network" {
		t.Errorf("Expected name 'test-network', got '%s'", resp.Name)
	}
	if resp.Purpose != "corporate" {
		t.Errorf("Expected purpose 'corporate', got '%s'", resp.Purpose)
	}
}

func TestClient_DoRequest_Error(t *testing.T) {
	// Create a mock HTTP server that returns an error
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"error": "Not Found"}`))
	}))
	defer ts.Close()

	c, err := client.NewClient(ts.URL, "test-api-key", "default", true)
	if err != nil {
		t.Fatalf("Expected no error creating client, got: %v", err)
	}

	err = c.DoRequest(context.Background(), "GET", "/test-path", nil, nil)
	if err == nil {
		t.Fatal("Expected error from DoRequest for 404 response, got nil")
	}
}

func TestClient_GetNetwork(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/proxy/network/integration/v1/sites/default/networks/net-id" {
			t.Errorf("Unexpected path: %s", r.URL.Path)
		}

		mockResp := client.NetworkDto{
			Name: "Corporate LAN", Purpose: "corporate",
		}

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(mockResp)
	}))
	defer ts.Close()

	c, err := client.NewClient(ts.URL, "test-key", "default", true)
	if err != nil {
		t.Fatalf("Expected no error creating client, got: %v", err)
	}

	net, err := c.Network.GetNetwork(context.Background(), "default", "net-id")
	if err != nil {
		t.Fatalf("Expected no error calling GetNetwork, got: %v", err)
	}

	if net == nil {
		t.Fatal("Expected network DTO to not be nil")
	}

	if net.Name != "Corporate LAN" {
		t.Errorf("Expected network name 'Corporate LAN', got '%s'", net.Name)
	}
}
