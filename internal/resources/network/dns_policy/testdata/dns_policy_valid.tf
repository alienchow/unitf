provider "unifi" {
	host    = "https://127.0.0.1"
	api_key = "test-key"
}

resource "unifi_dns_policy" "test" {
	site_id = "default"
	name    = "Test DNS Policy"
	type    = "A_RECORD"
	value   = "10.0.0.1"
	ttl     = 3600
}
