package main

import (
	"log"

	"github.com/gregdel/otb-go/lib/openwrt/ubus"
	"github.com/gregdel/otb-go/lib/openwrt/uci"
)

func main() {
	// info, err := ubus.SystemBoard()
	// if err != nil {
	// 	log.Fatalf("failed to get system info: %s", err)
	// }
	// pretty.Println(info)

	// links, err := netlink.LinkList()
	// if err != nil {
	// 	log.Fatalf("failed to get link list: %s", err)
	// }
	// pretty.Println(links)

	// interfaces, err := ubus.NetworkDump()
	// if err != nil {
	// 	log.Fatalf("failed to dump network: %s", err)
	// }
	// pretty.Println(interfaces)

	// configuration, err := ubus.UciGetConfig(uci.ConfigNetwork)
	// if err != nil {
	// 	log.Fatalf("failed to get network configuration: %s", err)
	// }
	// pretty.Println(configuration)

	cb := &uci.ConfigBlock{
		ConfigData: &uci.NetworkInterface{
			IfName:    "yolo0",
			Proto:     "static",
			IPAddr:    "10.10.10.10",
			Netmask:   "255.255.255.0",
			Gateway:   "10.10.10.11",
			Metric:    "150",
			IP4Table:  "150",
			Multipath: "off",
		},
	}

	if err := ubus.UciAdd(cb); err != nil {
		log.Fatalf("failed to add interface: %s", err)
	}

	if err := ubus.UciCommit(string(cb.ConfigType())); err != nil {
		log.Fatalf("failed to commit interface: %s", err)
	}
}
