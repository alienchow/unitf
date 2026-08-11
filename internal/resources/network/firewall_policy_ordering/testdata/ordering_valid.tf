provider "unifi" {
	host    = "https://127.0.0.1"
	api_key = "test-key"
}

resource "unifi_firewall_policy_ordering" "test" {
	site_id      = "default"
	from_zone_id = "zone-1"
	to_zone_id   = "zone-2"
	
	before_system_defined = ["policy-1"]
	after_system_defined  = ["policy-2", "policy-3"]
}
