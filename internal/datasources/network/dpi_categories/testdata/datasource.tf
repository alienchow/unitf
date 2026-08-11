
provider "unifi" {
	host    = "https://127.0.0.1"
	api_key = "test-key"
}

data "unifi_network_dpi_categories" "test" {
	site_id = "default"
}