provider "unifi" {
	host    = "https://127.0.0.1"
	api_key = "test-key"
}

resource "unifi_protect_siren" "test" {
	id     = "siren-uuid"
	name   = "Outdoor Siren"
	volume = 100
}
