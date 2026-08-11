# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/2adebba0-273f-424b-8780-5e89d29b2db1"
resource "unifi_network" "net_2adebba0_273f_424b_8780_5e89d29b2db1" {
  provider = unifi
  name     = "IoT"
  site_id  = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/f04c183c-422b-4f01-a927-d0c4b09456ac"
resource "unifi_firewall_zone" "zone_f04c183c_422b_4f01_a927_d0c4b09456ac" {
  provider = unitf
  name     = "Internal"
  site_id  = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
}

# __generated__ by OpenTofu from "64d7073b02e86c03e70003f4"
resource "unifi_protect_camera" "cam_64d7073b02e86c03e70003f4" {
  name = "Front Door"
}
