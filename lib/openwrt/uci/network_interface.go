package uci

func init() {
	registerConfigType(ConfigNetwork, "interface", func() ConfigData { return &NetworkInterface{} })
}

// NetworkInterface represents a network interface
type NetworkInterface struct {
	Type           string `json:"type"`
	IfName         string `json:"ifname"`
	MacAddr        string `json:"macaddr"`
	Proto          string `json:"proto"`
	IPAddr         string `json:"ipaddr"`
	Netmask        string `json:"netmask"`
	Gateway        string `json:"gateway"`
	Metric         string `json:"metric"`
	IP4Table       string `json:"ip4table"`
	TxQueueLen     string `json:"txqueuelen"`
	Multipath      string `json:"multipath"`
	MTU            string `json:"mtu"`
	TrafficControl string `json:"trafficcontrol"`
	Upload         string `json:"upload"`
	Download       string `json:"download"`
	APN            string `json:"apn"`
	Label          string `json:"label"`
	Device         string `json:"device"`
	PeerDNS        string `json:"peerdns"`
	IP6Table       string `json:"ip6table"`
	Service        string `json:"service"`
}

// ConfigType implements the ConfigData interface
func (ni *NetworkInterface) ConfigType() ConfigType { return ConfigNetwork }

// ConfigSubType implements the ConfigData interface
func (ni *NetworkInterface) ConfigSubType() string { return "interface" }
