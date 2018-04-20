package uci

func init() {
	registerConfigType(ConfigNetwork, "switch_vlan", func() ConfigData { return &NetworkSwitchVLAN{} })
}

// NetworkSwitchVLAN represents network VLANswitch
type NetworkSwitchVLAN struct {
	Device string `json:"device,omitempty"`
	VLAN   string `json:"vlan,omitempty"`
	Ports  string `json:"ports,omitempty"`
}

// ConfigType implements the ConfigData interface
func (nr *NetworkSwitchVLAN) ConfigType() ConfigType { return ConfigNetwork }

// ConfigSubType implements the ConfigData interface
func (nr *NetworkSwitchVLAN) ConfigSubType() string { return "switch_vlan" }

// ConfigName implements the ConfigData interface
func (nr *NetworkSwitchVLAN) ConfigName() string { return "" }
