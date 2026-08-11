provider "unifi" {
	host    = "https://127.0.0.1"
	api_key = "test-key"
}

resource "unifi_firewall_zone" "test" {
	site_id = "default"
	name    = "Test Zone"
	network_ids = ["network-123"]
}
