provider "unifi" {
	host    = "https://127.0.0.1"
	api_key = "test-key"
}

resource "unifi_protect_chime" "test" {
	id     = "chime-uuid"
	name   = "Hallway Chime"
	volume = 50
}
