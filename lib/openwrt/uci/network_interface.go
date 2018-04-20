package uci

func init() {
	registerConfigType(ConfigNetwork, "interface", func() ConfigData { return &NetworkInterface{} })
}

// NetworkInterface represents a network interface
type NetworkInterface struct {
	Type           string `json:"type,omitempty"`
	IfName         string `json:"ifname,omitempty"`
	MacAddr        string `json:"macaddr,omitempty"`
	Proto          string `json:"proto,omitempty"`
	IPAddr         string `json:"ipaddr,omitempty"`
	Netmask        string `json:"netmask,omitempty"`
	Gateway        string `json:"gateway,omitempty"`
	Metric         string `json:"metric,omitempty"`
	IP4Table       string `json:"ip4table,omitempty"`
	TxQueueLen     string `json:"txqueuelen,omitempty"`
	Multipath      string `json:"multipath,omitempty"`
	MTU            string `json:"mtu,omitempty"`
	TrafficControl string `json:"trafficcontrol,omitempty"`
	Upload         string `json:"upload,omitempty"`
	Download       string `json:"download,omitempty"`
	APN            string `json:"apn,omitempty"`
	Label          string `json:"label,omitempty"`
	Device         string `json:"device,omitempty"`
	PeerDNS        string `json:"peerdns,omitempty"`
	IP6Table       string `json:"ip6table,omitempty"`
	Service        string `json:"service,omitempty"`
}

// ConfigType implements the ConfigData interface
func (ni *NetworkInterface) ConfigType() ConfigType { return ConfigNetwork }

// ConfigSubType implements the ConfigData interface
func (ni *NetworkInterface) ConfigSubType() string { return "interface" }

// ConfigName implements the ConfigData interface
func (ni *NetworkInterface) ConfigName() string { return ni.IfName }
