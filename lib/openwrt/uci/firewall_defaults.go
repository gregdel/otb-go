package uci

func init() {
	registerConfigType(ConfigFirewall, "defaults", func() ConfigData { return &FirewallDefaults{} })
}

// FirewallDefaults represents a firewall default configuration
type FirewallDefaults struct {
	SynFlood string `json:"syn_flood,omitempty"`
	Input    string `json:"input,omitempty"`
	Output   string `json:"output,omitempty"`
	Forward  string `json:"forward,omitempty"`
}

// ConfigType implements the ConfigData interface
func (fd *FirewallDefaults) ConfigType() ConfigType { return ConfigFirewall }

// ConfigSubType implements the ConfigData interface
func (fd *FirewallDefaults) ConfigSubType() string { return "defaults" }

// ConfigName implements the ConfigData interface
func (fd *FirewallDefaults) ConfigName() string { return "" }
