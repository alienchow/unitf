provider "unifi" {
	host    = "https://127.0.0.1"
	api_key = "test-key"
}

resource "unifi_acl_rule" "test" {
	site_id    = "default"
	name       = "Test Rule"
	action     = "ACCEPT"
	ip_version = "IPV4"
	protocols  = ["TCP"]
	source = {
		network_id = "network-1"
	}
	destination = {
		network_id = "network-2"
		port       = "443"
	}
}
