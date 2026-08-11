provider "unifi" {
	host    = "https://127.0.0.1"
	api_key = "test-key"
}

resource "unifi_wifi_broadcast" "test" {
	site_id    = "default"
	name       = "Corporate WiFi"
	ssid       = "Corporate"
	network_id = "network-123"
	security   = "WPA3"
	passphrase = "supersecretpassword"
	mode       = "STANDARD"
}
