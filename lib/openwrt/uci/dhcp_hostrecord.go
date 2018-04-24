package uci

func init() {
	registerConfigType(ConfigDHCP, "hostrecord", func() ConfigData { return &DHCPHostRecord{} })
}

// DHCPHostRecord represents a FQDN entry
type DHCPHostRecord struct {
	Name string `json:"name,omitempty"`
	IP   string `json:"ip,omitempty"`
}

// ConfigType implements the ConfigData interface
func (dd *DHCPHostRecord) ConfigType() ConfigType { return ConfigDHCP }

// ConfigSubType implements the ConfigData interface
func (dd *DHCPHostRecord) ConfigSubType() string { return "hostrecord" }

// ConfigName implements the ConfigData interface
func (dd *DHCPHostRecord) ConfigName() string { return "" }
