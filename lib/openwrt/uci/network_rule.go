package uci

func init() {
	registerConfigType(ConfigNetwork, "rule", func() ConfigData { return &NetworkRule{} })
}

// NetworkRule represents a network rule
type NetworkRule struct {
	In       string `json:"in"`
	Out      string `json:"out"`
	Src      string `json:"src"`
	Dest     string `json:"dest"`
	ToS      string `json:"tos"`
	Mark     string `json:"mark"`
	Invert   string `json:"invert"`
	Priority string `json:"priority"`
	Lookup   string `json:"lookup"`
	Goto     string `json:"goto"`
	Action   string `json:"action"`
}

// ConfigType implements the ConfigData interface
func (nr *NetworkRule) ConfigType() ConfigType { return ConfigNetwork }

// ConfigSubType implements the ConfigData interface
func (nr *NetworkRule) ConfigSubType() string { return "rule" }
