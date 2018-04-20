package uci

func init() {
	registerConfigType(ConfigNetwork, "rule", func() ConfigData { return &NetworkRule{} })
}

// NetworkRule represents a network rule
type NetworkRule struct {
	In       string `json:"in,omitempty"`
	Out      string `json:"out,omitempty"`
	Src      string `json:"src,omitempty"`
	Dest     string `json:"dest,omitempty"`
	ToS      string `json:"tos,omitempty"`
	Mark     string `json:"mark,omitempty"`
	Invert   string `json:"invert,omitempty"`
	Priority string `json:"priority,omitempty"`
	Lookup   string `json:"lookup,omitempty"`
	Goto     string `json:"goto,omitempty"`
	Action   string `json:"action,omitempty"`
}

// ConfigType implements the ConfigData interface
func (nr *NetworkRule) ConfigType() ConfigType { return ConfigNetwork }

// ConfigSubType implements the ConfigData interface
func (nr *NetworkRule) ConfigSubType() string { return "rule" }

// ConfigName implements the ConfigData interface
func (nr *NetworkRule) ConfigName() string { return "" }
