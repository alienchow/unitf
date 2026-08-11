provider "unifi" {
	host    = "https://127.0.0.1"
	api_key = "test-key"
}

resource "unifi_traffic_matching_list" "test" {
	site_id   = "default"
	name      = "Test IP List"
	type      = "IPV4"
	addresses = ["10.0.0.1", "10.0.0.2"]
}
