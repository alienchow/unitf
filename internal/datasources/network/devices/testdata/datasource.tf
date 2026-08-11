
provider "unifi" {
	host    = "https://127.0.0.1"
	api_key = "test-key"
}

data "unifi_network_devices" "test" {
	site_id = "default"
}