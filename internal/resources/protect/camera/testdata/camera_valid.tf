provider "unifi" {
	host    = "https://127.0.0.1"
	api_key = "test-key"
}

resource "unifi_protect_camera" "test" {
	id                = "camera-uuid"
	name              = "Front Door"
	video_mode        = "highFps"
	record_everything = true

	osd_settings = {
		is_name_enabled = true
		is_date_enabled = true
	}

	smart_detect_settings = {
		object_types = ["person", "vehicle"]
	}
}
