package uci

func init() {
	registerConfigType(ConfigNetwork, "globals", func() ConfigData { return &NetworkGlobals{} })
}

// NetworkGlobals represents a network device
type NetworkGlobals struct {
	ULAPrefix string `json:"ula_prefix,omitempty"`
	Multipath string `json:"multipath,omitempty"`
}

// ConfigType implements the ConfigData interface
func (nd *NetworkGlobals) ConfigType() ConfigType { return ConfigNetwork }

// ConfigSubType implements the ConfigData interface
func (nd *NetworkGlobals) ConfigSubType() string { return "globals" }

// ConfigName implements the ConfigData interface
func (nd *NetworkGlobals) ConfigName() string { return "globals" }
