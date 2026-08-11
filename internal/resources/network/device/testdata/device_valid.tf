provider "unifi" {
	host    = "https://127.0.0.1"
	api_key = "test-key"
}

resource "unifi_device" "test" {
	site_id = "default"
	mac     = "00:11:22:33:44:55"
	name    = "Core Switch"
	adopt   = true
}
