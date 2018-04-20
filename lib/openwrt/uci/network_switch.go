package uci

func init() {
	registerConfigType(ConfigNetwork, "switch", func() ConfigData { return &NetworkSwitch{} })
}

// NetworkSwitch represents a network switch
type NetworkSwitch struct {
	Name       string `json:"name,omitempty"`
	Reset      string `json:"reset,omitempty"`
	EnableVLAN string `json:"enable_vlan,omitempty"`
}

// ConfigType implements the ConfigData interface
func (nr *NetworkSwitch) ConfigType() ConfigType { return ConfigNetwork }

// ConfigSubType implements the ConfigData interface
func (nr *NetworkSwitch) ConfigSubType() string { return "switch" }

// ConfigName implements the ConfigData interface
func (nr *NetworkSwitch) ConfigName() string { return "" }
