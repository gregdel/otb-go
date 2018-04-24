package uci

func init() {
	registerConfigType(ConfigFirewall, "redirect", func() ConfigData { return &FirewallRedirect{} })
}

// FirewallRedirect represents a firewall redirection
type FirewallRedirect struct {
	Name     string `json:"name,omitempty"`
	Src      string `json:"src,omitempty"`
	SrcIP    string `json:"src_ip,omitempty"`
	SrcPort  string `json:"src_port,omitempty"`
	Dest     string `json:"dest,omitempty"`
	DestIP   string `json:"dest_ip,omitempty"`
	DestPort string `json:"dest_port,omitempty"`
	Proto    string `json:"proto,omitempty"`
	Limit    string `json:"limit,omitempty"`
	Family   string `json:"family,omitempty"`
	Target   string `json:"target,omitempty"`
}

// ConfigType implements the ConfigData interface
func (fr *FirewallRedirect) ConfigType() ConfigType { return ConfigFirewall }

// ConfigSubType implements the ConfigData interface
func (fr *FirewallRedirect) ConfigSubType() string { return "redirect" }

// ConfigName implements the ConfigData interface
func (fr *FirewallRedirect) ConfigName() string { return "" }
