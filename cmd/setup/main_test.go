package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCleanProviderAttribute(t *testing.T) {
	input := `resource "unifi_network" "test" {
  provider = unifi
  name     = "IoT"
}`
	expected := `resource "unifi_network" "test" {
  name     = "IoT"
}`
	result := cleanProviderAttribute(input)
	if strings.TrimSpace(result) != strings.TrimSpace(expected) {
		t.Errorf("cleanProviderAttribute failed.\nGot:\n%s\nExpected:\n%s", result, expected)
	}
}

func TestExtractSiteID(t *testing.T) {
	content, err := os.ReadFile(filepath.Join("testdata", "baseline_sample.tf"))
	if err != nil {
		t.Fatalf("Failed to load testdata: %v", err)
	}

	siteID := extractSiteID(string(content))
	expected := "88f7af54-98f8-306a-a1c7-c9349722b1f6"
	if siteID != expected {
		t.Errorf("extractSiteID failed. Got %q, expected %q", siteID, expected)
	}
}

func TestGetTofuOrTerraform(t *testing.T) {
	cmd := getTofuOrTerraform()
	if cmd != "tofu" && cmd != "terraform" {
		t.Errorf("getTofuOrTerraform returned unexpected binary: %s", cmd)
	}
}

func TestRefactorAndCleanup(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "setup_test_*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	content, err := os.ReadFile(filepath.Join("testdata", "baseline_sample.tf"))
	if err != nil {
		t.Fatalf("Failed to load testdata: %v", err)
	}

	baselineFile := filepath.Join(tempDir, "baseline.tf")
	// #nosec G304,G703
	if err := os.WriteFile(baselineFile, content, 0600); err != nil {
		t.Fatalf("Failed to write baseline.tf: %v", err)
	}

	if err := runStep3RefactorAndCleanup(tempDir); err != nil {
		t.Fatalf("runStep3RefactorAndCleanup failed: %v", err)
	}

	// Verify baseline.tf was removed
	if _, err := os.Stat(baselineFile); !os.IsNotExist(err) {
		t.Errorf("baseline.tf was not deleted")
	}

	// Verify networks.tf was generated
	netFile := filepath.Join(tempDir, "networks.tf")
	// #nosec G304
	netContent, err := os.ReadFile(netFile)
	if err != nil {
		t.Fatalf("networks.tf was not generated: %v", err)
	}

	if !strings.Contains(string(netContent), "unifi_network") {
		t.Errorf("networks.tf does not contain unifi_network resource")
	}

	if !strings.Contains(string(netContent), "local.site_id") {
		t.Errorf("networks.tf site_id was not replaced with local.site_id")
	}

	// Verify firewall_zones.tf was generated
	fzFile := filepath.Join(tempDir, "firewall_zones.tf")
	// #nosec G304
	if _, err := os.ReadFile(fzFile); err != nil {
		t.Fatalf("firewall_zones.tf was not generated: %v", err)
	}

	// Verify protect_cameras.tf was generated
	camFile := filepath.Join(tempDir, "protect_cameras.tf")
	// #nosec G304
	if _, err := os.ReadFile(camFile); err != nil {
		t.Fatalf("protect_cameras.tf was not generated: %v", err)
	}
}
