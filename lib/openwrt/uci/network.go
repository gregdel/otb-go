package uci

// NetworkInterface represents a network interface
type NetworkInterface struct {
	Type       string `json:"type"`
	IfName     string `json:"ifname"`
	MacAddr    string `json:"macaddr"`
	Proto      string `json:"proto"`
	IPAddr     string `json:"ipaddr"`
	Netmask    string `json:"netmask"`
	Gateway    string `json:"gateway"`
	Metric     string `json:"metric"`
	IP4Table   string `json:"ip4table"`
	TxQueueLen string `json:"txqueuelen"`
}

// ConfigType implements the configuration type
func (ni *NetworkInterface) ConfigType() ConfigType { return ConfigNetwork }

// ConfigSubType implements the configuration type
func (ni *NetworkInterface) ConfigSubType() string { return "interface" }
