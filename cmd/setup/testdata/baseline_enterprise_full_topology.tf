# Enterprise Full Topology Baseline Test Data - Generated for Exhaustive Testing


# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/6cb78a5b-419e-4eda-b257-5effa4d63f78"
resource "unifi_network" "net_6cb78a5b_419e_4eda_b257_5effa4d63f78" {
  provider = unifi
  gateway_managed = {
    dhcp_guarding = false
    ipv4          = "10.1.0.1/24"
    multicast_dns = false
    purpose       = "CORPORATE"
    vlan_id       = 101
  }
  name    = "Enterprise-VLAN-101-Segment-1"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  type    = "GATEWAY_MANAGED"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/88e3a007-dc9b-4c41-be6d-7749300f2cca"
resource "unifi_network" "net_88e3a007_dc9b_4c41_be6d_7749300f2cca" {
  provider = unifi
  gateway_managed = {
    dhcp_guarding = true
    ipv4          = "10.2.0.1/24"
    multicast_dns = false
    purpose       = "CORPORATE"
    vlan_id       = 102
  }
  name    = "Enterprise-VLAN-102-Segment-2"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  type    = "GATEWAY_MANAGED"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/acdd3d20-b869-41d9-b3d1-ab3ab13c25e5"
resource "unifi_network" "net_acdd3d20_b869_41d9_b3d1_ab3ab13c25e5" {
  provider = unifi
  gateway_managed = {
    dhcp_guarding = false
    ipv4          = "10.3.0.1/24"
    multicast_dns = true
    purpose       = "CORPORATE"
    vlan_id       = 103
  }
  name    = "Enterprise-VLAN-103-Segment-3"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  type    = "GATEWAY_MANAGED"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/269a5d50-5633-409b-96e2-2f4e64bacbdb"
resource "unifi_network" "net_269a5d50_5633_409b_96e2_2f4e64bacbdb" {
  provider = unifi
  gateway_managed = {
    dhcp_guarding = true
    ipv4          = "10.4.0.1/24"
    multicast_dns = false
    purpose       = "CORPORATE"
    vlan_id       = 104
  }
  name    = "Enterprise-VLAN-104-Segment-4"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  type    = "GATEWAY_MANAGED"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/e29f54c4-db44-414d-813b-0d2b5cbb4072"
resource "unifi_network" "net_e29f54c4_db44_414d_813b_0d2b5cbb4072" {
  provider = unifi
  gateway_managed = {
    dhcp_guarding = false
    ipv4          = "10.5.0.1/24"
    multicast_dns = false
    purpose       = "CORPORATE"
    vlan_id       = 105
  }
  name    = "Enterprise-VLAN-105-Segment-5"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  type    = "GATEWAY_MANAGED"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/1a3e6532-9f7a-421d-81b5-159d18266d25"
resource "unifi_network" "net_1a3e6532_9f7a_421d_81b5_159d18266d25" {
  provider = unifi
  gateway_managed = {
    dhcp_guarding = true
    ipv4          = "10.6.0.1/24"
    multicast_dns = true
    purpose       = "CORPORATE"
    vlan_id       = 106
  }
  name    = "Enterprise-VLAN-106-Segment-6"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  type    = "GATEWAY_MANAGED"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/c64aeb21-a45c-46dd-ab2d-c74582abaea6"
resource "unifi_network" "net_c64aeb21_a45c_46dd_ab2d_c74582abaea6" {
  provider = unifi
  gateway_managed = {
    dhcp_guarding = false
    ipv4          = "10.7.0.1/24"
    multicast_dns = false
    purpose       = "CORPORATE"
    vlan_id       = 107
  }
  name    = "Enterprise-VLAN-107-Segment-7"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  type    = "GATEWAY_MANAGED"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/400ba948-0bc6-4f4d-ba16-dead77893d2c"
resource "unifi_network" "net_400ba948_0bc6_4f4d_ba16_dead77893d2c" {
  provider = unifi
  gateway_managed = {
    dhcp_guarding = true
    ipv4          = "10.8.0.1/24"
    multicast_dns = false
    purpose       = "CORPORATE"
    vlan_id       = 108
  }
  name    = "Enterprise-VLAN-108-Segment-8"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  type    = "GATEWAY_MANAGED"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/e95c3407-0764-40d0-9171-029bb59c1734"
resource "unifi_network" "net_e95c3407_0764_40d0_9171_029bb59c1734" {
  provider = unifi
  gateway_managed = {
    dhcp_guarding = false
    ipv4          = "10.9.0.1/24"
    multicast_dns = true
    purpose       = "CORPORATE"
    vlan_id       = 109
  }
  name    = "Enterprise-VLAN-109-Segment-9"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  type    = "GATEWAY_MANAGED"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/12269655-b4f4-4926-a8f7-0b311031fe95"
resource "unifi_network" "net_12269655_b4f4_4926_a8f7_0b311031fe95" {
  provider = unifi
  gateway_managed = {
    dhcp_guarding = true
    ipv4          = "10.10.0.1/24"
    multicast_dns = false
    purpose       = "CORPORATE"
    vlan_id       = 110
  }
  name    = "Enterprise-VLAN-110-Segment-10"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  type    = "GATEWAY_MANAGED"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/55802a6f-645f-4167-8a1d-ca36c7e57852"
resource "unifi_network" "net_55802a6f_645f_4167_8a1d_ca36c7e57852" {
  provider = unifi
  gateway_managed = {
    dhcp_guarding = false
    ipv4          = "10.11.0.1/24"
    multicast_dns = false
    purpose       = "CORPORATE"
    vlan_id       = 111
  }
  name    = "Enterprise-VLAN-111-Segment-11"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  type    = "GATEWAY_MANAGED"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/9987ec14-efd9-4404-8064-f0f10a679c40"
resource "unifi_network" "net_9987ec14_efd9_4404_8064_f0f10a679c40" {
  provider = unifi
  gateway_managed = {
    dhcp_guarding = true
    ipv4          = "10.12.0.1/24"
    multicast_dns = true
    purpose       = "CORPORATE"
    vlan_id       = 112
  }
  name    = "Enterprise-VLAN-112-Segment-12"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  type    = "GATEWAY_MANAGED"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/a8693827-1932-4b27-babc-6e279de84136"
resource "unifi_network" "net_a8693827_1932_4b27_babc_6e279de84136" {
  provider = unifi
  gateway_managed = {
    dhcp_guarding = false
    ipv4          = "10.13.0.1/24"
    multicast_dns = false
    purpose       = "CORPORATE"
    vlan_id       = 113
  }
  name    = "Enterprise-VLAN-113-Segment-13"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  type    = "GATEWAY_MANAGED"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/8355ec54-0bc5-461e-bb38-369c105a122d"
resource "unifi_network" "net_8355ec54_0bc5_461e_bb38_369c105a122d" {
  provider = unifi
  gateway_managed = {
    dhcp_guarding = true
    ipv4          = "10.14.0.1/24"
    multicast_dns = false
    purpose       = "CORPORATE"
    vlan_id       = 114
  }
  name    = "Enterprise-VLAN-114-Segment-14"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  type    = "GATEWAY_MANAGED"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/d2fca4c8-f167-495f-aca6-1a1bb3fa7cc6"
resource "unifi_network" "net_d2fca4c8_f167_495f_aca6_1a1bb3fa7cc6" {
  provider = unifi
  gateway_managed = {
    dhcp_guarding = false
    ipv4          = "10.15.0.1/24"
    multicast_dns = true
    purpose       = "CORPORATE"
    vlan_id       = 115
  }
  name    = "Enterprise-VLAN-115-Segment-15"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  type    = "GATEWAY_MANAGED"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/2bb10df2-ca4b-41fc-8eeb-07e1c3be2da7"
resource "unifi_network" "net_2bb10df2_ca4b_41fc_8eeb_07e1c3be2da7" {
  provider = unifi
  gateway_managed = {
    dhcp_guarding = true
    ipv4          = "10.16.0.1/24"
    multicast_dns = false
    purpose       = "CORPORATE"
    vlan_id       = 116
  }
  name    = "Enterprise-VLAN-116-Segment-16"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  type    = "GATEWAY_MANAGED"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/22d87da1-5eb8-4cfa-b879-41b7a78a088a"
resource "unifi_network" "net_22d87da1_5eb8_4cfa_b879_41b7a78a088a" {
  provider = unifi
  gateway_managed = {
    dhcp_guarding = false
    ipv4          = "10.17.0.1/24"
    multicast_dns = false
    purpose       = "CORPORATE"
    vlan_id       = 117
  }
  name    = "Enterprise-VLAN-117-Segment-17"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  type    = "GATEWAY_MANAGED"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/d8d992a4-26d2-421c-bcf3-6aa74e33c291"
resource "unifi_network" "net_d8d992a4_26d2_421c_bcf3_6aa74e33c291" {
  provider = unifi
  gateway_managed = {
    dhcp_guarding = true
    ipv4          = "10.18.0.1/24"
    multicast_dns = true
    purpose       = "CORPORATE"
    vlan_id       = 118
  }
  name    = "Enterprise-VLAN-118-Segment-18"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  type    = "GATEWAY_MANAGED"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/498a7ff9-0fb7-45e4-87e6-d1a2eb8e9162"
resource "unifi_network" "net_498a7ff9_0fb7_45e4_87e6_d1a2eb8e9162" {
  provider = unifi
  gateway_managed = {
    dhcp_guarding = false
    ipv4          = "10.19.0.1/24"
    multicast_dns = false
    purpose       = "CORPORATE"
    vlan_id       = 119
  }
  name    = "Enterprise-VLAN-119-Segment-19"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  type    = "GATEWAY_MANAGED"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/fc1cc9b8-1fce-4a7c-b095-85510770b6c8"
resource "unifi_network" "net_fc1cc9b8_1fce_4a7c_b095_85510770b6c8" {
  provider = unifi
  gateway_managed = {
    dhcp_guarding = true
    ipv4          = "10.20.0.1/24"
    multicast_dns = false
    purpose       = "CORPORATE"
    vlan_id       = 120
  }
  name    = "Enterprise-VLAN-120-Segment-20"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  type    = "GATEWAY_MANAGED"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/2d9c0ca3-752e-43cc-856e-5417ce5fe2ad"
resource "unifi_firewall_zone" "zone_2d9c0ca3_752e_43cc_856e_5417ce5fe2ad" {
  provider = unitf
  name        = "Internal"
  network_ids = ["5e941f4c-5db0-43fe-afbb-b61f934f4612", "6ab3536a-efd6-44a3-ae70-8dfc1848e5f0"]
  site_id     = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/8f0d0761-01b3-402d-b52c-5ebae07a5aae"
resource "unifi_firewall_zone" "zone_8f0d0761_01b3_402d_b52c_5ebae07a5aae" {
  provider = unitf
  name        = "DMZ"
  network_ids = ["b7b5af23-6e83-4387-bd93-3ba818630047", "7f4e27d2-1396-4549-b970-d4990de9452d"]
  site_id     = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/0de2115e-6b08-45d6-9362-64a8b4b893a5"
resource "unifi_firewall_zone" "zone_0de2115e_6b08_45d6_9362_64a8b4b893a5" {
  provider = unitf
  name        = "IoT-Zone"
  network_ids = ["576f87c3-38b6-407d-8c69-ce83b9c14f1f", "0ae3691e-912e-42c5-bcd3-c3d7c8579d74"]
  site_id     = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/01b32de0-c183-40a5-b966-e942407a8fa6"
resource "unifi_firewall_zone" "zone_01b32de0_c183_40a5_b966_e942407a8fa6" {
  provider = unitf
  name        = "Guest-Zone"
  network_ids = ["989d7700-6423-4f8e-8cac-25817cbdd11e", "94a9eace-94ef-4a19-9d5e-7dea54e2928f"]
  site_id     = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/506b6ed4-9a96-4448-a6d3-82d59889783c"
resource "unifi_firewall_zone" "zone_506b6ed4_9a96_4448_a6d3_82d59889783c" {
  provider = unitf
  name        = "VoIP-Zone"
  network_ids = ["5f013f9e-3afd-44ce-81fc-cbe91c6587ca", "59db4763-6e00-46e8-b93b-189d0ac8460a"]
  site_id     = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/82194965-4dd8-4ddb-ae58-a8adb17973b9"
resource "unifi_firewall_zone" "zone_82194965_4dd8_4ddb_ae58_a8adb17973b9" {
  provider = unitf
  name        = "Management-Zone"
  network_ids = ["fa18b634-ec71-4308-b33a-73af918fe6af", "670a589e-3445-4a84-80b5-ee61be8cb437"]
  site_id     = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/589c4df2-026c-49f6-8937-3b3f8d55f056"
resource "unifi_firewall_zone" "zone_589c4df2_026c_49f6_8937_3b3f8d55f056" {
  provider = unitf
  name        = "Storage-SAN"
  network_ids = ["2109095f-4a1a-4b3d-8afc-9bb68f4d88f9", "bec51610-35f0-427e-aa8e-9737cb650ca5"]
  site_id     = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/78ddf0cf-bb6d-49a1-abc9-2a8fa40f2c59"
resource "unifi_firewall_zone" "zone_78ddf0cf_bb6d_49a1_abc9_2a8fa40f2c59" {
  provider = unitf
  name        = "PCI-DSS-Zone"
  network_ids = ["250a8a85-9ba9-40a6-a02a-d3902ca5f50b", "a431ffd6-0221-425a-89eb-4fde3505258a"]
  site_id     = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/6badfcbd-838f-4f06-a3d9-4e31ea35c7b1"
resource "unifi_firewall_zone" "zone_6badfcbd_838f_4f06_a3d9_4e31ea35c7b1" {
  provider = unitf
  name        = "Development-Zone"
  network_ids = ["d63db4a4-73f9-4fe8-90d9-df21b8d73cd3", "cef5e5eb-806b-49c8-89c2-2f9cb5d25636"]
  site_id     = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/ff684e73-7286-4a67-86c7-37597bf80e1d"
resource "unifi_firewall_zone" "zone_ff684e73_7286_4a67_86c7_37597bf80e1d" {
  provider = unitf
  name        = "Staging-Zone"
  network_ids = ["148018a8-dc27-43b3-82c2-3f13e9d2cfce", "5d1f1f13-2194-4a75-897f-91e18627a365"]
  site_id     = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/0cb1afc4-a9d9-44de-811a-125a9bc8769f"
resource "unifi_firewall_zone" "zone_0cb1afc4_a9d9_44de_811a_125a9bc8769f" {
  provider = unitf
  name        = "Production-Zone"
  network_ids = ["96e3c8dd-1644-4c73-9bbc-20e96c5a0333", "c958d7b4-4db1-436d-8045-2c00ddc1ca6e"]
  site_id     = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/bb85076d-75d6-476f-8bc7-73c0fc8d82c7"
resource "unifi_firewall_zone" "zone_bb85076d_75d6_476f_8bc7_73c0fc8d82c7" {
  provider = unitf
  name        = "Security-Cameras"
  network_ids = ["056f28d7-b840-4170-8f6e-9ea7738c6481", "97e6ea10-1e4b-4530-ae0e-ce89e9ba3a92"]
  site_id     = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/1c4b5de9-ac11-4d0d-abcd-c72bf88e32ca"
resource "unifi_firewall_zone" "zone_1c4b5de9_ac11_4d0d_abcd_c72bf88e32ca" {
  provider = unitf
  name        = "Access-Control"
  network_ids = ["a7ff62b4-2e2f-4e86-93c1-6d4712df73dc", "32be2f99-ca93-4135-8202-a3d6c3df7464"]
  site_id     = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/635eefac-0d9a-47a1-9d3f-2693270898f6"
resource "unifi_firewall_zone" "zone_635eefac_0d9a_47a1_9d3f_2693270898f6" {
  provider = unitf
  name        = "Backup-Vault"
  network_ids = ["d1038b8c-c669-4bb0-88cc-22ad48ff546f", "7374262a-af7e-4ae2-bb12-be4658727db1"]
  site_id     = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/9d9981ad-5bde-4b25-a881-fe2168d481bc"
resource "unifi_firewall_zone" "zone_9d9981ad_5bde_4b25_a881_fe2168d481bc" {
  provider = unitf
  name        = "Quarantine"
  network_ids = ["798a3f33-5578-442b-b922-e128e1775709", "ea54b66a-cedd-477b-a4be-77f39cc1f41c"]
  site_id     = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/f8b87551-7207-4dfc-bd0d-22eee22363e1"
resource "unifi_firewall_policy" "fw_f8b87551_7207_4dfc_bd0d_22eee22363e1" {
  action      = "DROP"
  enabled     = true
  name        = "Rule-001-Filter-DROP-Traffic-1"
  site_id     = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  destination = {
    port = "81"
  }
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/9b188f9d-db05-4529-86b0-5125f71bbce0"
resource "unifi_firewall_policy" "fw_9b188f9d_db05_4529_86b0_5125f71bbce0" {
  action      = "ACCEPT"
  enabled     = true
  name        = "Rule-002-Filter-ACCEPT-Traffic-2"
  site_id     = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  destination = {
    port = "82"
  }
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/c600b11e-9d74-4efc-aee6-94e901215f61"
resource "unifi_firewall_policy" "fw_c600b11e_9d74_4efc_aee6_94e901215f61" {
  action      = "DROP"
  enabled     = true
  name        = "Rule-003-Filter-DROP-Traffic-3"
  site_id     = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  destination = {
    port = "83"
  }
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/fe805f37-9496-4c77-9a7c-5ca96c10cfb0"
resource "unifi_firewall_policy" "fw_fe805f37_9496_4c77_9a7c_5ca96c10cfb0" {
  action      = "ACCEPT"
  enabled     = false
  name        = "Rule-004-Filter-ACCEPT-Traffic-4"
  site_id     = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  destination = {
    port = "84"
  }
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/2c643aa0-b4ae-4e52-9b1d-264ce87d6940"
resource "unifi_firewall_policy" "fw_2c643aa0_b4ae_4e52_9b1d_264ce87d6940" {
  action      = "DROP"
  enabled     = true
  name        = "Rule-005-Filter-DROP-Traffic-5"
  site_id     = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  destination = {
    port = "85"
  }
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/c8259a4c-ad41-4483-84d8-214b4b941bb7"
resource "unifi_firewall_policy" "fw_c8259a4c_ad41_4483_84d8_214b4b941bb7" {
  action      = "ACCEPT"
  enabled     = true
  name        = "Rule-006-Filter-ACCEPT-Traffic-6"
  site_id     = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  destination = {
    port = "86"
  }
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/e4616b63-5e85-4ee3-acd0-ed7eb3a0816e"
resource "unifi_firewall_policy" "fw_e4616b63_5e85_4ee3_acd0_ed7eb3a0816e" {
  action      = "DROP"
  enabled     = true
  name        = "Rule-007-Filter-DROP-Traffic-7"
  site_id     = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  destination = {
    port = "87"
  }
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/e9c3796c-fc91-486b-ae36-ff3ca2d6679c"
resource "unifi_firewall_policy" "fw_e9c3796c_fc91_486b_ae36_ff3ca2d6679c" {
  action      = "ACCEPT"
  enabled     = false
  name        = "Rule-008-Filter-ACCEPT-Traffic-8"
  site_id     = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  destination = {
    port = "88"
  }
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/d1504a57-36aa-4aad-887f-29c5f85debd4"
resource "unifi_firewall_policy" "fw_d1504a57_36aa_4aad_887f_29c5f85debd4" {
  action      = "DROP"
  enabled     = true
  name        = "Rule-009-Filter-DROP-Traffic-9"
  site_id     = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  destination = {
    port = "89"
  }
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/af62636e-f228-4c20-aa0f-0bd95b1c5c92"
resource "unifi_firewall_policy" "fw_af62636e_f228_4c20_aa0f_0bd95b1c5c92" {
  action      = "ACCEPT"
  enabled     = true
  name        = "Rule-010-Filter-ACCEPT-Traffic-10"
  site_id     = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  destination = {
    port = "90"
  }
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/c3a049c0-bbb2-4dbe-a75d-e84b14ef6311"
resource "unifi_firewall_policy" "fw_c3a049c0_bbb2_4dbe_a75d_e84b14ef6311" {
  action      = "DROP"
  enabled     = true
  name        = "Rule-011-Filter-DROP-Traffic-11"
  site_id     = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  destination = {
    port = "91"
  }
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/f021feef-41f1-4ada-8a2f-79528690f173"
resource "unifi_firewall_policy" "fw_f021feef_41f1_4ada_8a2f_79528690f173" {
  action      = "ACCEPT"
  enabled     = false
  name        = "Rule-012-Filter-ACCEPT-Traffic-12"
  site_id     = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  destination = {
    port = "92"
  }
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/19988390-9f7a-4544-a813-06212aec74ff"
resource "unifi_firewall_policy" "fw_19988390_9f7a_4544_a813_06212aec74ff" {
  action      = "DROP"
  enabled     = true
  name        = "Rule-013-Filter-DROP-Traffic-13"
  site_id     = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  destination = {
    port = "93"
  }
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/0743aee3-7b2a-475e-9abf-bb61b1389803"
resource "unifi_firewall_policy" "fw_0743aee3_7b2a_475e_9abf_bb61b1389803" {
  action      = "ACCEPT"
  enabled     = true
  name        = "Rule-014-Filter-ACCEPT-Traffic-14"
  site_id     = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  destination = {
    port = "94"
  }
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/b0adad24-0831-4ad0-ac1a-e1836a70bb53"
resource "unifi_firewall_policy" "fw_b0adad24_0831_4ad0_ac1a_e1836a70bb53" {
  action      = "DROP"
  enabled     = true
  name        = "Rule-015-Filter-DROP-Traffic-15"
  site_id     = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  destination = {
    port = "95"
  }
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/fd725a11-b76d-4c42-9e0a-117702077e86"
resource "unifi_firewall_policy" "fw_fd725a11_b76d_4c42_9e0a_117702077e86" {
  action      = "ACCEPT"
  enabled     = false
  name        = "Rule-016-Filter-ACCEPT-Traffic-16"
  site_id     = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  destination = {
    port = "96"
  }
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/0af3b16b-57a4-4abd-92af-d276d9366f02"
resource "unifi_firewall_policy" "fw_0af3b16b_57a4_4abd_92af_d276d9366f02" {
  action      = "DROP"
  enabled     = true
  name        = "Rule-017-Filter-DROP-Traffic-17"
  site_id     = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  destination = {
    port = "97"
  }
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/b55d4507-52ff-4b7b-833d-410446d78f87"
resource "unifi_firewall_policy" "fw_b55d4507_52ff_4b7b_833d_410446d78f87" {
  action      = "ACCEPT"
  enabled     = true
  name        = "Rule-018-Filter-ACCEPT-Traffic-18"
  site_id     = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  destination = {
    port = "98"
  }
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/748f0c92-bed4-45f6-a69e-08c2b40950e3"
resource "unifi_firewall_policy" "fw_748f0c92_bed4_45f6_a69e_08c2b40950e3" {
  action      = "DROP"
  enabled     = true
  name        = "Rule-019-Filter-DROP-Traffic-19"
  site_id     = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  destination = {
    port = "99"
  }
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/9dc2aa35-07df-4b45-bb25-86fc7c76bda2"
resource "unifi_firewall_policy" "fw_9dc2aa35_07df_4b45_bb25_86fc7c76bda2" {
  action      = "ACCEPT"
  enabled     = false
  name        = "Rule-020-Filter-ACCEPT-Traffic-20"
  site_id     = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  destination = {
    port = "100"
  }
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/3b635e76-ebbb-4aab-a22b-e2cda0ce6561"
resource "unifi_firewall_policy" "fw_3b635e76_ebbb_4aab_a22b_e2cda0ce6561" {
  action      = "DROP"
  enabled     = true
  name        = "Rule-021-Filter-DROP-Traffic-21"
  site_id     = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  destination = {
    port = "101"
  }
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/d8c7fd9b-2d19-4649-980d-fa22bc6f1792"
resource "unifi_firewall_policy" "fw_d8c7fd9b_2d19_4649_980d_fa22bc6f1792" {
  action      = "ACCEPT"
  enabled     = true
  name        = "Rule-022-Filter-ACCEPT-Traffic-22"
  site_id     = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  destination = {
    port = "102"
  }
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/73e4e5ab-7a9f-4f0d-8fc6-3e4667135807"
resource "unifi_firewall_policy" "fw_73e4e5ab_7a9f_4f0d_8fc6_3e4667135807" {
  action      = "DROP"
  enabled     = true
  name        = "Rule-023-Filter-DROP-Traffic-23"
  site_id     = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  destination = {
    port = "103"
  }
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/0d4beb23-aafc-47dd-994e-5796ced2d0f3"
resource "unifi_firewall_policy" "fw_0d4beb23_aafc_47dd_994e_5796ced2d0f3" {
  action      = "ACCEPT"
  enabled     = false
  name        = "Rule-024-Filter-ACCEPT-Traffic-24"
  site_id     = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  destination = {
    port = "104"
  }
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/dd81cc8e-f312-4a20-9fec-9c33ae12f746"
resource "unifi_firewall_policy" "fw_dd81cc8e_f312_4a20_9fec_9c33ae12f746" {
  action      = "DROP"
  enabled     = true
  name        = "Rule-025-Filter-DROP-Traffic-25"
  site_id     = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  destination = {
    port = "105"
  }
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/662f4153-2926-4ba2-8aed-fc8c7c9949cd"
resource "unifi_firewall_policy" "fw_662f4153_2926_4ba2_8aed_fc8c7c9949cd" {
  action      = "ACCEPT"
  enabled     = true
  name        = "Rule-026-Filter-ACCEPT-Traffic-26"
  site_id     = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  destination = {
    port = "106"
  }
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/3e413a41-953d-4db9-abe9-1351de384c82"
resource "unifi_firewall_policy" "fw_3e413a41_953d_4db9_abe9_1351de384c82" {
  action      = "DROP"
  enabled     = true
  name        = "Rule-027-Filter-DROP-Traffic-27"
  site_id     = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  destination = {
    port = "107"
  }
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/8d7630a4-51f3-4e8d-9838-769dd024cf4f"
resource "unifi_firewall_policy" "fw_8d7630a4_51f3_4e8d_9838_769dd024cf4f" {
  action      = "ACCEPT"
  enabled     = false
  name        = "Rule-028-Filter-ACCEPT-Traffic-28"
  site_id     = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  destination = {
    port = "108"
  }
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/ca033755-e05a-4fa2-b83f-95e4f787a343"
resource "unifi_firewall_policy" "fw_ca033755_e05a_4fa2_b83f_95e4f787a343" {
  action      = "DROP"
  enabled     = true
  name        = "Rule-029-Filter-DROP-Traffic-29"
  site_id     = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  destination = {
    port = "109"
  }
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/ac16b47f-acc1-430d-9bbd-f1b1483e9a96"
resource "unifi_firewall_policy" "fw_ac16b47f_acc1_430d_9bbd_f1b1483e9a96" {
  action      = "ACCEPT"
  enabled     = true
  name        = "Rule-030-Filter-ACCEPT-Traffic-30"
  site_id     = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  destination = {
    port = "110"
  }
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/b2b98fe7-67f8-4ac8-9d57-1d8e2596305b"
resource "unifi_dns_policy" "dns_b2b98fe7_67f8_4ac8_9d57_1d8e2596305b" {
  provider = unifi
  enabled = true
  name    = "node-1.corp.internal"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  ttl     = 600
  type    = "AAAA_RECORD"
  value   = "10.0.1.50"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/b18c2c49-1a50-444e-9915-6aa4ba533da9"
resource "unifi_dns_policy" "dns_b18c2c49_1a50_444e_9915_6aa4ba533da9" {
  provider = unifi
  enabled = true
  name    = "node-2.corp.internal"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  ttl     = 900
  type    = "CNAME_RECORD"
  value   = "10.0.2.50"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/b0f0cce0-8ff8-48d5-be3e-6b6e28667955"
resource "unifi_dns_policy" "dns_b0f0cce0_8ff8_48d5_be3e_6b6e28667955" {
  provider = unifi
  enabled = true
  name    = "node-3.corp.internal"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  ttl     = 1200
  type    = "TXT_RECORD"
  value   = "10.0.3.50"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/54c83772-3237-4ba5-a1c8-62e946126947"
resource "unifi_dns_policy" "dns_54c83772_3237_4ba5_a1c8_62e946126947" {
  provider = unifi
  enabled = true
  name    = "node-4.corp.internal"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  ttl     = 1500
  type    = "FORWARD_DOMAIN"
  value   = "10.0.4.50"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/cecd16b8-e579-4f44-a1c0-60e8e6f83300"
resource "unifi_dns_policy" "dns_cecd16b8_e579_4f44_a1c0_60e8e6f83300" {
  provider = unifi
  enabled = true
  name    = "node-5.corp.internal"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  ttl     = 300
  type    = "A_RECORD"
  value   = "10.0.5.50"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/cc3a6255-16f1-4136-88ad-5ffebcff26bd"
resource "unifi_dns_policy" "dns_cc3a6255_16f1_4136_88ad_5ffebcff26bd" {
  provider = unifi
  enabled = true
  name    = "node-6.corp.internal"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  ttl     = 600
  type    = "AAAA_RECORD"
  value   = "10.0.6.50"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/bb504e8b-ae80-4da3-a0a3-818db284b95e"
resource "unifi_dns_policy" "dns_bb504e8b_ae80_4da3_a0a3_818db284b95e" {
  provider = unifi
  enabled = true
  name    = "node-7.corp.internal"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  ttl     = 900
  type    = "CNAME_RECORD"
  value   = "10.0.7.50"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/d84d3dcc-a493-4766-bbc2-706c47178a74"
resource "unifi_dns_policy" "dns_d84d3dcc_a493_4766_bbc2_706c47178a74" {
  provider = unifi
  enabled = true
  name    = "node-8.corp.internal"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  ttl     = 1200
  type    = "TXT_RECORD"
  value   = "10.0.8.50"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/d23fff83-65b7-43ad-ae5f-be8483caa794"
resource "unifi_dns_policy" "dns_d23fff83_65b7_43ad_ae5f_be8483caa794" {
  provider = unifi
  enabled = true
  name    = "node-9.corp.internal"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  ttl     = 1500
  type    = "FORWARD_DOMAIN"
  value   = "10.0.9.50"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/cdc0f4c4-860b-406e-bcf8-e87391476c46"
resource "unifi_dns_policy" "dns_cdc0f4c4_860b_406e_bcf8_e87391476c46" {
  provider = unifi
  enabled = true
  name    = "node-10.corp.internal"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  ttl     = 300
  type    = "A_RECORD"
  value   = "10.0.10.50"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/d3511904-19fc-4cfd-bdb6-337547aa5c6d"
resource "unifi_dns_policy" "dns_d3511904_19fc_4cfd_bdb6_337547aa5c6d" {
  provider = unifi
  enabled = true
  name    = "node-11.corp.internal"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  ttl     = 600
  type    = "AAAA_RECORD"
  value   = "10.0.11.50"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/e5b1d59e-b787-4cb1-92b9-698460845dfa"
resource "unifi_dns_policy" "dns_e5b1d59e_b787_4cb1_92b9_698460845dfa" {
  provider = unifi
  enabled = true
  name    = "node-12.corp.internal"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  ttl     = 900
  type    = "CNAME_RECORD"
  value   = "10.0.12.50"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/b7c2c9f6-b3b1-49f8-b1b1-1ce20c7d3695"
resource "unifi_dns_policy" "dns_b7c2c9f6_b3b1_49f8_b1b1_1ce20c7d3695" {
  provider = unifi
  enabled = true
  name    = "node-13.corp.internal"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  ttl     = 1200
  type    = "TXT_RECORD"
  value   = "10.0.13.50"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/1e51e948-6436-4efc-a0a1-1f1876cdb084"
resource "unifi_dns_policy" "dns_1e51e948_6436_4efc_a0a1_1f1876cdb084" {
  provider = unifi
  enabled = true
  name    = "node-14.corp.internal"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  ttl     = 1500
  type    = "FORWARD_DOMAIN"
  value   = "10.0.14.50"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/4813cf49-1bd0-4303-a539-3e24f6d3391e"
resource "unifi_dns_policy" "dns_4813cf49_1bd0_4303_a539_3e24f6d3391e" {
  provider = unifi
  enabled = true
  name    = "node-15.corp.internal"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  ttl     = 300
  type    = "A_RECORD"
  value   = "10.0.15.50"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/fa139475-dfb2-4c81-866b-a2d2e01534ae"
resource "unifi_traffic_matching_list" "traffic_fa139475_dfb2_4c81_866b_a2d2e01534ae" {
  name    = "Traffic-Match-Group-1"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  targets = ["APP-GROUP-1", "DOMAIN-BLOCK-1"]
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/eec590db-7cd8-4731-ba3f-2f3bb2d5421e"
resource "unifi_traffic_matching_list" "traffic_eec590db_7cd8_4731_ba3f_2f3bb2d5421e" {
  name    = "Traffic-Match-Group-2"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  targets = ["APP-GROUP-2", "DOMAIN-BLOCK-2"]
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/2cd2e344-9766-4b0e-ac14-e73450677787"
resource "unifi_traffic_matching_list" "traffic_2cd2e344_9766_4b0e_ac14_e73450677787" {
  name    = "Traffic-Match-Group-3"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  targets = ["APP-GROUP-3", "DOMAIN-BLOCK-3"]
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/035c4675-0b72-4202-af4d-83a328aa75a6"
resource "unifi_traffic_matching_list" "traffic_035c4675_0b72_4202_af4d_83a328aa75a6" {
  name    = "Traffic-Match-Group-4"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  targets = ["APP-GROUP-4", "DOMAIN-BLOCK-4"]
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/0a83df12-0880-4ff9-9870-05db9e8ad9cd"
resource "unifi_traffic_matching_list" "traffic_0a83df12_0880_4ff9_9870_05db9e8ad9cd" {
  name    = "Traffic-Match-Group-5"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  targets = ["APP-GROUP-5", "DOMAIN-BLOCK-5"]
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/7727ff6b-763a-4561-ae3c-eb5b02dcd7c3"
resource "unifi_traffic_matching_list" "traffic_7727ff6b_763a_4561_ae3c_eb5b02dcd7c3" {
  name    = "Traffic-Match-Group-6"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  targets = ["APP-GROUP-6", "DOMAIN-BLOCK-6"]
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/14f06617-5538-4eb2-8554-2cea5a5195c0"
resource "unifi_traffic_matching_list" "traffic_14f06617_5538_4eb2_8554_2cea5a5195c0" {
  name    = "Traffic-Match-Group-7"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  targets = ["APP-GROUP-7", "DOMAIN-BLOCK-7"]
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/4d148255-3ee9-4173-bc6e-40409959fa8c"
resource "unifi_traffic_matching_list" "traffic_4d148255_3ee9_4173_bc6e_40409959fa8c" {
  name    = "Traffic-Match-Group-8"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  targets = ["APP-GROUP-8", "DOMAIN-BLOCK-8"]
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/5a0da1e0-6566-4315-be6a-156db003e9a9"
resource "unifi_traffic_matching_list" "traffic_5a0da1e0_6566_4315_be6a_156db003e9a9" {
  name    = "Traffic-Match-Group-9"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  targets = ["APP-GROUP-9", "DOMAIN-BLOCK-9"]
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/86d03534-4948-4bc6-a5c0-1cedf888ca94"
resource "unifi_traffic_matching_list" "traffic_86d03534_4948_4bc6_a5c0_1cedf888ca94" {
  name    = "Traffic-Match-Group-10"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  targets = ["APP-GROUP-10", "DOMAIN-BLOCK-10"]
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/844f2f12-0da8-4ba3-a3eb-d871901c1638"
resource "unifi_acl_rule" "acl_844f2f12_0da8_4ba3_a3eb_d871901c1638" {
  name    = "ACL-Rule-Strict-1"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  action  = "REJECT"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/18434928-1429-48d3-89ae-52f18d4ab0e6"
resource "unifi_acl_rule" "acl_18434928_1429_48d3_89ae_52f18d4ab0e6" {
  name    = "ACL-Rule-Strict-2"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  action  = "REJECT"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/d6900e08-cef3-4334-8eb9-018f4fed2253"
resource "unifi_acl_rule" "acl_d6900e08_cef3_4334_8eb9_018f4fed2253" {
  name    = "ACL-Rule-Strict-3"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  action  = "REJECT"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/b4715ffb-8e3c-49d0-b5ad-ddb2641a7394"
resource "unifi_acl_rule" "acl_b4715ffb_8e3c_49d0_b5ad_ddb2641a7394" {
  name    = "ACL-Rule-Strict-4"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  action  = "REJECT"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/7ca8302f-d5cb-48d3-88f2-6ffb513faa93"
resource "unifi_acl_rule" "acl_7ca8302f_d5cb_48d3_88f2_6ffb513faa93" {
  name    = "ACL-Rule-Strict-5"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  action  = "REJECT"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/f46fc164-bc1a-4c71-ab75-5e35094e8642"
resource "unifi_acl_rule" "acl_f46fc164_bc1a_4c71_ab75_5e35094e8642" {
  name    = "ACL-Rule-Strict-6"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  action  = "REJECT"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/62aca731-3f2b-43c1-9197-cc63efa16e9c"
resource "unifi_acl_rule" "acl_62aca731_3f2b_43c1_9197_cc63efa16e9c" {
  name    = "ACL-Rule-Strict-7"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  action  = "REJECT"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/6724ea80-f632-4729-9f45-c91e72a81513"
resource "unifi_acl_rule" "acl_6724ea80_f632_4729_9f45_c91e72a81513" {
  name    = "ACL-Rule-Strict-8"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  action  = "REJECT"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/1e3e5949-44b4-4e00-a24c-5cf3044944ac"
resource "unifi_acl_rule" "acl_1e3e5949_44b4_4e00_a24c_5cf3044944ac" {
  name    = "ACL-Rule-Strict-9"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  action  = "REJECT"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/8a07ff02-eb86-43f5-b3d5-aa23d1568239"
resource "unifi_acl_rule" "acl_8a07ff02_eb86_43f5_b3d5_aa23d1568239" {
  name    = "ACL-Rule-Strict-10"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  action  = "REJECT"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/466ed34a-1e3d-4d3b-8b02-e99ee7f0262a"
resource "unifi_device" "dev_466ed34a_1e3d_4d3b_8b02_e99ee7f0262a" {
  name    = "USW-Enterprise-48-PoE-Switch-1"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  mac     = "74:83:c2:01:11:22"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/60a06a8e-8cd6-48c9-9a72-8161aba13b12"
resource "unifi_device" "dev_60a06a8e_8cd6_48c9_9a72_8161aba13b12" {
  name    = "USW-Enterprise-48-PoE-Switch-2"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  mac     = "74:83:c2:02:11:22"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/6ebc28b8-2006-44d9-aaf5-ca3f94238611"
resource "unifi_device" "dev_6ebc28b8_2006_44d9_aaf5_ca3f94238611" {
  name    = "USW-Enterprise-48-PoE-Switch-3"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  mac     = "74:83:c2:03:11:22"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/5d9ecbc3-0cdb-4b2f-814f-6f4e126d392e"
resource "unifi_device" "dev_5d9ecbc3_0cdb_4b2f_814f_6f4e126d392e" {
  name    = "USW-Enterprise-48-PoE-Switch-4"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  mac     = "74:83:c2:04:11:22"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/8252e3ec-7bfe-429d-9040-3c6783a118de"
resource "unifi_device" "dev_8252e3ec_7bfe_429d_9040_3c6783a118de" {
  name    = "USW-Enterprise-48-PoE-Switch-5"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  mac     = "74:83:c2:05:11:22"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/7a4fcc8f-c714-46cb-addc-5afa99bbfe71"
resource "unifi_device" "dev_7a4fcc8f_c714_46cb_addc_5afa99bbfe71" {
  name    = "USW-Enterprise-48-PoE-Switch-6"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  mac     = "74:83:c2:06:11:22"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/d71a2a65-2221-452a-b59a-905f42d8f3a9"
resource "unifi_device" "dev_d71a2a65_2221_452a_b59a_905f42d8f3a9" {
  name    = "USW-Enterprise-48-PoE-Switch-7"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  mac     = "74:83:c2:07:11:22"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/4e6be466-4941-4ba1-975e-addd87d62a57"
resource "unifi_device" "dev_4e6be466_4941_4ba1_975e_addd87d62a57" {
  name    = "USW-Enterprise-48-PoE-Switch-8"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  mac     = "74:83:c2:08:11:22"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/72ee34b6-d644-43d2-8796-59f155d5669d"
resource "unifi_device" "dev_72ee34b6_d644_43d2_8796_59f155d5669d" {
  name    = "USW-Enterprise-48-PoE-Switch-9"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  mac     = "74:83:c2:09:11:22"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/245496b5-fa28-4b77-a4fc-a8073f9f4fc1"
resource "unifi_device" "dev_245496b5_fa28_4b77_a4fc_a8073f9f4fc1" {
  name    = "USW-Enterprise-48-PoE-Switch-10"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  mac     = "74:83:c2:0a:11:22"
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/fw-ordering-master"
resource "unifi_firewall_policy_ordering" "fw_ordering_master" {
  site_id    = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  policy_ids = ["p1", "p2", "p3"]
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/acl-ordering-master"
resource "unifi_acl_rule_ordering" "acl_ordering_master" {
  site_id  = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  rule_ids = ["a1", "a2", "a3"]
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/07e198a7-a38b-4430-94cf-7971eb7fde7d"
resource "unifi_hotspot_voucher" "voucher_07e198a7_a38b_4430_94cf_7971eb7fde7d" {
  site_id          = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  code             = "VIP-CODE-0001"
  duration_minutes = 1440
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/c7e6d272-1856-41b6-a8fa-765a6eed70c0"
resource "unifi_hotspot_voucher" "voucher_c7e6d272_1856_41b6_a8fa_765a6eed70c0" {
  site_id          = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  code             = "VIP-CODE-0002"
  duration_minutes = 2880
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/2849d000-7d62-4900-99bd-1a555bfd897a"
resource "unifi_hotspot_voucher" "voucher_2849d000_7d62_4900_99bd_1a555bfd897a" {
  site_id          = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  code             = "VIP-CODE-0003"
  duration_minutes = 4320
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/cf58a5fd-5042-4e7c-a513-82f5fe42c31e"
resource "unifi_hotspot_voucher" "voucher_cf58a5fd_5042_4e7c_a513_82f5fe42c31e" {
  site_id          = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  code             = "VIP-CODE-0004"
  duration_minutes = 5760
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/15c6996c-3d5e-4a77-835f-f07ba696d898"
resource "unifi_hotspot_voucher" "voucher_15c6996c_3d5e_4a77_835f_f07ba696d898" {
  site_id          = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  code             = "VIP-CODE-0005"
  duration_minutes = 7200
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/5bd2c5f8-79e1-4a97-8846-250d38eb3ca6"
resource "unifi_wifi_broadcast" "wifi_5bd2c5f8_79e1_4a97_8846_250d38eb3ca6" {
  name    = "Corporate-SSID-1"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  enabled = true
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/e5c8cc3f-dba8-4e34-bf61-cf8d5530f789"
resource "unifi_wifi_broadcast" "wifi_e5c8cc3f_dba8_4e34_bf61_cf8d5530f789" {
  name    = "Corporate-SSID-2"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  enabled = true
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/6dc68fe4-2c15-42d8-8ecf-b57e366fdeba"
resource "unifi_wifi_broadcast" "wifi_6dc68fe4_2c15_42d8_8ecf_b57e366fdeba" {
  name    = "Corporate-SSID-3"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  enabled = true
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/ef36f67d-979c-4658-b249-5414c6c5b964"
resource "unifi_wifi_broadcast" "wifi_ef36f67d_979c_4658_b249_5414c6c5b964" {
  name    = "Corporate-SSID-4"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  enabled = true
}

# __generated__ by OpenTofu from "88f7af54-98f8-306a-a1c7-c9349722b1f6/7f881cca-3972-4023-a2ff-cd95cf6b9cce"
resource "unifi_wifi_broadcast" "wifi_7f881cca_3972_4023_a2ff_cd95cf6b9cce" {
  name    = "Corporate-SSID-5"
  site_id = "88f7af54-98f8-306a-a1c7-c9349722b1f6"
  enabled = true
}

# __generated__ by OpenTofu from "314bb60d-405b-4335-a320-49e79c938cd4"
resource "unifi_protect_camera" "cam_314bb60d_405b_4335_a320_49e79c938cd4" {
  name = "G5-Bullet-HQ-Zone-1"
}

# __generated__ by OpenTofu from "a1c85624-02f4-4c7b-996e-6c5f14b23fca"
resource "unifi_protect_camera" "cam_a1c85624_02f4_4c7b_996e_6c5f14b23fca" {
  name = "G5-Bullet-HQ-Zone-2"
}

# __generated__ by OpenTofu from "51885e50-f2bc-4b4d-8057-6f73e9a13fb1"
resource "unifi_protect_camera" "cam_51885e50_f2bc_4b4d_8057_6f73e9a13fb1" {
  name = "G5-Bullet-HQ-Zone-3"
}

# __generated__ by OpenTofu from "92c05aaf-a9a3-450a-bd67-1dd1adff9f4b"
resource "unifi_protect_camera" "cam_92c05aaf_a9a3_450a_bd67_1dd1adff9f4b" {
  name = "G5-Bullet-HQ-Zone-4"
}

# __generated__ by OpenTofu from "c3e59f90-a535-438c-9d7c-7d0aef06b817"
resource "unifi_protect_camera" "cam_c3e59f90_a535_438c_9d7c_7d0aef06b817" {
  name = "G5-Bullet-HQ-Zone-5"
}

# __generated__ by OpenTofu from "5934429d-3eda-4f34-83d9-f0489257ec17"
resource "unifi_protect_camera" "cam_5934429d_3eda_4f34_83d9_f0489257ec17" {
  name = "G5-Bullet-HQ-Zone-6"
}

# __generated__ by OpenTofu from "5f17636c-827e-4e04-bd67-5e2036f69632"
resource "unifi_protect_camera" "cam_5f17636c_827e_4e04_bd67_5e2036f69632" {
  name = "G5-Bullet-HQ-Zone-7"
}

# __generated__ by OpenTofu from "778135d2-1c9c-4ec2-a35c-9668de5e8ef6"
resource "unifi_protect_camera" "cam_778135d2_1c9c_4ec2_a35c_9668de5e8ef6" {
  name = "G5-Bullet-HQ-Zone-8"
}

# __generated__ by OpenTofu from "fbb1ee89-fc4e-4cef-b75f-3bcb1486f0ff"
resource "unifi_protect_camera" "cam_fbb1ee89_fc4e_4cef_b75f_3bcb1486f0ff" {
  name = "G5-Bullet-HQ-Zone-9"
}

# __generated__ by OpenTofu from "6f41019f-7dd8-42a7-87ba-5ac076e1d476"
resource "unifi_protect_camera" "cam_6f41019f_7dd8_42a7_87ba_5ac076e1d476" {
  name = "G5-Bullet-HQ-Zone-10"
}

# __generated__ by OpenTofu from "8c9db9bc-11de-4b55-9210-d3a4e6b55ab3"
resource "unifi_protect_camera" "cam_8c9db9bc_11de_4b55_9210_d3a4e6b55ab3" {
  name = "G5-Bullet-HQ-Zone-11"
}

# __generated__ by OpenTofu from "0b34b9f0-94c4-4ef8-9161-05b5427c8fb8"
resource "unifi_protect_camera" "cam_0b34b9f0_94c4_4ef8_9161_05b5427c8fb8" {
  name = "G5-Bullet-HQ-Zone-12"
}

# __generated__ by OpenTofu from "caa2e690-c39f-416c-88d6-c774bc209ae3"
resource "unifi_protect_camera" "cam_caa2e690_c39f_416c_88d6_c774bc209ae3" {
  name = "G5-Bullet-HQ-Zone-13"
}

# __generated__ by OpenTofu from "a1413b3e-331a-4607-9832-1f51d0d51fdd"
resource "unifi_protect_camera" "cam_a1413b3e_331a_4607_9832_1f51d0d51fdd" {
  name = "G5-Bullet-HQ-Zone-14"
}

# __generated__ by OpenTofu from "801aad41-db8c-4c59-a9d1-a96fd4c59748"
resource "unifi_protect_camera" "cam_801aad41_db8c_4c59_a9d1_a96fd4c59748" {
  name = "G5-Bullet-HQ-Zone-15"
}

# __generated__ by OpenTofu from "fac5e000-8d9e-490d-a25f-a2b0ca50f82b"
resource "unifi_protect_chime" "chime_fac5e000_8d9e_490d_a25f_a2b0ca50f82b" {
  name   = "Chime-Building-1"
  volume = 71
}

# __generated__ by OpenTofu from "f55bbadc-5b60-4907-9166-39387b286aad"
resource "unifi_protect_chime" "chime_f55bbadc_5b60_4907_9166_39387b286aad" {
  name   = "Chime-Building-2"
  volume = 72
}

# __generated__ by OpenTofu from "2221809b-aa1e-4f61-9196-7bf9a76423c1"
resource "unifi_protect_chime" "chime_2221809b_aa1e_4f61_9196_7bf9a76423c1" {
  name   = "Chime-Building-3"
  volume = 73
}

# __generated__ by OpenTofu from "529db545-4ea7-4294-b831-aac2b50482c7"
resource "unifi_protect_chime" "chime_529db545_4ea7_4294_b831_aac2b50482c7" {
  name   = "Chime-Building-4"
  volume = 74
}

# __generated__ by OpenTofu from "151c404a-d41d-4823-a7dd-0680eca9a602"
resource "unifi_protect_chime" "chime_151c404a_d41d_4823_a7dd_0680eca9a602" {
  name   = "Chime-Building-5"
  volume = 75
}

# __generated__ by OpenTofu from "4b2436a9-e569-4940-bf1e-eb947fedcb62"
resource "unifi_protect_light" "light_4b2436a9_e569_4940_bf1e_eb947fedcb62" {
  name = "Smart-Floodlight-1"
}

# __generated__ by OpenTofu from "6cddfc1a-562b-4f24-9edf-d69a4b13bb48"
resource "unifi_protect_light" "light_6cddfc1a_562b_4f24_9edf_d69a4b13bb48" {
  name = "Smart-Floodlight-2"
}

# __generated__ by OpenTofu from "329768b3-b81f-432e-9704-d08a4d36f774"
resource "unifi_protect_light" "light_329768b3_b81f_432e_9704_d08a4d36f774" {
  name = "Smart-Floodlight-3"
}

# __generated__ by OpenTofu from "fd752ecd-2a28-4a55-8d77-63c079072c0e"
resource "unifi_protect_light" "light_fd752ecd_2a28_4a55_8d77_63c079072c0e" {
  name = "Smart-Floodlight-4"
}

# __generated__ by OpenTofu from "45d320a6-5bb0-4f5a-9504-58a15cd200bf"
resource "unifi_protect_light" "light_45d320a6_5bb0_4f5a_9504_58a15cd200bf" {
  name = "Smart-Floodlight-5"
}

# __generated__ by OpenTofu from "e22b111c-5c17-436a-86b1-a3c81e8510ab"
resource "unifi_protect_liveview" "lv_e22b111c_5c17_436a_86b1_a3c81e8510ab" {
  name = "NOC-Monitoring-Matrix-1"
}

# __generated__ by OpenTofu from "b7cf6e3f-e24b-42e1-965f-dbecfca190c7"
resource "unifi_protect_liveview" "lv_b7cf6e3f_e24b_42e1_965f_dbecfca190c7" {
  name = "NOC-Monitoring-Matrix-2"
}

# __generated__ by OpenTofu from "88322441-6fcf-4e0d-a122-1b149f622e57"
resource "unifi_protect_liveview" "lv_88322441_6fcf_4e0d_a122_1b149f622e57" {
  name = "NOC-Monitoring-Matrix-3"
}

# __generated__ by OpenTofu from "ba4b517a-9310-447d-baab-751ea0d6ef25"
resource "unifi_protect_liveview" "lv_ba4b517a_9310_447d_baab_751ea0d6ef25" {
  name = "NOC-Monitoring-Matrix-4"
}

# __generated__ by OpenTofu from "a8767f2d-e9c0-43dc-8e00-6b5bc711fdbd"
resource "unifi_protect_liveview" "lv_a8767f2d_e9c0_43dc_8e00_6b5bc711fdbd" {
  name = "NOC-Monitoring-Matrix-5"
}

# __generated__ by OpenTofu from "0e830fe9-9bec-41db-8555-db033ed89cb6"
resource "unifi_protect_relay" "relay_0e830fe9_9bec_41db_8555_db033ed89cb6" {
  name = "Access-Gate-Relay-1"
}

# __generated__ by OpenTofu from "a660820b-cc19-493e-87b7-9a7ebb18fe14"
resource "unifi_protect_relay" "relay_a660820b_cc19_493e_87b7_9a7ebb18fe14" {
  name = "Access-Gate-Relay-2"
}

# __generated__ by OpenTofu from "5efec4d3-bc77-479d-babc-51485b27d826"
resource "unifi_protect_relay" "relay_5efec4d3_bc77_479d_babc_51485b27d826" {
  name = "Access-Gate-Relay-3"
}

# __generated__ by OpenTofu from "41f532c6-ea95-471d-8314-f37597254ed7"
resource "unifi_protect_relay" "relay_41f532c6_ea95_471d_8314_f37597254ed7" {
  name = "Access-Gate-Relay-4"
}

# __generated__ by OpenTofu from "b9f8905b-08ae-4b13-9281-dc60e0dd11cd"
resource "unifi_protect_relay" "relay_b9f8905b_08ae_4b13_9281_dc60e0dd11cd" {
  name = "Access-Gate-Relay-5"
}

# __generated__ by OpenTofu from "5d3155c5-03a2-4152-99c7-cea5762e3546"
resource "unifi_protect_sensor" "sensor_5d3155c5_03a2_4152_99c7_cea5762e3546" {
  name  = "Multi-Sensor-Room-1"
  alarm = true
}

# __generated__ by OpenTofu from "0e55cf95-3eb0-427f-a9de-92f53596d487"
resource "unifi_protect_sensor" "sensor_0e55cf95_3eb0_427f_a9de_92f53596d487" {
  name  = "Multi-Sensor-Room-2"
  alarm = true
}

# __generated__ by OpenTofu from "e6e15bd6-5c17-4cec-8e43-1437289afaee"
resource "unifi_protect_sensor" "sensor_e6e15bd6_5c17_4cec_8e43_1437289afaee" {
  name  = "Multi-Sensor-Room-3"
  alarm = true
}

# __generated__ by OpenTofu from "da49ed09-ca18-415e-a165-956f836e231f"
resource "unifi_protect_sensor" "sensor_da49ed09_ca18_415e_a165_956f836e231f" {
  name  = "Multi-Sensor-Room-4"
  alarm = true
}

# __generated__ by OpenTofu from "c124bff8-e01c-4260-bf99-08fdf2d14016"
resource "unifi_protect_sensor" "sensor_c124bff8_e01c_4260_bf99_08fdf2d14016" {
  name  = "Multi-Sensor-Room-5"
  alarm = true
}

# __generated__ by OpenTofu from "3791ee0b-68a6-407b-9593-750a3ec31671"
resource "unifi_protect_siren" "siren_3791ee0b_68a6_407b_9593_750a3ec31671" {
  name   = "Strobe-Siren-Floor-1"
  volume = 100
}

# __generated__ by OpenTofu from "300ba208-7fe0-4b1f-b3d4-c7e5a5cce387"
resource "unifi_protect_siren" "siren_300ba208_7fe0_4b1f_b3d4_c7e5a5cce387" {
  name   = "Strobe-Siren-Floor-2"
  volume = 100
}

# __generated__ by OpenTofu from "3c367704-db02-437e-bbc0-20f4c3a3fc2a"
resource "unifi_protect_siren" "siren_3c367704_db02_437e_bbc0_20f4c3a3fc2a" {
  name   = "Strobe-Siren-Floor-3"
  volume = 100
}

# __generated__ by OpenTofu from "aff8e4c0-463e-4bef-acc5-b05dcc467e8e"
resource "unifi_protect_siren" "siren_aff8e4c0_463e_4bef_acc5_b05dcc467e8e" {
  name   = "Strobe-Siren-Floor-4"
  volume = 100
}

# __generated__ by OpenTofu from "6f63f834-560c-4009-a2db-0e800793dadc"
resource "unifi_protect_siren" "siren_6f63f834_560c_4009_a2db_0e800793dadc" {
  name   = "Strobe-Siren-Floor-5"
  volume = 100
}

# __generated__ by OpenTofu from "774ab766-a011-451d-92dc-7081e4e1e0e9"
resource "unifi_protect_viewer" "viewer_774ab766_a011_451d_92dc_7081e4e1e0e9" {
  name = "Viewport-Lobby-Display-1"
}

# __generated__ by OpenTofu from "1660e615-45d8-4c9a-bf7f-46c33b55a707"
resource "unifi_protect_viewer" "viewer_1660e615_45d8_4c9a_bf7f_46c33b55a707" {
  name = "Viewport-Lobby-Display-2"
}

# __generated__ by OpenTofu from "5965392c-1507-41a5-8aff-3b87d5fa5db3"
resource "unifi_protect_viewer" "viewer_5965392c_1507_41a5_8aff_3b87d5fa5db3" {
  name = "Viewport-Lobby-Display-3"
}

# __generated__ by OpenTofu from "2e5341d5-9b37-4723-8237-f7cbc7ca9727"
resource "unifi_protect_viewer" "viewer_2e5341d5_9b37_4723_8237_f7cbc7ca9727" {
  name = "Viewport-Lobby-Display-4"
}

# __generated__ by OpenTofu from "4b3fb37c-8d24-4d46-aa13-672d0920aebe"
resource "unifi_protect_viewer" "viewer_4b3fb37c_8d24_4d46_aa13_672d0920aebe" {
  name = "Viewport-Lobby-Display-5"
}