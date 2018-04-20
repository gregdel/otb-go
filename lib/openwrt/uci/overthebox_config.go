package uci

func init() {
	registerConfigType(ConfigOverthebox, "config", func() ConfigData { return &OvertheboxConfig{} })
}

// OvertheboxConfig represents the overthebox config
type OvertheboxConfig struct {
	DeviceID        string `json:"device_id,omitempty"`
	Token           string `json:"token,omitempty"`
	Service         string `json:"service,omitempty"`
	NeedsActivation string `json:"needs_activation,omitempty"`
}

// ConfigType implements the ConfigData interface
func (oc *OvertheboxConfig) ConfigType() ConfigType { return ConfigOverthebox }

// ConfigSubType implements the ConfigData interface
func (oc *OvertheboxConfig) ConfigSubType() string { return "config" }

// ConfigName implements the ConfigData interface
func (oc *OvertheboxConfig) ConfigName() string { return "me" }
