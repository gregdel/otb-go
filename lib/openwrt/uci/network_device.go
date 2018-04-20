package uci

func init() {
	registerConfigType(ConfigNetwork, "device", func() ConfigData { return &NetworkDevice{} })
}

// NetworkDevice represents a network device
type NetworkDevice struct {
	Type    string `json:"type,omitempty"`
	Name    string `json:"name,omitempty"`
	IfName  string `json:"ifname,omitempty"`
	VID     string `json:"vid,omitempty"`
	MacAddr string `json:"macaddr,omitempty"`
}

// ConfigType implements the ConfigData interface
func (nd *NetworkDevice) ConfigType() ConfigType { return ConfigNetwork }

// ConfigSubType implements the ConfigData interface
func (nd *NetworkDevice) ConfigSubType() string { return "device" }

// ConfigName implements the ConfigData interface
func (nd *NetworkDevice) ConfigName() string {
	if nd.Name != "" {
		return nd.Name
	}
	return nd.IfName + "_dev"
}
