package uci

func init() {
	registerConfigType(ConfigNetwork, "switch", func() ConfigData { return &NetworkSwitch{} })
}

// NetworkSwitch represents a network switch
type NetworkSwitch struct {
	Name       string `json:"name"`
	Reset      string `json:"reset"`
	EnableVLAN string `json:"enable_vlan"`
}

// ConfigType implements the ConfigData interface
func (nr *NetworkSwitch) ConfigType() ConfigType { return ConfigNetwork }

// ConfigSubType implements the ConfigData interface
func (nr *NetworkSwitch) ConfigSubType() string { return "switch" }
