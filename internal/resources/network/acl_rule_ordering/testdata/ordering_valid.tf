provider "unifi" {
	host    = "https://127.0.0.1"
	api_key = "test-key"
}

resource "unifi_acl_rule_ordering" "test" {
	site_id  = "default"
	rule_ids = ["rule-1", "rule-2", "rule-3"]
}
