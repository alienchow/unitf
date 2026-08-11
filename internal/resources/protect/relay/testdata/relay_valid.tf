provider "unifi" {
	host    = "https://127.0.0.1"
	api_key = "test-key"
}

resource "unifi_protect_relay" "test" {
	id   = "relay-uuid"
	name = "Garage Door"
}
