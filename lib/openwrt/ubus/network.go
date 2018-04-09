package ubus

// NetworkInterface represents a network interface
type NetworkInterface struct {
	Interface     string        `json:"interface"`
	Up            bool          `json:"up"`
	Pending       bool          `json:"pending"`
	Avalable      bool          `json:"avalable"`
	AutoStart     bool          `json:"autostart"`
	Dynamic       bool          `json:"dynamic"`
	Uptime        int           `json:"uptime"`
	L3Device      string        `json:"l3_device"`
	Proto         string        `json:"proto"`
	Device        string        `json:"device"`
	Metric        int           `json:"metric"`
	IP4Table      int           `json:"ip4table"`
	DNSMetric     int           `json:"dns_metric"`
	Delegation    bool          `json:"delegation"`
	IPv4Addresses []IPv4Address `json:"ipv4-address"`
	Routes        []Route       `json:"route"`
	DNSServer     []string      `json:"dns-server"`
}

// IPv4Address represents an IPv4 address
type IPv4Address struct {
	Address string `json:"address"`
	Mask    int    `json:"mask"`
}

// Route represents a route
type Route struct {
	Target  string `json:"target"`
	Mask    int    `json:"mask"`
	Nexthop string `json:"nexthop"`
	Source  string `json:"source"`
}

// NetworkDump dumps the network configuraion seen by UCI
func NetworkDump() ([]*NetworkInterface, error) {
	out := struct {
		Interfaces []*NetworkInterface `json:"interface"`
	}{
		Interfaces: []*NetworkInterface{},
	}

	if err := call(&out, "network.interface", "dump"); err != nil {
		return nil, err
	}

	return out.Interfaces, nil
}
