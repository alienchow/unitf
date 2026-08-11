provider "unifi" {
	host    = "https://127.0.0.1"
	api_key = "test-key"
}

resource "unifi_network" "test" {
	site_id = "default"
	name    = "Test Network"
	type    = "GATEWAY_MANAGED"

	gateway_managed = {
		vlan_id = 10
		purpose = "CORPORATE"
		ipv4 = {
			enabled = true
			subnet_mask = "255.255.255.0"
			dhcp_server = {
				range_start = "10.0.10.10"
				range_stop  = "10.0.10.100"
			}
		}
	}
}
