package main

import (
	"embed"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

//go:embed testdata/*
var testDataFS embed.FS

func getTestData(t *testing.T, filename string) string {
	t.Helper()
	b, err := testDataFS.ReadFile("testdata/" + filename)
	if err != nil {
		t.Fatalf("Failed to load testdata %s: %v", filename, err)
	}
	return string(b)
}

func TestCleanProviderAttribute(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "unifi provider",
			input: `resource "unifi_network" "test" {
  provider = unifi
  name     = "IoT"
}`,
			expected: `resource "unifi_network" "test" {
  name     = "IoT"
}`,
		},
		{
			name: "unitf provider",
			input: `resource "unifi_network" "test" {
  provider = unitf
  name     = "IoT"
}`,
			expected: `resource "unifi_network" "test" {
  name     = "IoT"
}`,
		},
		{
			name: "no provider line",
			input: `resource "unifi_network" "test" {
  name     = "IoT"
}`,
			expected: `resource "unifi_network" "test" {
  name     = "IoT"
}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := cleanProviderAttribute(tt.input)
			if strings.TrimSpace(result) != strings.TrimSpace(tt.expected) {
				t.Errorf("cleanProviderAttribute failed for %s.\nGot:\n%s\nExpected:\n%s", tt.name, result, tt.expected)
			}
		})
	}
}

func TestExtractSiteID(t *testing.T) {
	content := getTestData(t, "baseline_sample.tf")

	siteID := extractSiteID(content)
	expected := "88f7af54-98f8-306a-a1c7-c9349722b1f6"
	if siteID != expected {
		t.Errorf("extractSiteID failed. Got %q, expected %q", siteID, expected)
	}

	// Edge case: no site_id
	if noSite := extractSiteID(`resource "test" "t" { name = "foo" }`); noSite != "" {
		t.Errorf("Expected empty string when site_id missing, got %q", noSite)
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

	content := getTestData(t, "baseline_sample.tf")

	baselineFile := filepath.Join(tempDir, "baseline.tf")
	// #nosec G304,G703
	if err := os.WriteFile(baselineFile, []byte(content), 0600); err != nil {
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

func TestRefactorAndCleanup_NonExistentBaseline(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "setup_test_empty_*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Should handle missing baseline gracefully without erroring out
	if err := runStep3RefactorAndCleanup(tempDir); err != nil {
		t.Errorf("Expected no error when baseline.tf missing, got %v", err)
	}
}
