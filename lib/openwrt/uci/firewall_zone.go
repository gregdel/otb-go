package uci

func init() {
	registerConfigType(ConfigFirewall, "zone", func() ConfigData { return &FirewallZone{} })
}

// FirewallZone represents a firewall zone configuration
type FirewallZone struct {
	Name    string `json:"name,omitempty"`
	Input   string `json:"input,omitempty"`
	Output  string `json:"output,omitempty"`
	Forward string `json:"forward,omitempty"`
	Masq    string `json:"masq,omitempty"`
	MTUFix  string `json:"mtu_fix,omitempty"`
	Network List   `json:"network,omitempty"`
}

// ConfigType implements the ConfigData interface
func (fz *FirewallZone) ConfigType() ConfigType { return ConfigFirewall }

// ConfigSubType implements the ConfigData interface
func (fz *FirewallZone) ConfigSubType() string { return "zone" }

// ConfigName implements the ConfigData interface
func (fz *FirewallZone) ConfigName() string { return fz.Name }
