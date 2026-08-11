provider "unifi" {
	host    = "https://127.0.0.1"
	api_key = "test-key"
}

resource "unifi_protect_sensor" "test" {
	id         = "sensor-uuid"
	name       = "Server Room"
	alarm      = true
	temp_limit = 25
}
