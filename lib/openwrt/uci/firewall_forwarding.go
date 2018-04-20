package uci

func init() {
	registerConfigType(ConfigFirewall, "forwarding", func() ConfigData { return &FirewallForwarding{} })
}

// FirewallForwarding represents a firewall forwarding configuration
type FirewallForwarding struct {
	Src  string `json:"src,omitempty"`
	Dest string `json:"dest,omitempty"`
}

// ConfigType implements the ConfigData interface
func (fd *FirewallForwarding) ConfigType() ConfigType { return ConfigFirewall }

// ConfigSubType implements the ConfigData interface
func (fd *FirewallForwarding) ConfigSubType() string { return "forwarding" }

// ConfigName implements the ConfigData interface
func (fd *FirewallForwarding) ConfigName() string { return "" }
