package node

import (
	"bytes"
)

func ClientBody(name string) *bytes.Buffer {
	body := `{
  "name": "` + name + `",
  "admin": false
}`

	return string_to_reader(body)
}

func EndRunBody(run_uuid string, start_time string) *bytes.Buffer {
	body := `{
  "action": "end",
  "resources": [
    {
      "type": "execute",
      "name": "an execute resource",
      "id": "echo 'hello'",
      "after": {},
      "before": {},
      "duration": "15340",
      "delta": "",
      "result": "run",
      "cookbook_name": "dotfiles",
      "cookbook_version": "1.1.1"
    },
        {
      "type": "users_manage",
      "name": "sysadmin",
      "id": "sysadmin",
      "after": {
        "cookbook": "users",
        "data_bag": "users",
        "group_id": 2300,
        "group_name": "sysadmin",
        "search_group": "sysadmin"
      },
      "before": {
        
      },
      "duration": "903",
      "delta": "",
      "result": "create",
      "cookbook_name": "users",
      "cookbook_version": "1.8.2"
    },
    {
      "type": "template",
      "name": "\/etc\/sudoers",
      "id": "\/etc\/sudoers",
      "after": {
        "checksum": "72cd01dc3a319d596b32090d23e84b836c9bf15d948f06c861df8b718aa19976",
        "owner": "root",
        "group": "root",
        "mode": "0440"
      },
      "before": {
        "checksum": "ce7a6431cd28b20515e6503dafe1f51884045b894a643711c43d8041ba7ed5fd",
        "owner": "root",
        "group": "root",
        "mode": "0440"
      },
      "duration": "19",
      "delta": "REDACTED",
      "result": "create",
      "cookbook_name": "sudo",
      "cookbook_version": "2.7.1"
    }
  ],
  "status": "success",
  "run_list": "[\"role[base]\"]",
  "start_time": "` + start_time + `",
  "end_time": "` + timestamp() + `",
  "total_res_count": "46",
  "data": {}
}`

	return string_to_reader(body)
}

func StartRunBody(run_uuid string, start_time string) *bytes.Buffer {
	body := `{
  "action": "start",
  "run_id": "` + run_uuid + `",
  "start_time": "` + start_time + `"
}`

	return string_to_reader(body)
}

func CookbookVersionsBody() *bytes.Buffer {
	body := `{
  "run_list": [
    "apache2",
    "apt",
    "build-essential",
    "chef-client@4.3.0",
    "cron",
    "dotfiles",
    "java",
    "nginx",
    "ntp@1.8.6",
    "perl",
    "postfix",
    "runit",
    "users@1.8.2",
    "windows",
    "yum",
    "yum-epel"
  ]
}`

	return string_to_reader(body)
}

func NodeBody(name string) *bytes.Buffer {
	body := `{
  "name":"` + name + `",
  "chef_environment": "_default",
  "chef_environment": "rs-prod",
  "run_list": [
    "role[base]"
  ],
	"automatic": {
		"languages": {
			"python": {
				"version": "2.7.6",
				"builddate": "Mar 22 2014, 22:59:56"
			},
			"ruby": {
				"platform": "x86_64-linux",
				"version": "2.1.6",
				"release_date": "2015-04-13",
				"target": "x86_64-unknown-linux-gnu",
				"target_cpu": "x86_64",
				"target_vendor": "unknown",
				"target_os": "linux",
				"host": "x86_64-unknown-linux-gnu",
				"host_cpu": "x86_64",
				"host_os": "linux-gnu",
				"host_vendor": "unknown",
				"bin_dir": "/opt/chef/embedded/bin",
				"ruby_bin": "/opt/chef/embedded/bin/ruby",
				"gems_dir": "/opt/chef/embedded/lib/ruby/gems/2.1.0",
				"gem_bin": "/opt/chef/embedded/bin/gem"
			},
			"perl": {
				"version": "5.18.2",
				"archname": "x86_64-linux-gnu-thread-multi"
			}
		},
		"network": {
			"interfaces": {
				"lo": {
					"mtu": "65536",
					"flags": ["LOOPBACK", "UP", "LOWER_UP"],
					"encapsulation": "Loopback",
					"addresses": {
						"127.0.0.1": {
							"family": "inet",
							"prefixlen": "8",
							"netmask": "255.0.0.0",
							"scope": "Node"
						},
						"::1": {
							"family": "inet6",
							"prefixlen": "128",
							"scope": "Node"
						}
					},
					"state": "unknown"
				},
				"eth0": {
					"type": "eth",
					"number": "0",
					"mtu": "1500",
					"flags": ["BROADCAST", "MULTICAST", "UP", "LOWER_UP"],
					"encapsulation": "Ethernet",
					"addresses": {
						"00:11:22:33:44:92": {
							"family": "lladdr"
						},
						"192.0.0.100": {
							"family": "inet",
							"prefixlen": "22",
							"netmask": "255.255.255.0",
							"broadcast": "192.0.0.255",
							"scope": "Global"
						},
						"fe80::20c:1234:1234:1234": {
							"family": "inet6",
							"prefixlen": "64",
							"scope": "Link"
						}
					},
					"state": "up",
					"routes": [{
						"destination": "default",
						"family": "inet",
						"via": "192.0.0.1"
					}, {
						"destination": "192.0.0.0/24",
						"family": "inet",
						"scope": "link",
						"proto": "kernel",
						"src": "172.28.1.35"
					}, {
						"destination": "fe80::/64",
						"family": "inet6",
						"metric": "256",
						"proto": "kernel"
					}]
				}
			},
			"default_interface": "eth0",
			"default_gateway": "172.28.0.1"
		},
		"counters": {
			"network": {
				"interfaces": {
					"lo": {
						"rx": {
							"bytes": "0",
							"packets": "0",
							"errors": "0",
							"drop": "0",
							"overrun": "0"
						},
						"tx": {
							"bytes": "0",
							"packets": "0",
							"errors": "0",
							"drop": "0",
							"carrier": "0",
							"collisions": "0"
						}
					},
					"eth0": {
						"tx": {
							"queuelen": "1000",
							"bytes": "1794153",
							"packets": "18992",
							"errors": "0",
							"drop": "0",
							"carrier": "0",
							"collisions": "0"
						},
						"rx": {
							"bytes": "45138617",
							"packets": "40575",
							"errors": "0",
							"drop": "0",
							"overrun": "0"
						}
					}
				}
			}
		},
		"ipaddress": "192.0.0.100",
		"macaddress": "00:11:22:33:44:55",
		"cpu": {
			"0": {
				"vendor_id": "GenuineIntel",
				"family": "6",
				"model": "70",
				"model_name": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
				"stepping": "1",
				"mhz": "2794.041",
				"cache_size": "6144 KB",
				"physical_id": "0",
				"core_id": "0",
				"cores": "1",
				"flags": ["fpu", "vme", "de", "pse", "tsc", "msr", "pae", "mce", "cx8", "apic", "sep", "mtrr", "pge", "mca", "cmov", "pat", "pse36", "clflush", "dts", "mmx", "fxsr", "sse", "sse2", "ss", "syscall", "nx", "pdpe1gb", "rdtscp", "lm", "constant_tsc", "arch_perfmon", "pebs", "bts", "nopl", "xtopology", "tsc_reliable", "nonstop_tsc", "aperfmperf", "eagerfpu", "pni", "pclmulqdq", "ssse3", "fma", "cx16", "pcid", "sse4_1", "sse4_2", "x2apic", "movbe", "popcnt", "aes", "xsave", "avx", "f16c", "rdrand", "hypervisor", "lahf_lm", "ida", "arat", "epb", "xsaveopt", "pln", "pts", "dtherm", "fsgsbase", "smep"]
			},
			"1": {
				"vendor_id": "GenuineIntel",
				"family": "6",
				"model": "70",
				"model_name": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
				"stepping": "1",
				"mhz": "2794.041",
				"cache_size": "6144 KB",
				"physical_id": "2",
				"core_id": "0",
				"cores": "1",
				"flags": ["fpu", "vme", "de", "pse", "tsc", "msr", "pae", "mce", "cx8", "apic", "sep", "mtrr", "pge", "mca", "cmov", "pat", "pse36", "clflush", "dts", "mmx", "fxsr", "sse", "sse2", "ss", "syscall", "nx", "pdpe1gb", "rdtscp", "lm", "constant_tsc", "arch_perfmon", "pebs", "bts", "nopl", "xtopology", "tsc_reliable", "nonstop_tsc", "aperfmperf", "eagerfpu", "pni", "pclmulqdq", "ssse3", "fma", "cx16", "pcid", "sse4_1", "sse4_2", "x2apic", "movbe", "popcnt", "aes", "xsave", "avx", "f16c", "rdrand", "hypervisor", "lahf_lm", "ida", "arat", "epb", "xsaveopt", "pln", "pts", "dtherm", "fsgsbase", "smep"]
			},
			"total": 2,
			"real": 2
		},
		"ip6address": "fe80::20c:29ff:fef3:5f92",
		"memory": {
			"swap": {
				"cached": "0kB",
				"total": "2094076kB",
				"free": "2094076kB"
			},
			"total": "2042644kB",
			"free": "1289132kB",
			"buffers": "27492kB",
			"cached": "492648kB",
			"active": "377088kB",
			"inactive": "242700kB",
			"dirty": "48kB",
			"writeback": "0kB",
			"anon_pages": "99612kB",
			"mapped": "18212kB",
			"slab": "76660kB",
			"slab_reclaimable": "55704kB",
			"slab_unreclaim": "20956kB",
			"page_tables": "2716kB",
			"nfs_unstable": "0kB",
			"bounce": "0kB",
			"commit_limit": "3115396kB",
			"committed_as": "189948kB",
			"vmalloc_total": "34359738367kB",
			"vmalloc_used": "181280kB",
			"vmalloc_chunk": "34359535100kB"
		},
		"lsb": {
			"id": "Ubuntu",
			"release": "14.04",
			"codename": "trusty",
			"description": "Ubuntu 14.04.2 LTS"
		},
		"kernel": {
			"name": "Linux",
			"release": "3.16.0-30-generic",
			"version": "#40~14.04.1-Ubuntu SMP Thu Jan 15 17:43:14 UTC 2015",
			"machine": "x86_64",
			"os": "GNU/Linux",
			"modules": {
				"btusb": {
					"size": "32497",
					"refcount": "0"
				},
				"hid_generic": {
					"size": "12559",
					"refcount": "0"
				},
				"bluetooth": {
					"size": "446409",
					"refcount": "2"
				},
				"6lowpan_iphc": {
					"size": "18702",
					"refcount": "1"
				},
				"ppdev": {
					"size": "17671",
					"refcount": "0"
				},
				"coretemp": {
					"size": "13441",
					"refcount": "0"
				},
				"vmw_balloon": {
					"size": "13415",
					"refcount": "0"
				},
				"crct10dif_pclmul": {
					"size": "14307",
					"refcount": "0"
				},
				"crc32_pclmul": {
					"size": "13133",
					"refcount": "0"
				},
				"ghash_clmulni_intel": {
					"size": "13230",
					"refcount": "0"
				},
				"aesni_intel": {
					"size": "152552",
					"refcount": "1"
				},
				"snd_ens1371": {
					"size": "25445",
					"refcount": "0"
				},
				"aes_x86_64": {
					"size": "17131",
					"refcount": "1"
				},
				"lrw": {
					"size": "13286",
					"refcount": "1"
				},
				"gf128mul": {
					"size": "14951",
					"refcount": "1"
				},
				"glue_helper": {
					"size": "13990",
					"refcount": "1"
				},
				"ablk_helper": {
					"size": "13597",
					"refcount": "1"
				},
				"cryptd": {
					"size": "20359",
					"refcount": "3"
				},
				"serio_raw": {
					"size": "13483",
					"refcount": "0"
				},
				"snd_ac97_codec": {
					"size": "130476",
					"refcount": "1"
				},
				"ac97_bus": {
					"size": "12730",
					"refcount": "1"
				},
				"gameport": {
					"size": "15758",
					"refcount": "1"
				},
				"snd_rawmidi": {
					"size": "30876",
					"refcount": "1"
				},
				"snd_seq_device": {
					"size": "14497",
					"refcount": "1"
				},
				"usbhid": {
					"size": "52616",
					"refcount": "0"
				},
				"hid": {
					"size": "110426",
					"refcount": "2"
				},
				"snd_pcm": {
					"size": "104112",
					"refcount": "2"
				},
				"snd_timer": {
					"size": "29562",
					"refcount": "1"
				},
				"vmwgfx": {
					"size": "179698",
					"refcount": "1"
				},
				"snd": {
					"size": "79468",
					"refcount": "6"
				},
				"soundcore": {
					"size": "15047",
					"refcount": "1"
				},
				"ttm": {
					"size": "85314",
					"refcount": "1"
				},
				"drm_kms_helper": {
					"size": "61574",
					"refcount": "1"
				},
				"i2c_piix4": {
					"size": "22166",
					"refcount": "0"
				},
				"drm": {
					"size": "311018",
					"refcount": "4"
				},
				"shpchp": {
					"size": "37047",
					"refcount": "0"
				},
				"vmw_vmci": {
					"size": "62976",
					"refcount": "0"
				},
				"parport_pc": {
					"size": "32741",
					"refcount": "1"
				},
				"mac_hid": {
					"size": "13227",
					"refcount": "0"
				},
				"lp": {
					"size": "17759",
					"refcount": "0"
				},
				"parport": {
					"size": "42348",
					"refcount": "3"
				},
				"psmouse": {
					"size": "106561",
					"refcount": "0"
				},
				"ahci": {
					"size": "34062",
					"refcount": "0"
				},
				"libahci": {
					"size": "32424",
					"refcount": "1"
				},
				"mptspi": {
					"size": "22560",
					"refcount": "2"
				},
				"mptscsih": {
					"size": "40150",
					"refcount": "1"
				},
				"e1000": {
					"size": "133256",
					"refcount": "0"
				},
				"mptbase": {
					"size": "101822",
					"refcount": "2"
				},
				"scsi_transport_spi": {
					"size": "30732",
					"refcount": "1"
				},
				"pata_acpi": {
					"size": "13053",
					"refcount": "0"
				},
				"floppy": {
					"size": "69662",
					"refcount": "0"
				}
			}
		},
		"os": "linux",
		"os_version": "3.16.0-30-generic",
		"platform": "ubuntu",
		"platform_version": "14.04",
		"platform_family": "debian",
		"filesystem": {
			"/dev/sda1": {
				"kb_size": "18447100",
				"kb_used": "1357564",
				"kb_available": "16129436",
				"percent_used": "8%",
				"mount": "/",
				"total_inodes": "1179648",
				"inodes_used": "79273",
				"inodes_available": "1100375",
				"inodes_percent_used": "7%",
				"fs_type": "ext4",
				"mount_options": ["rw", "errors=remount-ro"],
				"uuid": "cfd79998-a6d6-49bf-8649-2b56e75e7a63"
			},
			"none": {
				"kb_size": "102400",
				"kb_used": "0",
				"kb_available": "102400",
				"percent_used": "0%",
				"mount": "/sys/fs/pstore",
				"total_inodes": "255330",
				"inodes_used": "3",
				"inodes_available": "255327",
				"inodes_percent_used": "1%",
				"fs_type": "pstore",
				"mount_options": ["rw"]
			},
			"udev": {
				"kb_size": "1010552",
				"kb_used": "4",
				"kb_available": "1010548",
				"percent_used": "1%",
				"mount": "/dev",
				"total_inodes": "252638",
				"inodes_used": "467",
				"inodes_available": "252171",
				"inodes_percent_used": "1%",
				"fs_type": "devtmpfs",
				"mount_options": ["rw", "mode=0755"]
			},
			"tmpfs": {
				"kb_size": "204268",
				"kb_used": "980",
				"kb_available": "203288",
				"percent_used": "1%",
				"mount": "/run",
				"total_inodes": "255330",
				"inodes_used": "512",
				"inodes_available": "254818",
				"inodes_percent_used": "1%",
				"fs_type": "tmpfs",
				"mount_options": ["rw", "noexec", "nosuid", "size=10%", "mode=0755"]
			},
			"proc": {
				"mount": "/proc",
				"fs_type": "proc",
				"mount_options": ["rw", "noexec", "nosuid", "nodev"]
			},
			"sysfs": {
				"mount": "/sys",
				"fs_type": "sysfs",
				"mount_options": ["rw", "noexec", "nosuid", "nodev"]
			},
			"devpts": {
				"mount": "/dev/pts",
				"fs_type": "devpts",
				"mount_options": ["rw", "noexec", "nosuid", "gid=5", "mode=0620"]
			},
			"systemd": {
				"mount": "/sys/fs/cgroup/systemd",
				"fs_type": "cgroup",
				"mount_options": ["rw", "noexec", "nosuid", "nodev", "none", "name=systemd"]
			},
			"/dev/sda5": {
				"fs_type": "swap",
				"uuid": "4d7ea8cf-b83f-429d-ad59-d0fe7f0039fb"
			},
			"rootfs": {
				"mount": "/",
				"fs_type": "rootfs",
				"mount_options": ["rw"]
			},
			"/dev/disk/by-uuid/cfd79998-a6d6-49bf-8649-2b56e75e7a63": {
				"mount": "/",
				"fs_type": "ext4",
				"mount_options": ["rw", "relatime", "errors=remount-ro", "data=ordered"]
			}
		},
		"cloud_v2": null,
		"virtualization": {
			"systems": {
				"vmware": "guest"
			},
			"system": "vmware",
			"role": "guest"
		},
		"hostname": "ubuntu",
		"machinename": "ubuntu",
		"fqdn": "ubuntu",
		"domain": null,
		"root_group": "root",
		"dmi": {
			"dmidecode_version": "2.12",
			"smbios_version": "2.4",
			"structures": {
				"count": "364",
				"size": "16646"
			},
			"table_location": "0x000E0010",
			"bios": {
				"all_records": [{
					"record_id": "0x0000",
					"size": "0",
					"application_identifier": "BIOS Information",
					"Vendor": "Phoenix Technologies LTD",
					"Version": "6.00",
					"Release Date": "07/31/2013",
					"Address": "0xEA050",
					"Runtime Size": "90032 bytes",
					"ROM Size": "64 kB",
					"Characteristics": {
						"ISA is supported": null,
						"PCI is supported": null,
						"PC Card (PCMCIA) is supported": null,
						"PNP is supported": null,
						"APM is supported": null,
						"BIOS is upgradeable": null,
						"BIOS shadowing is allowed": null,
						"ESCD support is available": null,
						"Boot from CD is supported": null,
						"Selectable boot is supported": null,
						"EDD is supported": null,
						"Print screen service is supported (int 5h)": null,
						"8042 keyboard services are supported (int 9h)": null,
						"Serial services are supported (int 14h)": null,
						"Printer services are supported (int 17h)": null,
						"CGA/mono video services are supported (int 10h)": null,
						"ACPI is supported": null,
						"Smart battery is supported": null,
						"BIOS boot specification is supported": null,
						"Function key-initiated network boot is supported": null,
						"Targeted content distribution is supported": null
					},
					"BIOS Revision": "4.6",
					"Firmware Revision": "0.0"
				}],
				"vendor": "Phoenix Technologies LTD",
				"version": "6.00",
				"release_date": "07/31/2013",
				"address": "0xEA050",
				"runtime_size": "90032 bytes",
				"rom_size": "64 kB",
				"bios_revision": "4.6",
				"firmware_revision": "0.0"
			},
			"system": {
				"all_records": [{
					"record_id": "0x0001",
					"size": "1",
					"application_identifier": "System Information",
					"Manufacturer": "VMware, Inc.",
					"Product Name": "VMware Virtual Platform",
					"Version": "None",
					"Serial Number": "VMware-56 4d 3f de b6 77 4b 6d-0e 61 3b 4f 11 f3 5f 92",
					"UUID": "564D3FDE-B677-4B6D-0E61-3B4F11F35F92",
					"Wake-up Type": "Power Switch",
					"SKU Number": "Not Specified",
					"Family": "Not Specified"
				}],
				"manufacturer": "VMware, Inc.",
				"product_name": "VMware Virtual Platform",
				"version": "None",
				"serial_number": "VMware-56 4d 3f de b6 77 4b 6d-0e 61 3b 4f 11 f3 5f 92",
				"uuid": "564D3FDE-B677-4B6D-0E61-3B4F11F35F92",
				"wake_up_type": "Power Switch",
				"sku_number": "Not Specified",
				"family": "Not Specified"
			},
			"base_board": {
				"all_records": [{
					"record_id": "0x0002",
					"size": "2",
					"application_identifier": "Base Board Information",
					"Manufacturer": "Intel Corporation",
					"Product Name": "440BX Desktop Reference Platform",
					"Version": "None",
					"Serial Number": "None",
					"Asset Tag": "Not Specified",
					"Features": "None",
					"Location In Chassis": "Not Specified",
					"Chassis Handle": "0x0000",
					"Type": "Unknown",
					"Contained Object Handles": "0"
				}],
				"manufacturer": "Intel Corporation",
				"product_name": "440BX Desktop Reference Platform",
				"version": "None",
				"serial_number": "None",
				"asset_tag": "Not Specified",
				"features": "None",
				"location_in_chassis": "Not Specified",
				"chassis_handle": "0x0000",
				"type": "Unknown",
				"contained_object_handles": "0"
			},
			"chassis": {
				"all_records": [{
					"record_id": "0x0003",
					"size": "3",
					"application_identifier": "Chassis Information",
					"Manufacturer": "No Enclosure",
					"Type": "Other",
					"Lock": "Not Present",
					"Version": "N/A",
					"Serial Number": "None",
					"Asset Tag": "No Asset Tag",
					"Boot-up State": "Safe",
					"Power Supply State": "Safe",
					"Thermal State": "Safe",
					"Security Status": "None",
					"OEM Information": "0x00001234",
					"Height": "Unspecified",
					"Number Of Power Cords": "Unspecified",
					"Contained Elements": "0"
				}],
				"manufacturer": "No Enclosure",
				"type": "Other",
				"lock": "Not Present",
				"version": "N/A",
				"serial_number": "None",
				"asset_tag": "No Asset Tag",
				"boot_up_state": "Safe",
				"power_supply_state": "Safe",
				"thermal_state": "Safe",
				"security_status": "None",
				"oem_information": "0x00001234",
				"height": "Unspecified",
				"number_of_power_cords": "Unspecified",
				"contained_elements": "0"
			},
			"processor": {
				"all_records": [{
					"record_id": "0x0004",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #0",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 04 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Enabled",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x0054",
					"L2 Cache Handle": "0x0055",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x0005",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #1",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Enabled",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x0056",
					"L2 Cache Handle": "0x0057",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x0006",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #2",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x0058",
					"L2 Cache Handle": "0x0059",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x0007",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #3",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x005A",
					"L2 Cache Handle": "0x005B",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x0008",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #4",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x005C",
					"L2 Cache Handle": "0x005D",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x0009",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #5",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x005E",
					"L2 Cache Handle": "0x005F",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x000A",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #6",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x0060",
					"L2 Cache Handle": "0x0061",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x000B",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #7",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x0062",
					"L2 Cache Handle": "0x0063",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x000C",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #8",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x0064",
					"L2 Cache Handle": "0x0065",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x000D",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #9",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x0066",
					"L2 Cache Handle": "0x0067",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x000E",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #10",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x0068",
					"L2 Cache Handle": "0x0069",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x000F",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #11",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x006A",
					"L2 Cache Handle": "0x006B",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x0010",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #12",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x006C",
					"L2 Cache Handle": "0x006D",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x0011",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #13",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x006E",
					"L2 Cache Handle": "0x006F",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x0012",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #14",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x0070",
					"L2 Cache Handle": "0x0071",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x0013",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #15",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x0072",
					"L2 Cache Handle": "0x0073",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x0014",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #16",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x0074",
					"L2 Cache Handle": "0x0075",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x0015",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #17",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x0076",
					"L2 Cache Handle": "0x0077",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x0016",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #18",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x0078",
					"L2 Cache Handle": "0x0079",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x0017",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #19",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x007A",
					"L2 Cache Handle": "0x007B",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x0018",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #20",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x007C",
					"L2 Cache Handle": "0x007D",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x0019",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #21",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x007E",
					"L2 Cache Handle": "0x007F",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x001A",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #22",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x0080",
					"L2 Cache Handle": "0x0081",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x001B",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #23",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x0082",
					"L2 Cache Handle": "0x0083",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x001C",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #24",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x0084",
					"L2 Cache Handle": "0x0085",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x001D",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #25",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x0086",
					"L2 Cache Handle": "0x0087",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x001E",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #26",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x0088",
					"L2 Cache Handle": "0x0089",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x001F",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #27",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x008A",
					"L2 Cache Handle": "0x008B",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x0020",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #28",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x008C",
					"L2 Cache Handle": "0x008D",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x0021",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #29",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x008E",
					"L2 Cache Handle": "0x008F",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x0022",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #30",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x0090",
					"L2 Cache Handle": "0x0091",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x0023",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #31",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x0092",
					"L2 Cache Handle": "0x0093",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x0024",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #32",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x0094",
					"L2 Cache Handle": "0x0095",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x0025",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #33",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x0096",
					"L2 Cache Handle": "0x0097",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x0026",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #34",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x0098",
					"L2 Cache Handle": "0x0099",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x0027",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #35",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x009A",
					"L2 Cache Handle": "0x009B",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x0028",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #36",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x009C",
					"L2 Cache Handle": "0x009D",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x0029",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #37",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x009E",
					"L2 Cache Handle": "0x009F",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x002A",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #38",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x00A0",
					"L2 Cache Handle": "0x00A1",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x002B",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #39",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x00A2",
					"L2 Cache Handle": "0x00A3",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x002C",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #40",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x00A4",
					"L2 Cache Handle": "0x00A5",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x002D",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #41",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x00A6",
					"L2 Cache Handle": "0x00A7",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x002E",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #42",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x00A8",
					"L2 Cache Handle": "0x00A9",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x002F",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #43",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x00AA",
					"L2 Cache Handle": "0x00AB",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x0030",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #44",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x00AC",
					"L2 Cache Handle": "0x00AD",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x0031",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #45",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x00AE",
					"L2 Cache Handle": "0x00AF",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x0032",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #46",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x00B0",
					"L2 Cache Handle": "0x00B1",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x0033",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #47",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x00B2",
					"L2 Cache Handle": "0x00B3",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x0034",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #48",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x00B4",
					"L2 Cache Handle": "0x00B5",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x0035",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #49",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x00B6",
					"L2 Cache Handle": "0x00B7",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x0036",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #50",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x00B8",
					"L2 Cache Handle": "0x00B9",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x0037",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #51",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x00BA",
					"L2 Cache Handle": "0x00BB",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x0038",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #52",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x00BC",
					"L2 Cache Handle": "0x00BD",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x0039",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #53",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x00BE",
					"L2 Cache Handle": "0x00BF",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x003A",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #54",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x00C0",
					"L2 Cache Handle": "0x00C1",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x003B",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #55",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x00C2",
					"L2 Cache Handle": "0x00C3",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x003C",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #56",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x00C4",
					"L2 Cache Handle": "0x00C5",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x003D",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #57",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x00C6",
					"L2 Cache Handle": "0x00C7",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x003E",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #58",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x00C8",
					"L2 Cache Handle": "0x00C9",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x003F",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #59",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x00CA",
					"L2 Cache Handle": "0x00CB",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x0040",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #60",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x00CC",
					"L2 Cache Handle": "0x00CD",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x0041",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #61",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x00CE",
					"L2 Cache Handle": "0x00CF",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x0042",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #62",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x00D0",
					"L2 Cache Handle": "0x00D1",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}, {
					"record_id": "0x0043",
					"size": "4",
					"application_identifier": "Processor Information",
					"Socket Designation": "CPU socket #63",
					"Type": "Central Processor",
					"Family": "Unknown",
					"Manufacturer": "GenuineIntel",
					"ID": "61 06 00 00 FF FB AB 0F",
					"Version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
					"Voltage": "3.3 V",
					"External Clock": "Unknown",
					"Max Speed": "30000 MHz",
					"Current Speed": "2800 MHz",
					"Status": "Populated, Disabled By BIOS",
					"Upgrade": "ZIF Socket",
					"L1 Cache Handle": "0x00D2",
					"L2 Cache Handle": "0x00D3",
					"L3 Cache Handle": "Not Provided",
					"Serial Number": "Not Specified",
					"Asset Tag": "Not Specified",
					"Part Number": "Not Specified"
				}],
				"type": "Central Processor",
				"family": "Unknown",
				"manufacturer": "GenuineIntel",
				"version": "Intel(R) Core(TM) i7-4980HQ CPU @ 2.80GHz",
				"voltage": "3.3 V",
				"external_clock": "Unknown",
				"max_speed": "30000 MHz",
				"current_speed": "2800 MHz",
				"upgrade": "ZIF Socket",
				"l3_cache_handle": "Not Provided",
				"serial_number": "Not Specified",
				"asset_tag": "Not Specified",
				"part_number": "Not Specified"
			},
			"memory_module": {
				"all_records": [{
					"record_id": "0x0045",
					"size": "6",
					"application_identifier": "Memory Module Information",
					"Socket Designation": "RAM socket #0",
					"Bank Connections": "None",
					"Current Speed": "Unknown",
					"Type": "EDO DIMM",
					"Installed Size": "2048 MB (Single-bank Connection)",
					"Enabled Size": "2048 MB (Single-bank Connection)",
					"Error Status": "OK"
				}, {
					"record_id": "0x0046",
					"size": "6",
					"application_identifier": "Memory Module Information",
					"Socket Designation": "RAM socket #1",
					"Bank Connections": "None",
					"Current Speed": "Unknown",
					"Type": "DIMM",
					"Installed Size": "Not Installed",
					"Enabled Size": "Not Installed",
					"Error Status": "OK"
				}, {
					"record_id": "0x0047",
					"size": "6",
					"application_identifier": "Memory Module Information",
					"Socket Designation": "RAM socket #2",
					"Bank Connections": "None",
					"Current Speed": "Unknown",
					"Type": "DIMM",
					"Installed Size": "Not Installed",
					"Enabled Size": "Not Installed",
					"Error Status": "OK"
				}, {
					"record_id": "0x0048",
					"size": "6",
					"application_identifier": "Memory Module Information",
					"Socket Designation": "RAM socket #3",
					"Bank Connections": "None",
					"Current Speed": "Unknown",
					"Type": "DIMM",
					"Installed Size": "Not Installed",
					"Enabled Size": "Not Installed",
					"Error Status": "OK"
				}, {
					"record_id": "0x0049",
					"size": "6",
					"application_identifier": "Memory Module Information",
					"Socket Designation": "RAM socket #4",
					"Bank Connections": "None",
					"Current Speed": "Unknown",
					"Type": "DIMM",
					"Installed Size": "Not Installed",
					"Enabled Size": "Not Installed",
					"Error Status": "OK"
				}, {
					"record_id": "0x004A",
					"size": "6",
					"application_identifier": "Memory Module Information",
					"Socket Designation": "RAM socket #5",
					"Bank Connections": "None",
					"Current Speed": "Unknown",
					"Type": "DIMM",
					"Installed Size": "Not Installed",
					"Enabled Size": "Not Installed",
					"Error Status": "OK"
				}, {
					"record_id": "0x004B",
					"size": "6",
					"application_identifier": "Memory Module Information",
					"Socket Designation": "RAM socket #6",
					"Bank Connections": "None",
					"Current Speed": "Unknown",
					"Type": "DIMM",
					"Installed Size": "Not Installed",
					"Enabled Size": "Not Installed",
					"Error Status": "OK"
				}, {
					"record_id": "0x004C",
					"size": "6",
					"application_identifier": "Memory Module Information",
					"Socket Designation": "RAM socket #7",
					"Bank Connections": "None",
					"Current Speed": "Unknown",
					"Type": "DIMM",
					"Installed Size": "Not Installed",
					"Enabled Size": "Not Installed",
					"Error Status": "OK"
				}, {
					"record_id": "0x004D",
					"size": "6",
					"application_identifier": "Memory Module Information",
					"Socket Designation": "RAM socket #8",
					"Bank Connections": "None",
					"Current Speed": "Unknown",
					"Type": "DIMM",
					"Installed Size": "Not Installed",
					"Enabled Size": "Not Installed",
					"Error Status": "OK"
				}, {
					"record_id": "0x004E",
					"size": "6",
					"application_identifier": "Memory Module Information",
					"Socket Designation": "RAM socket #9",
					"Bank Connections": "None",
					"Current Speed": "Unknown",
					"Type": "DIMM",
					"Installed Size": "Not Installed",
					"Enabled Size": "Not Installed",
					"Error Status": "OK"
				}, {
					"record_id": "0x004F",
					"size": "6",
					"application_identifier": "Memory Module Information",
					"Socket Designation": "RAM socket #10",
					"Bank Connections": "None",
					"Current Speed": "Unknown",
					"Type": "DIMM",
					"Installed Size": "Not Installed",
					"Enabled Size": "Not Installed",
					"Error Status": "OK"
				}, {
					"record_id": "0x0050",
					"size": "6",
					"application_identifier": "Memory Module Information",
					"Socket Designation": "RAM socket #11",
					"Bank Connections": "None",
					"Current Speed": "Unknown",
					"Type": "DIMM",
					"Installed Size": "Not Installed",
					"Enabled Size": "Not Installed",
					"Error Status": "OK"
				}, {
					"record_id": "0x0051",
					"size": "6",
					"application_identifier": "Memory Module Information",
					"Socket Designation": "RAM socket #12",
					"Bank Connections": "None",
					"Current Speed": "Unknown",
					"Type": "DIMM",
					"Installed Size": "Not Installed",
					"Enabled Size": "Not Installed",
					"Error Status": "OK"
				}, {
					"record_id": "0x0052",
					"size": "6",
					"application_identifier": "Memory Module Information",
					"Socket Designation": "RAM socket #13",
					"Bank Connections": "None",
					"Current Speed": "Unknown",
					"Type": "DIMM",
					"Installed Size": "Not Installed",
					"Enabled Size": "Not Installed",
					"Error Status": "OK"
				}, {
					"record_id": "0x0053",
					"size": "6",
					"application_identifier": "Memory Module Information",
					"Socket Designation": "RAM socket #14",
					"Bank Connections": "None",
					"Current Speed": "Unknown",
					"Type": "DIMM",
					"Installed Size": "Not Installed",
					"Enabled Size": "Not Installed",
					"Error Status": "OK"
				}],
				"bank_connections": "None",
				"current_speed": "Unknown",
				"error_status": "OK"
			},
			"oem_strings": {
				"all_records": [{
					"record_id": "0x00E0",
					"size": "11",
					"application_identifier": "OEM Strings",
					"String 1": "[MS_VM_CERT/SHA1/27d66596a61c48dd3dc7216fd715126e33f59ae7]",
					"String 2": "Welcome to the Virtual Machine"
				}],
				"string_1": "[MS_VM_CERT/SHA1/27d66596a61c48dd3dc7216fd715126e33f59ae7]",
				"string_2": "Welcome to the Virtual Machine"
			}
		},
		"ohai_time": 1435875334.1333106,
		"init_package": "init",
		"uptime_seconds": 1263,
		"uptime": "21 minutes 03 seconds",
		"idletime_seconds": 2494,
		"idletime": "41 minutes 34 seconds",
		"block_device": {
			"fd0": {
				"size": "0",
				"removable": "1",
				"rotational": "1"
			},
			"sda": {
				"size": "41943040",
				"removable": "0",
				"model": "VMware Virtual S",
				"rev": "1.0",
				"state": "running",
				"timeout": "30",
				"vendor": "VMware,",
				"rotational": "1"
			},
			"sr0": {
				"size": "2097151",
				"removable": "1",
				"model": "VMware SATA CD01",
				"rev": "1.00",
				"state": "running",
				"timeout": "30",
				"vendor": "NECVMWar",
				"rotational": "1"
			},
			"ram0": {
				"size": "131072",
				"removable": "0",
				"rotational": "1"
			},
			"ram1": {
				"size": "131072",
				"removable": "0",
				"rotational": "1"
			},
			"ram2": {
				"size": "131072",
				"removable": "0",
				"rotational": "1"
			},
			"ram3": {
				"size": "131072",
				"removable": "0",
				"rotational": "1"
			},
			"ram4": {
				"size": "131072",
				"removable": "0",
				"rotational": "1"
			},
			"ram5": {
				"size": "131072",
				"removable": "0",
				"rotational": "1"
			},
			"ram6": {
				"size": "131072",
				"removable": "0",
				"rotational": "1"
			},
			"ram7": {
				"size": "131072",
				"removable": "0",
				"rotational": "1"
			},
			"ram8": {
				"size": "131072",
				"removable": "0",
				"rotational": "1"
			},
			"ram9": {
				"size": "131072",
				"removable": "0",
				"rotational": "1"
			},
			"loop0": {
				"size": "0",
				"removable": "0",
				"rotational": "1"
			},
			"loop1": {
				"size": "0",
				"removable": "0",
				"rotational": "1"
			},
			"loop2": {
				"size": "0",
				"removable": "0",
				"rotational": "1"
			},
			"loop3": {
				"size": "0",
				"removable": "0",
				"rotational": "1"
			},
			"loop4": {
				"size": "0",
				"removable": "0",
				"rotational": "1"
			},
			"loop5": {
				"size": "0",
				"removable": "0",
				"rotational": "1"
			},
			"loop6": {
				"size": "0",
				"removable": "0",
				"rotational": "1"
			},
			"loop7": {
				"size": "0",
				"removable": "0",
				"rotational": "1"
			},
			"ram10": {
				"size": "131072",
				"removable": "0",
				"rotational": "1"
			},
			"ram11": {
				"size": "131072",
				"removable": "0",
				"rotational": "1"
			},
			"ram12": {
				"size": "131072",
				"removable": "0",
				"rotational": "1"
			},
			"ram13": {
				"size": "131072",
				"removable": "0",
				"rotational": "1"
			},
			"ram14": {
				"size": "131072",
				"removable": "0",
				"rotational": "1"
			},
			"ram15": {
				"size": "131072",
				"removable": "0",
				"rotational": "1"
			}
		},
		"chef_packages": {
			"ohai": {
				"version": "8.5.0",
				"ohai_root": "/opt/chef/embedded/lib/ruby/gems/2.1.0/gems/ohai-8.5.0/lib/ohai"
			},
			"chef": {
				"version": "12.4.0",
				"chef_root": "/opt/chef/embedded/apps/chef/lib"
			}
		},
		"command": {
			"ps": "ps -ef"
		},
		"etc": {
			"passwd": {
				"root": {
					"dir": "/root",
					"gid": 0,
					"uid": 0,
					"shell": "/bin/bash",
					"gecos": "root"
				},
				"daemon": {
					"dir": "/usr/sbin",
					"gid": 1,
					"uid": 1,
					"shell": "/usr/sbin/nologin",
					"gecos": "daemon"
				},
				"bin": {
					"dir": "/bin",
					"gid": 2,
					"uid": 2,
					"shell": "/usr/sbin/nologin",
					"gecos": "bin"
				},
				"sys": {
					"dir": "/dev",
					"gid": 3,
					"uid": 3,
					"shell": "/usr/sbin/nologin",
					"gecos": "sys"
				},
				"sync": {
					"dir": "/bin",
					"gid": 65534,
					"uid": 4,
					"shell": "/bin/sync",
					"gecos": "sync"
				},
				"games": {
					"dir": "/usr/games",
					"gid": 60,
					"uid": 5,
					"shell": "/usr/sbin/nologin",
					"gecos": "games"
				},
				"man": {
					"dir": "/var/cache/man",
					"gid": 12,
					"uid": 6,
					"shell": "/usr/sbin/nologin",
					"gecos": "man"
				},
				"lp": {
					"dir": "/var/spool/lpd",
					"gid": 7,
					"uid": 7,
					"shell": "/usr/sbin/nologin",
					"gecos": "lp"
				},
				"mail": {
					"dir": "/var/mail",
					"gid": 8,
					"uid": 8,
					"shell": "/usr/sbin/nologin",
					"gecos": "mail"
				},
				"news": {
					"dir": "/var/spool/news",
					"gid": 9,
					"uid": 9,
					"shell": "/usr/sbin/nologin",
					"gecos": "news"
				},
				"uucp": {
					"dir": "/var/spool/uucp",
					"gid": 10,
					"uid": 10,
					"shell": "/usr/sbin/nologin",
					"gecos": "uucp"
				},
				"proxy": {
					"dir": "/bin",
					"gid": 13,
					"uid": 13,
					"shell": "/usr/sbin/nologin",
					"gecos": "proxy"
				},
				"www-data": {
					"dir": "/var/www",
					"gid": 33,
					"uid": 33,
					"shell": "/usr/sbin/nologin",
					"gecos": "www-data"
				},
				"backup": {
					"dir": "/var/backups",
					"gid": 34,
					"uid": 34,
					"shell": "/usr/sbin/nologin",
					"gecos": "backup"
				},
				"list": {
					"dir": "/var/list",
					"gid": 38,
					"uid": 38,
					"shell": "/usr/sbin/nologin",
					"gecos": "Mailing List Manager"
				},
				"irc": {
					"dir": "/var/run/ircd",
					"gid": 39,
					"uid": 39,
					"shell": "/usr/sbin/nologin",
					"gecos": "ircd"
				},
				"gnats": {
					"dir": "/var/lib/gnats",
					"gid": 41,
					"uid": 41,
					"shell": "/usr/sbin/nologin",
					"gecos": "Gnats Bug-Reporting System (admin)"
				},
				"nobody": {
					"dir": "/nonexistent",
					"gid": 65534,
					"uid": 65534,
					"shell": "/usr/sbin/nologin",
					"gecos": "nobody"
				},
				"libuuid": {
					"dir": "/var/lib/libuuid",
					"gid": 101,
					"uid": 100,
					"shell": "",
					"gecos": ""
				},
				"syslog": {
					"dir": "/home/syslog",
					"gid": 104,
					"uid": 101,
					"shell": "/bin/false",
					"gecos": ""
				},
				"messagebus": {
					"dir": "/var/run/dbus",
					"gid": 106,
					"uid": 102,
					"shell": "/bin/false",
					"gecos": ""
				},
				"landscape": {
					"dir": "/var/lib/landscape",
					"gid": 109,
					"uid": 103,
					"shell": "/bin/false",
					"gecos": ""
				},
				"sshd": {
					"dir": "/var/run/sshd",
					"gid": 65534,
					"uid": 104,
					"shell": "/usr/sbin/nologin",
					"gecos": ""
				},
				"paulmooring": {
					"dir": "/home/paulmooring",
					"gid": 1000,
					"uid": 1000,
					"shell": "/bin/bash",
					"gecos": "Paul Mooring,,,"
				},
				"ntp": {
					"dir": "/home/ntp",
					"gid": 112,
					"uid": 105,
					"shell": "/bin/false",
					"gecos": ""
				},
				"pwm": {
					"dir": "/home/pwm",
					"gid": 1001,
					"uid": 1001,
					"shell": "/bin/bash",
					"gecos": ""
				}
			},
			"group": {
				"root": {
					"gid": 0,
					"members": []
				},
				"daemon": {
					"gid": 1,
					"members": []
				},
				"bin": {
					"gid": 2,
					"members": []
				},
				"sys": {
					"gid": 3,
					"members": []
				},
				"adm": {
					"gid": 4,
					"members": ["syslog", "paulmooring"]
				},
				"tty": {
					"gid": 5,
					"members": []
				},
				"disk": {
					"gid": 6,
					"members": []
				},
				"lp": {
					"gid": 7,
					"members": []
				},
				"mail": {
					"gid": 8,
					"members": []
				},
				"news": {
					"gid": 9,
					"members": []
				},
				"uucp": {
					"gid": 10,
					"members": []
				},
				"man": {
					"gid": 12,
					"members": []
				},
				"proxy": {
					"gid": 13,
					"members": []
				},
				"kmem": {
					"gid": 15,
					"members": []
				},
				"dialout": {
					"gid": 20,
					"members": []
				},
				"fax": {
					"gid": 21,
					"members": []
				},
				"voice": {
					"gid": 22,
					"members": []
				},
				"cdrom": {
					"gid": 24,
					"members": ["paulmooring"]
				},
				"floppy": {
					"gid": 25,
					"members": []
				},
				"tape": {
					"gid": 26,
					"members": []
				},
				"sudo": {
					"gid": 27,
					"members": ["paulmooring", "pwm"]
				},
				"audio": {
					"gid": 29,
					"members": []
				},
				"dip": {
					"gid": 30,
					"members": ["paulmooring"]
				},
				"www-data": {
					"gid": 33,
					"members": []
				},
				"backup": {
					"gid": 34,
					"members": []
				},
				"operator": {
					"gid": 37,
					"members": []
				},
				"list": {
					"gid": 38,
					"members": []
				},
				"irc": {
					"gid": 39,
					"members": []
				},
				"src": {
					"gid": 40,
					"members": []
				},
				"gnats": {
					"gid": 41,
					"members": []
				},
				"shadow": {
					"gid": 42,
					"members": []
				},
				"utmp": {
					"gid": 43,
					"members": []
				},
				"video": {
					"gid": 44,
					"members": []
				},
				"sasl": {
					"gid": 45,
					"members": []
				},
				"plugdev": {
					"gid": 46,
					"members": ["paulmooring"]
				},
				"staff": {
					"gid": 50,
					"members": []
				},
				"games": {
					"gid": 60,
					"members": []
				},
				"users": {
					"gid": 100,
					"members": []
				},
				"nogroup": {
					"gid": 65534,
					"members": []
				},
				"libuuid": {
					"gid": 101,
					"members": []
				},
				"netdev": {
					"gid": 102,
					"members": []
				},
				"crontab": {
					"gid": 103,
					"members": []
				},
				"syslog": {
					"gid": 104,
					"members": []
				},
				"fuse": {
					"gid": 105,
					"members": []
				},
				"messagebus": {
					"gid": 106,
					"members": []
				},
				"mlocate": {
					"gid": 107,
					"members": []
				},
				"ssh": {
					"gid": 108,
					"members": []
				},
				"landscape": {
					"gid": 109,
					"members": []
				},
				"paulmooring": {
					"gid": 1000,
					"members": []
				},
				"lpadmin": {
					"gid": 110,
					"members": ["paulmooring"]
				},
				"sambashare": {
					"gid": 111,
					"members": ["paulmooring"]
				},
				"ntp": {
					"gid": 112,
					"members": []
				},
				"pwm": {
					"gid": 1001,
					"members": []
				}
			}
		},
		"current_user": "root",
		"recipes": ["packages::default", "ntp::default", "users::sysadmins", "sudo::default", "ntp::apparmor"],
		"expanded_run_list": ["packages::default", "ntp::default", "users::sysadmins", "sudo::default"],
		"roles": ["linux-base"],
		"cookbooks": {
			"ntp": {
				"version": "1.8.6"
			},
			"packages": {
				"version": "0.1.0"
			},
			"sudo": {
				"version": "2.7.1"
			},
			"users": {
				"version": "1.8.2"
			}
		}
	},
  "normal": {
    "tags": [
      "chef-swarm"
    ],
    "timestamp": ` + timestamp() + `
    "runit": {
      "sv_bin": "/usr/bin/sv",
      "service_dir": "/etc/service",
      "sv_dir": "/etc/sv"
    },
    "apache": {
      "dir": "/etc/apache2",
      "log_dir": "/var/log/apache2",
      "user": "www-data",
      "group": "www-data",
      "binary": "/usr/sbin/apache2",
      "icondir": "/usr/share/apache2/icons",
      "cache_dir": "/var/cache/apache2",
      "pid_file": "/var/run/apache2.pid",
      "lib_dir": "/usr/lib/apache2",
      "libexecdir": "/usr/lib/apache2/modules"
    },
    "ec2opts": {
      "lvm": {
        "use_ephemeral": true,
        "ephemeral_mountpoint": "/mnt",
        "ephemeral_volume_group": "ephemeral",
        "ephemeral_logical_volume": "store",
        "ephemeral_devices": {
          "m1.small": [
            "/dev/sda2"
          ],
          "m1.large": [
            "/dev/sdb",
            "/dev/sdc"
          ],
          "m1.xlarge": [
            "/dev/sdb",
            "/dev/sdc",
            "/dev/sdd",
            "/dev/sde"
          ]
        }
      }
    },
    "nagios": {
      "plugin_dir": "/usr/lib/nagios/plugins",
      "client": {
        "install_method": "package"
      },
      "nrpe": {
        "home": "/usr/lib/nagios",
        "conf_dir": "/etc/nagios",
        "init_style": "runit"
      },
      "server": {
        "install_method": "package",
        "service_name": "nagios3"
      },
      "home": "/usr/lib/nagios3",
      "conf_dir": "/etc/nagios3",
      "config_dir": "/etc/nagios3/conf.d",
      "log_dir": "/var/log/nagios3",
      "apachelog_dir": "/var/log/apache2",
      "cache_dir": "/var/cache/nagios3",
      "state_dir": "/var/lib/nagios3",
      "run_dir": "/var/run/nagios3",
      "docroot": "/usr/share/nagios3/htdocs"
    },
    "mysql": {
      "server_id": "1",
      "conf_dir": "/etc/mysql",
      "socket": "/var/run/mysqld/mysqld.sock",
      "pid_file": "/var/run/mysqld/mysqld.pid",
      "old_passwords": 0
    },
    "src_dir": "/usr/local/src",
    "db_src_dir": "/usr/local/db_src",
    "postgresql": {
      "config": {
        "listen_addresses": "*",
        "log_lock_waits": "on",
        "log_min_duration_statement": 1000,
        "log_line_prefix": "%u %d %p %m %c ",
        "log_temp_files": 0,
        "log_checkpoints": "on"
      },
      "enable_pitti_ppa": true,
      "server": {
        "config": false
      },
      "pg_hba": [
        {
          "type": "local",
          "db": "all",
          "user": "postgres",
          "addr": null,
          "method": "ident"
        },
        {
          "type": "local",
          "db": "all",
          "user": "all",
          "addr": null,
          "method": "ident"
        },
        {
          "type": "host",
          "db": "all",
          "user": "all",
          "addr": "127.0.0.1/32",
          "method": "md5"
        },
        {
          "type": "host",
          "db": "all",
          "user": "all",
          "addr": "::1/128",
          "method": "md5"
        },
        {
          "type": "host",
          "db": "all",
          "user": "all",
          "addr": "0.0.0.0/0",
          "method": "md5"
        },
        {
          "type": "host",
          "db": "all",
          "user": "all",
          "addr": "::0/0",
          "method": "md5"
        }
      ]
    },
    "build_essential": {
      "compiletime": true
    },
    "chef_client": {
      "bin": "/usr/bin/chef-client"
    }
  }
}`

	return string_to_reader(body)
}
