provider "unifi" {
	host    = "https://127.0.0.1"
	api_key = "test-key"
}

resource "unifi_network" "test_invalid" {
	site_id = "default"
	name    = "Test Network"
	type    = "INVALID_TYPE"
}
