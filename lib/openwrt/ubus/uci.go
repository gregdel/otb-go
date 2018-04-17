package ubus

import "github.com/gregdel/otb-go/lib/openwrt/uci"

// UciGet gets the configuration in uci
func UciGet(result interface{}, config string) error {
	params := map[string]string{"config": config}

	return callWithParams(result, params, "uci", "get")
}

// UciGetConfig returns a raw configuration
func UciGetConfig(t uci.ConfigType) (*uci.Configuration, error) {
	c := uci.NewConfiguration(t)
	return c, UciGet(c, string(t))
}
