package uci

func init() {
	registerConfigType(ConfigNetwork, "device", func() ConfigData { return &NetworkDevice{} })
}

// NetworkDevice represents a network device
type NetworkDevice struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	IfName  string `json:"ifname"`
	VID     string `json:"vid"`
	MacAddr string `json:"macaddr"`
}

// ConfigType implements the ConfigData interface
func (nd *NetworkDevice) ConfigType() ConfigType { return ConfigNetwork }

// ConfigSubType implements the ConfigData interface
func (nd *NetworkDevice) ConfigSubType() string { return "device" }
