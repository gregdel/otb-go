package main

import (
	"log"

	"github.com/gregdel/otb-go/lib/openwrt/ubus"
	"github.com/gregdel/otb-go/lib/openwrt/uci"
	"github.com/kr/pretty"
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

	configuration, err := ubus.UciGetConfig(uci.ConfigNetwork)
	if err != nil {
		log.Fatalf("failed to get network configuration: %s", err)
	}
	pretty.Println(configuration)
}
