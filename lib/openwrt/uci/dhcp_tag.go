package uci

func init() {
	registerConfigType(ConfigDHCP, "tag", func() ConfigData { return &DHCPTag{} })
}

// DHCPTag represents a DHCP tag configuration
type DHCPTag struct {
	Options List `json:"dhcp_option,omitempty"`
}

// ConfigType implements the ConfigData interface
func (dt *DHCPTag) ConfigType() ConfigType { return ConfigDHCP }

// ConfigSubType implements the ConfigData interface
func (dt *DHCPTag) ConfigSubType() string { return "tag" }

// ConfigName implements the ConfigData interface
func (dt *DHCPTag) ConfigName() string { return "" }
