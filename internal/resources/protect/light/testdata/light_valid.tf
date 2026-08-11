provider "unifi" {
	host    = "https://127.0.0.1"
	api_key = "test-key"
}

resource "unifi_protect_light" "test" {
	id                   = "light-uuid"
	name                 = "Porch Light"
	is_indicator_enabled = true
	led_level            = 5
}
