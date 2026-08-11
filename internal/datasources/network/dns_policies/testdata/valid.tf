provider "unifi" {
	host    = "https://127.0.0.1"
	api_key = "test-key"
}

data "unifi_network_dns_policies" "all" {
	site_id = "default"
}
