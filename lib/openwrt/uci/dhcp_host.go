package uci

func init() {
	registerConfigType(ConfigDHCP, "host", func() ConfigData { return &DHCPHost{} })
}

// DHCPHost represents a static lease
type DHCPHost struct {
	IP        string `json:"ip,omitempty"`
	Mac       string `json:"mac,omitempty"`
	HostID    string `json:"hostid,omitempty"`
	Name      string `json:"name,omitempty"`
	Tag       string `json:"tag,omitempty"`
	DNS       string `json:"dns,omitempty"`
	Broadcast string `json:"broadcast,omitempty"`
	LeaseTime string `json:"leasetime,omitempty"`
}

// ConfigType implements the ConfigData interface
func (dd *DHCPHost) ConfigType() ConfigType { return ConfigDHCP }

// ConfigSubType implements the ConfigData interface
func (dd *DHCPHost) ConfigSubType() string { return "host" }

// ConfigName implements the ConfigData interface
func (dd *DHCPHost) ConfigName() string { return "" }
