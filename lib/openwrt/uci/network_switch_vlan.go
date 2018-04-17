package uci

func init() {
	registerConfigType(ConfigNetwork, "switch_vlan", func() ConfigData { return &NetworkSwitchVLAN{} })
}

// NetworkSwitchVLAN represents network VLANswitch
type NetworkSwitchVLAN struct {
	Device string `json:"device"`
	VLAN   string `json:"vlan"`
	Ports  string `json:"ports"`
}

// ConfigType implements the ConfigData interface
func (nr *NetworkSwitchVLAN) ConfigType() ConfigType { return ConfigNetwork }

// ConfigSubType implements the ConfigData interface
func (nr *NetworkSwitchVLAN) ConfigSubType() string { return "switch_vlan" }
