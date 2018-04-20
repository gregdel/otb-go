package uci

func init() {
	registerConfigType(ConfigDHCP, "dhcp", func() ConfigData { return &DHCPConfig{} })
}

// DHCPConfig represents a DHCP configuration
type DHCPConfig struct {
	Interface      string `json:"interface,omitempty"`
	Start          string `json:"start,omitempty"`
	Limit          string `json:"limit,omitempty"`
	LeaseTime      string `json:"leasetime,omitempty"`
	Dynamic        string `json:"dynamicdhcp,omitempty"`
	Force          string `json:"force,omitempty"`
	Ignore         string `json:"ignore,omitempty"`
	InterfaceNames List   `json:"interface_name,omitempty"`
}

// ConfigType implements the ConfigData interface
func (dd *DHCPConfig) ConfigType() ConfigType { return ConfigDHCP }

// ConfigSubType implements the ConfigData interface
func (dd *DHCPConfig) ConfigSubType() string { return "dhcp" }

// ConfigName implements the ConfigData interface
func (dd *DHCPConfig) ConfigName() string { return dd.Interface }
