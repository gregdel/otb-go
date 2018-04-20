package ubus

import (
	"github.com/gregdel/otb-go/lib/openwrt/uci"
)

// UciParams reprensents the ubus parameters required by uci
type UciParams struct {
	Config  string      `json:"config,omitempty"`
	Section string      `json:"section,omitempty"`
	Name    string      `json:"name,omitempty"`
	Type    string      `json:"type,omitempty"`
	Option  string      `json:"option,omitempty"`
	Options []string    `json:"options,omitempty"`
	Match   string      `json:"match,omitempty"`
	Values  interface{} `json:"values,omitempty"`
}

// NewUciParams returns a new UciParams with a config
func NewUciParams(config string) *UciParams {
	return &UciParams{Config: config}
}

// UciGet gets the configuration in uci
func UciGet(result interface{}, config string) error {
	return callWithParams(result, NewUciParams(config), "uci", "get")
}

// UciGetConfig returns a raw configuration
func UciGetConfig(t uci.ConfigType) (*uci.Configuration, error) {
	c := uci.NewConfiguration(t)
	return c, UciGet(c, string(t))
}

// UciCommit commits the changes
func UciCommit(config string) error {
	return callWithParams(nil, NewUciParams(config), "uci", "commit")
}

// UciAdd adds a new configuration
func UciAdd(cb *uci.ConfigBlock) error {
	p := &UciParams{
		Config: string(cb.ConfigType()),
		Type:   cb.ConfigSubType(),
		Name:   cb.ConfigName(),
		Values: cb.ConfigData,
	}

	return callWithParams(nil, p, "uci", "add")
}
