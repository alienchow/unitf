# OpenTofu Provider for UniFi (unitf)

![License](https://img.shields.io/badge/license-MIT-blue.svg)
![Go Version](https://img.shields.io/badge/Go-1.26+-00ADD8?logo=go)
![OpenTofu](https://img.shields.io/badge/OpenTofu-1.6+-E0C323?logo=opentofu)

A declarative, idempotent OpenTofu and Terraform provider designed for Ubiquiti's UniFi Network and UniFi Protect applications.

Built from the ground up using the modern HashiCorp `terraform-plugin-framework` (Protocol v6), `unitf` interacts directly with the local API proxies of UniFi OS Consoles (such as the Dream Machine Pro, Cloud Key Gen2, and UNVR). It bypasses legacy dependencies and strictly enforces a highly declarative, predictable workflow for managing your entire UniFi infrastructure as code.

## Key Capabilities

- **Modern Architecture**: Developed on Terraform Plugin Framework v6, completely omitting legacy SDKv2 patterns to ensure strict state consistency and robust error handling.
- **Unified Ecosystem**: Manage both UniFi Network (networks, firewall policies, Wi-Fi broadcasts, traffic rules) and UniFi Protect (cameras, sensors, lights, chimes, viewers) from a single provider.
- **Local Console Direct Access**: Communicates directly via the local `/proxy/network/integration/v1/` and `/proxy/protect/integration/v1/` APIs.

## Installation

Add the provider to your OpenTofu or Terraform configuration:

```hcl
terraform {
  required_providers {
    unifi = {
      source  = "alienchow/unitf"
      version = "~> 0.1.0"
    }
  }
}
```

Run `tofu init` or `terraform init` to download and install the provider.

## Configuration

The provider requires authentication and connection details for your local UniFi Console. You can configure these natively in the provider block or supply them securely via environment variables.

```hcl
provider "unifi" {
  host     = "https://192.168.1.1"       # Or via environment variable: UNIFI_HOST
  api_key  = "your-secret-api-key"       # Or via environment variable: UNIFI_API_KEY
  site_id  = "default"                   # Or via environment variable: UNIFI_SITE_ID
  insecure = true                        # Optional: Skip TLS verification for self-signed certificates
}
```

> **Note**: For production environments, we strongly recommend omitting the `api_key` from your configuration files and relying entirely on the `UNIFI_API_KEY` environment variable.

## Usage Example

The following example demonstrates how to bootstrap a corporate network alongside a UniFi Protect camera setup.

```hcl
# 1. Retrieve the default site identity
data "unifi_sites" "default" {}

locals {
  site_id = data.unifi_sites.default.sites[0].id
}

# 2. Provision a Corporate Network with a managed DHCP server
resource "unifi_network" "corporate" {
  site_id = local.site_id
  name    = "Corporate VLAN"

  gateway_managed = {
    vlan_id                 = 10
    purpose                 = "CORPORATE"
    internet_access_enabled = true

    ipv4 = {
      host_ip_address    = "10.0.10.1"
      prefix_length      = 24
      auto_scale_enabled = false

      dhcp_server = {
        ip_address_range   = { start = "10.0.10.100", stop = "10.0.10.200" }
        lease_time_seconds = 86400
      }
    }
  }
}

# 3. Lookup a connected UniFi Protect Camera by Name
data "unifi_cameras" "all" {}

locals {
  front_cam = [for c in data.unifi_cameras.all.items : c if c.name == "Front Door"][0]
}

# 4. Enforce Protect Recording and Smart Detection Behaviors
resource "unifi_protect_camera" "front_door" {
  id = local.front_cam.id

  name              = "Front Door"
  video_mode        = "highFps"
  record_everything = true

  osd_settings = {
    is_name_enabled = true
    is_date_enabled = true
  }

  smart_detect_settings = {
    object_types = ["person", "vehicle", "package"]
  }
}
```

## Importing Existing Resources

OpenTofu allows you to generate configuration code natively for resources you already have running on your console using `tofu plan -generate-config-out=baseline.tf`.

### UniFi Protect
You can dynamically discover and generate code for all your Protect devices (Cameras, Chimes, Lights, etc.) by generating explicit `import` blocks. Create a `local_file` resource to unroll your devices into static imports:

```hcl
data "unifi_cameras" "all" {}

resource "local_file" "explicit_imports" {
  filename = "explicit_imports.tf"
  content  = join("\n", [
    for cam in data.unifi_cameras.all.items :
    "import {\n  id = \"${cam.id}\"\n  to = unifi_protect_camera.cam_${replace(cam.id, \"-\", \"_\")}\n}"
  ])
}
```
Run `tofu apply -target=local_file.explicit_imports` to create the file, then `tofu plan -generate-config-out=baseline.tf` to generate your infrastructure code!

### UniFi Network
Due to limitations in the UniFi proxy API, Network resources (like VLANs and Firewall Policies) **do not have list endpoints**. This means you cannot automatically discover them via `data` sources. To import existing networks, you must manually grab the UUID from your UniFi dashboard URL and write explicit import blocks:

```hcl
import {
  id = "your-network-uuid-here"
  to = unifi_network.my_network
}
```

## Contributing

We welcome community contributions. To maintain the integrity and consistency of the codebase, please review the [Agentic Guidelines (`AGENTS.md`)](AGENTS.md) before submitting pull requests. The guidelines enforce our architectural standards, including idiomatic Go practices, strictly isolated Domain Models, Dependency Injection, and the usage of the `testdata/` pattern for robust unit testing.

## License

This project is distributed under the MIT License.
