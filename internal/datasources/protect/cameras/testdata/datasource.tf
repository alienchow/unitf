
provider "unifi" {
	host    = "https://127.0.0.1"
	api_key = "test-key"
}

data "unifi_cameras" "test" {
}