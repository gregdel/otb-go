package uci

func init() {
	registerConfigType(ConfigFirewall, "include", func() ConfigData { return &FirewallInclude{} })
}

// FirewallInclude represents a firewall include type
type FirewallInclude struct {
	Path string `json:"path,omitempty"`
}

// ConfigType implements the ConfigData interface
func (fi *FirewallInclude) ConfigType() ConfigType { return ConfigFirewall }

// ConfigSubType implements the ConfigData interface
func (fi *FirewallInclude) ConfigSubType() string { return "include" }

// ConfigName implements the ConfigData interface
func (fi *FirewallInclude) ConfigName() string { return "" }
