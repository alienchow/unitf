provider "unifi" {
	host    = "https://127.0.0.1"
	api_key = "test-key"
}

resource "unifi_hotspot_voucher" "test" {
	site_id  = "default"
	quota    = 1
	duration = 1440
}
