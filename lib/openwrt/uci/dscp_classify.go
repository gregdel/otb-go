package uci

func init() {
	registerConfigType(ConfigDSCP, "classify", func() ConfigData { return &DSCPClassifyConfig{} })
}

// DSCPClassifyConfig represents a DSCP rule
type DSCPClassifyConfig struct {
	Direction string `json:"direction,omitempty"`
	Proto     string `json:"proto,omitempty"`
	DestIP    string `json:"dest_ip,omitempty"`
	DestPort  string `json:"dest_port,omitempty"`
	SrcIP     string `json:"src_ip,omitempty"`
	SrcPort   string `json:"src_port,omitempty"`
	Class     string `json:"class,omitempty"`
	Comment   string `json:"comment,omitempty"`
}

// ConfigType implements the ConfigData interface
func (dc *DSCPClassifyConfig) ConfigType() ConfigType { return ConfigDSCP }

// ConfigSubType implements the ConfigData interface
func (dc *DSCPClassifyConfig) ConfigSubType() string { return "classify" }

// ConfigName implements the ConfigData interface
func (dc *DSCPClassifyConfig) ConfigName() string { return "" }
