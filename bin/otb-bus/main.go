package main

import (
	"fmt"

	"github.com/gregdel/otb-go/lib/openwrt/ubus"
	"github.com/kr/pretty"
)

func main() {
	info, err := ubus.SystemBoard()
	if err != nil {
		fmt.Println("failed to get system info: ", err)
		return
	}
	pretty.Println(info)
}
