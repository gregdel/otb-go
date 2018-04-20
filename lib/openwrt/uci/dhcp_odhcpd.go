package uci

func init() {
	registerConfigType(ConfigDHCP, "odhcpd", func() ConfigData { return &DHCPOdhcpd{} })
}

// DHCPOdhcpd represents a DHCP configuration
type DHCPOdhcpd struct {
	MainDHCP     string `json:"maindhcp,omitempty"`
	LeaseFile    string `json:"leasefile,omitempty"`
	LeaseTrigger string `json:"leasetrigger,omitempty"`
	LogLevel     string `json:"loglevel,omitempty"`
}

// ConfigType implements the ConfigData interface
func (do *DHCPOdhcpd) ConfigType() ConfigType { return ConfigDHCP }

// ConfigSubType implements the ConfigData interface
func (do *DHCPOdhcpd) ConfigSubType() string { return "odhcpd" }

// ConfigName implements the ConfigData interface
func (do *DHCPOdhcpd) ConfigName() string { return "odhcpd" }
