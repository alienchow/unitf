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
		name         string
		inputFile    string
		expectedFile string
	}{
		{
			name:         "unifi provider",
			inputFile:    "clean_provider_input_unifi.tf",
			expectedFile: "clean_provider_expected_unifi.tf",
		},
		{
			name:         "unitf provider",
			inputFile:    "clean_provider_input_unitf.tf",
			expectedFile: "clean_provider_expected_unitf.tf",
		},
		{
			name:         "no provider line",
			inputFile:    "clean_provider_input_none.tf",
			expectedFile: "clean_provider_expected_none.tf",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := getTestData(t, tt.inputFile)
			expected := getTestData(t, tt.expectedFile)

			result := cleanProviderAttribute(input)
			if strings.TrimSpace(result) != strings.TrimSpace(expected) {
				t.Errorf("cleanProviderAttribute failed for %s.\nGot:\n%s\nExpected:\n%s", tt.name, result, expected)
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

func TestRunCommand(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "cmd_test_*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	if err := runCommand(tempDir, "echo", "hello"); err != nil {
		t.Errorf("runCommand failed: %v", err)
	}

	if err := runCommand(tempDir, "false"); err == nil {
		t.Errorf("Expected error from 'false' command, got nil")
	}
}

func setupMockBinary(t *testing.T) string {
	t.Helper()
	tempBinDir, err := os.MkdirTemp("", "mock_bin_*")
	if err != nil {
		t.Fatalf("Failed to create mock bin dir: %v", err)
	}

	mockTofu := filepath.Join(tempBinDir, "tofu")
	mockScript := "#!/bin/sh\n" +
		"if [ \"$1\" = \"plan\" ]; then\n" +
		"  touch baseline.tf\n" +
		"fi\n" +
		"exit 0\n"

	// #nosec G304,G703,G306
	if err := os.WriteFile(mockTofu, []byte(mockScript), 0755); err != nil {
		t.Fatalf("Failed to write mock binary: %v", err)
	}

	oldPATH := os.Getenv("PATH")
	t.Setenv("PATH", tempBinDir+":"+oldPATH)

	return tempBinDir
}

func TestRunStep1GenerateImports(t *testing.T) {
	mockBinDir := setupMockBinary(t)
	defer os.RemoveAll(mockBinDir)

	tempDir, err := os.MkdirTemp("", "step1_test_*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	if err := runStep1GenerateImports(tempDir); err != nil {
		t.Fatalf("runStep1GenerateImports failed: %v", err)
	}

	providersFile := filepath.Join(tempDir, "providers.tf")
	if _, err := os.Stat(providersFile); os.IsNotExist(err) {
		t.Errorf("providers.tf was not created by Step 1")
	}
}

func TestRunStep2GenerateBaseline(t *testing.T) {
	mockBinDir := setupMockBinary(t)
	defer os.RemoveAll(mockBinDir)

	tempDir, err := os.MkdirTemp("", "step2_test_*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	if err := runStep2GenerateBaseline(tempDir); err != nil {
		t.Fatalf("runStep2GenerateBaseline failed: %v", err)
	}

	baselineFile := filepath.Join(tempDir, "baseline.tf")
	if _, err := os.Stat(baselineFile); os.IsNotExist(err) {
		t.Errorf("baseline.tf was not created by Step 2")
	}
}

func TestRefactorAndCleanup_UnmappedResourceType(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "setup_unmapped_test_*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	content := `# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/unmapped-1"
resource "unifi_unknown_custom_resource" "unknown_1" {
  name    = "custom"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
}`

	baselineFile := filepath.Join(tempDir, "baseline.tf")
	// #nosec G304,G703
	if err := os.WriteFile(baselineFile, []byte(content), 0600); err != nil {
		t.Fatalf("Failed to write baseline.tf: %v", err)
	}

	if err := runStep3RefactorAndCleanup(tempDir); err != nil {
		t.Fatalf("runStep3RefactorAndCleanup failed: %v", err)
	}

	resFile := filepath.Join(tempDir, "resources.tf")
	// #nosec G304
	resContent, err := os.ReadFile(resFile)
	if err != nil {
		t.Fatalf("resources.tf was not generated for unmapped resource: %v", err)
	}

	if !strings.Contains(string(resContent), "unifi_unknown_custom_resource") {
		t.Errorf("resources.tf does not contain unknown resource")
	}
}

func TestRefactorAndCleanup_Sample(t *testing.T) {
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

	if _, err := os.Stat(baselineFile); !os.IsNotExist(err) {
		t.Errorf("baseline.tf was not deleted")
	}

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
}

func TestRefactorAndCleanup_EnterpriseFullTopology(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "setup_enterprise_test_*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	content := getTestData(t, "baseline_enterprise_full_topology.tf")

	baselineFile := filepath.Join(tempDir, "baseline.tf")
	// #nosec G304,G703
	if err := os.WriteFile(baselineFile, []byte(content), 0600); err != nil {
		t.Fatalf("Failed to write baseline.tf: %v", err)
	}

	if err := runStep3RefactorAndCleanup(tempDir); err != nil {
		t.Fatalf("runStep3RefactorAndCleanup failed on enterprise topology baseline: %v", err)
	}

	expectedFiles := []string{
		"acl_rules.tf",
		"acl_rule_orderings.tf",
		"devices.tf",
		"dns_policies.tf",
		"firewall_policies.tf",
		"firewall_policy_orderings.tf",
		"firewall_zones.tf",
		"hotspot_vouchers.tf",
		"networks.tf",
		"traffic_matching_lists.tf",
		"wifi_broadcasts.tf",
		"protect_cameras.tf",
		"protect_chimes.tf",
		"protect_lights.tf",
		"protect_liveviews.tf",
		"protect_relays.tf",
		"protect_sensors.tf",
		"protect_sirens.tf",
		"protect_viewers.tf",
	}

	for _, fName := range expectedFiles {
		fPath := filepath.Join(tempDir, fName)
		// #nosec G304
		data, err := os.ReadFile(fPath)
		if err != nil {
			t.Errorf("Expected segmented file %s was not generated: %v", fName, err)
			continue
		}
		if len(data) == 0 {
			t.Errorf("Segmented file %s is unexpectedly empty", fName)
		}
	}
}

func TestRefactorAndCleanup_NonExistentBaseline(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "setup_test_empty_*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	if err := runStep3RefactorAndCleanup(tempDir); err != nil {
		t.Errorf("Expected no error when baseline.tf missing, got %v", err)
	}
}
