package uci

func init() {
	registerConfigType(ConfigDHCP, "dnsmasq", func() ConfigData { return &DHCPDnsmasq{} })
}

// DHCPDnsmasq represents a dnsmasq configuration
type DHCPDnsmasq struct {
	DomainNeeded     string `json:"domainneeded,omitempty"`
	BogusPriv        string `json:"boguspriv,omitempty"`
	FilterWin2k      string `json:"filterwin2k,omitempty"`
	LocaliseQueries  string `json:"localise_queries,omitempty"`
	RebindProtection string `json:"rebind_protection,omitempty"`
	RebindLocalhost  string `json:"rebind_localhost,omitempty"`
	Local            string `json:"local,omitempty"`
	Domain           string `json:"domain,omitempty"`
	ExpandHosts      string `json:"expandhosts,omitempty"`
	Authoritative    string `json:"authoritative,omitempty"`
	ReadEthers       string `json:"readethers,omitempty"`
	LeaseFile        string `json:"leasefile,omitempty"`
	ResolvFile       string `json:"resolvfile,omitempty"`
	NonWildcard      string `json:"nonwildcard,omitempty"`
	LocalService     string `json:"localservice,omitempty"`
	CacheSize        string `json:"cachesize,omitempty"`
	FilterAAAA       string `json:"filter_aaaa,omitempty"`
	AllServers       string `json:"all_servers,omitempty"`
	QuietDHCP        string `json:"quietdhcp,omitempty"`
	Interfaces       List   `json:"interface,omitempty"`
}

// ConfigType implements the ConfigData interface
func (dd *DHCPDnsmasq) ConfigType() ConfigType { return ConfigDHCP }

// ConfigSubType implements the ConfigData interface
func (dd *DHCPDnsmasq) ConfigSubType() string { return "dnsmasq" }

// ConfigName implements the ConfigData interface
func (dd *DHCPDnsmasq) ConfigName() string { return "" }
