
provider "unifi" {
	host    = "https://127.0.0.1"
	api_key = "test-key"
}

data "unifi_mc_lag_domains" "test" {
	site_id = "default"
}