package ubus

// UciGet gets the configuration in uci
func UciGet(config string) error {
	params := map[string]string{"config": config}

	return callWithParams(nil, params, "uci", "get")
}
