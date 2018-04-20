package uci

func init() {
	registerConfigType(ConfigFirewall, "rule", func() ConfigData { return &FirewallRule{} })
}

// FirewallRule represents a firewall rule
type FirewallRule struct {
	Name      string `json:"name,omitempty"`
	Src       string `json:"src,omitempty"`
	SrcIP     string `json:"src_ip,omitempty"`
	SrcPort   string `json:"src_port,omitempty"`
	Dest      string `json:"dest,omitempty"`
	DestIP    string `json:"dest_ip,omitempty"`
	DestPort  string `json:"dest_port,omitempty"`
	Proto     string `json:"proto,omitempty"`
	Limit     string `json:"limit,omitempty"`
	Family    string `json:"family,omitempty"`
	Target    string `json:"target,omitempty"`
	ICMPTypes List   `json:"icmp_type,omitempty"`
}

// ConfigType implements the ConfigData interface
func (fr *FirewallRule) ConfigType() ConfigType { return ConfigFirewall }

// ConfigSubType implements the ConfigData interface
func (fr *FirewallRule) ConfigSubType() string { return "rule" }

// ConfigName implements the ConfigData interface
func (fr *FirewallRule) ConfigName() string { return "" }
