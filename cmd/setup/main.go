package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

var fileMapping = map[string]string{
	"unifi_network":               "networks.tf",
	"unifi_firewall_zone":         "firewall_zones.tf",
	"unifi_firewall_policy":       "firewall_policies.tf",
	"unifi_dns_policy":            "dns_policies.tf",
	"unifi_traffic_matching_list": "traffic_matching_lists.tf",
	"unifi_protect_camera":        "protect_cameras.tf",
}

func main() {
	dirFlag := flag.String("dir", ".", "Directory of the OpenTofu/Terraform project")
	flag.Parse()

	absDir, err := filepath.Abs(*dirFlag)
	if err != nil {
		log.Fatalf("Failed to get absolute path: %v", err)
	}

	log.Printf("Initializing workspace at: %s", absDir)

	if err := runStep1GenerateImports(absDir); err != nil {
		log.Fatalf("Step 1 (Generate Explicit Imports) failed: %v", err)
	}

	if err := runStep2GenerateBaseline(absDir); err != nil {
		log.Fatalf("Step 2 (Generate Baseline) failed: %v", err)
	}

	if err := runStep3RefactorAndCleanup(absDir); err != nil {
		log.Fatalf("Step 3 (Refactor & Cleanup) failed: %v", err)
	}

	log.Println("Workspace setup complete successfully!")
}

func runCommand(dir, name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func runStep1GenerateImports(dir string) error {
	log.Println("==> Step 1: Generating explicit imports...")

	providersFile := filepath.Join(dir, "providers.tf")
	if _, err := os.Stat(providersFile); os.IsNotExist(err) {
		providersContent := `terraform {
  required_providers {
    unifi = {
      source  = "alienchow/unitf"
      version = "0.2.4"
    }
  }
}

provider "unifi" {
  insecure = true
}

data "unifi_sites" "default" {}

locals {
  site_id = data.unifi_sites.default.sites[0].id
}
`
		if err := os.WriteFile(providersFile, []byte(providersContent), 0600); err != nil {
			return fmt.Errorf("failed to write providers.tf: %w", err)
		}
	}

	initDiscoveryFile := filepath.Join(dir, "init_discovery.tf")
	discoveryContent := `data "unifi_networks" "all" { site_id = local.site_id }
data "unifi_firewall_policies" "all" { site_id = local.site_id }
data "unifi_firewall_zones" "all" { site_id = local.site_id }
data "unifi_dns_policies" "all" { site_id = local.site_id }
data "unifi_traffic_matching_lists" "all" { site_id = local.site_id }
data "unifi_cameras" "all" {}

resource "local_file" "explicit_imports" {
  filename = "explicit_imports.tf"
  content = join("\n", concat([
    for item in data.unifi_networks.all.items :
    "import {\n  id = \"${local.site_id}/${item.id}\"\n  to = unifi_network.net_${replace(item.id, "-", "_")}\n}"
    ], [
    for item in data.unifi_firewall_policies.all.items :
    "import {\n  id = \"${local.site_id}/${item.id}\"\n  to = unifi_firewall_policy.fw_${replace(item.id, "-", "_")}\n}"
    ], [
    for item in data.unifi_firewall_zones.all.items :
    "import {\n  id = \"${local.site_id}/${item.id}\"\n  to = unifi_firewall_zone.zone_${replace(item.id, "-", "_")}\n}"
    ], [
    for item in data.unifi_dns_policies.all.items :
    "import {\n  id = \"${local.site_id}/${item.id}\"\n  to = unifi_dns_policy.dns_${replace(item.id, "-", "_")}\n}"
    ], [
    for item in data.unifi_traffic_matching_lists.all.items :
    "import {\n  id = \"${local.site_id}/${item.id}\"\n  to = unifi_traffic_matching_list.traffic_${replace(item.id, "-", "_")}\n}"
    ], [
    for cam in data.unifi_cameras.all.items :
    "import {\n  id = \"${cam.id}\"\n  to = unifi_protect_camera.cam_${replace(cam.id, "-", "_")}\n}"
  ]))
}
`
	if err := os.WriteFile(initDiscoveryFile, []byte(discoveryContent), 0600); err != nil {
		return fmt.Errorf("failed to write init_discovery.tf: %w", err)
	}

	tfCmd := getTofuOrTerraform()

	if err := runCommand(dir, tfCmd, "init"); err != nil {
		return fmt.Errorf("failed to run init: %w", err)
	}

	if err := runCommand(dir, tfCmd, "apply", "-target=local_file.explicit_imports", "-auto-approve"); err != nil {
		return fmt.Errorf("failed to apply explicit imports generator: %w", err)
	}

	_ = runCommand(dir, tfCmd, "state", "rm", "local_file.explicit_imports")
	_ = os.Remove(initDiscoveryFile)

	return nil
}

func runStep2GenerateBaseline(dir string) error {
	log.Println("==> Step 2: Generating baseline configuration...")

	tfCmd := getTofuOrTerraform()

	baselineFile := filepath.Join(dir, "baseline.tf")
	_ = runCommand(dir, tfCmd, "plan", "-generate-config-out=baseline.tf")

	content, err := os.ReadFile(baselineFile)
	if err != nil {
		return fmt.Errorf("failed to read generated baseline.tf: %w", err)
	}

	cleaned := cleanProviderAttribute(string(content))
	// #nosec G304,G703
	if err := os.WriteFile(filepath.Clean(baselineFile), []byte(cleaned), 0600); err != nil {
		return fmt.Errorf("failed to write cleaned baseline.tf: %w", err)
	}

	if err := runCommand(dir, tfCmd, "apply", "-auto-approve"); err != nil {
		log.Printf("Notice during baseline apply: %v", err)
	}

	return nil
}

func runStep3RefactorAndCleanup(dir string) error {
	log.Println("==> Step 3: Refactoring components into segmented TF files and cleaning up...")

	baselineFile := filepath.Join(dir, "baseline.tf")
	// #nosec G304
	contentBytes, err := os.ReadFile(filepath.Clean(baselineFile))
	if err != nil {
		log.Printf("Notice: baseline.tf not found or unreadable, skipping split: %v", err)
		return nil
	}

	content := string(contentBytes)

	siteID := extractSiteID(content)
	if siteID != "" {
		log.Printf("Detected site ID: %s", siteID)
		content = strings.ReplaceAll(content, fmt.Sprintf(`"%s"`, siteID), "local.site_id")
	}

	resourcePattern := regexp.MustCompile(`(?s)(# __generated__ by OpenTofu from [^\n]+\nresource "([^"]+)" "[^"]+" \{.*?\n\})`)
	matches := resourcePattern.FindAllStringSubmatch(content, -1)

	segmented := make(map[string][]string)
	for _, match := range matches {
		fullBlock := match[1]
		rType := match[2]
		targetFile, ok := fileMapping[rType]
		if !ok {
			targetFile = "resources.tf"
		}
		segmented[targetFile] = append(segmented[targetFile], fullBlock)
	}

	for fileName, blocks := range segmented {
		targetPath := filepath.Join(dir, fileName)
		header := fmt.Sprintf("# Managed UniFi resources for %s\n\n", fileName)
		fileContent := header + strings.Join(blocks, "\n\n") + "\n"
		if err := os.WriteFile(targetPath, []byte(fileContent), 0600); err != nil {
			return fmt.Errorf("failed to write %s: %w", fileName, err)
		}
		log.Printf("Segmented %d resources into %s", len(blocks), fileName)
	}

	_ = os.Remove(baselineFile)
	_ = os.Remove(filepath.Join(dir, "explicit_imports.tf"))

	tfCmd := getTofuOrTerraform()
	_ = runCommand(dir, tfCmd, "fmt")

	return nil
}

func cleanProviderAttribute(content string) string {
	re := regexp.MustCompile(`(?m)^\s*provider\s*=\s*.*$\n?`)
	return re.ReplaceAllString(content, "")
}

func extractSiteID(content string) string {
	re := regexp.MustCompile(`site_id\s*=\s*"([a-f0-9\-]{36})"`)
	match := re.FindStringSubmatch(content)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

func getTofuOrTerraform() string {
	if _, err := exec.LookPath("tofu"); err == nil {
		return "tofu"
	}
	return "terraform"
}
