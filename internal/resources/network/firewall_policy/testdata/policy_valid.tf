provider "unifi" {
	host    = "https://127.0.0.1"
	api_key = "test-key"
}

resource "unifi_firewall_policy" "test" {
	site_id = "default"
	name    = "Test Policy"
	enabled = true
	logging = true
	
	action = {
		accept = {}
	}
	
	ip_protocol_scope = {
		ip_version = "IPV4"
		protocols  = ["TCP", "UDP"]
	}

	source = {
		zone_id = "zone-1"
	}
	
	destination = {
		zone_id = "zone-2"
	}
}
