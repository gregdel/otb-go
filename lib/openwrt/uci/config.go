package uci

import (
	"encoding/json"
	"fmt"
	"sort"
)

// ConfigType represents a configuration type
type ConfigType string

// Possible configuration types
const (
	ConfigNetwork  ConfigType = "network"
	ConfigFirewall            = "firewall"
	ConfigDHCP                = "dhcp"
	ConfigDSCP                = "dscp"
)

// ConfigData is in interface to implement to be a configuration block
type ConfigData interface {
	ConfigType() ConfigType
	ConfigSubType() string
}

// ConfigBlock represents a configuration block
type ConfigBlock struct {
	Type ConfigType
	ConfigBlockHeader
	ConfigData
}

// UnmarshalJSON implements the marshaler interface
func (cb *ConfigBlock) UnmarshalJSON(in []byte) error {
	// Unmarshal the headers
	cb.ConfigBlockHeader = ConfigBlockHeader{}
	if err := json.Unmarshal(in, &cb.ConfigBlockHeader); err != nil {
		return err
	}

	// Check the block config
	if _, ok := unmarshalFuncs[cb.Type]; !ok {
		fmt.Printf("invalid config section: %q\n", cb.Type)
		return nil
	}

	// Get the proper function to call
	f, ok := unmarshalFuncs[cb.Type][cb.ConfigBlockHeader.Type]
	if !ok {
		// Fail softly for now
		fmt.Printf("invalid block type: %q\n", cb.ConfigBlockHeader.Type)
		return nil
	}

	// Get the data from the registered function
	data, err := f(in)
	if err != nil {
		return err
	}
	cb.ConfigData = data

	return nil
}

// ConfigBlockHeader represents the generic informations present in each uci
// section
type ConfigBlockHeader struct {
	Anonymous bool   `json:".anonymous"`
	Index     int    `json:".index"`
	Type      string `json:".type"`
	Name      string `json:".name"`
}

// Configuration represents a configuration
type Configuration struct {
	Type         ConfigType
	ConfigBlocks []*ConfigBlock
}

// UnmarshalJSON implements the unmarshaler interface
func (c *Configuration) UnmarshalJSON(in []byte) error {
	var rawData = struct {
		Values map[string]json.RawMessage `json:"values"`
	}{Values: map[string]json.RawMessage{}}

	if err := json.Unmarshal(in, &rawData); err != nil {
		return err
	}

	c.ConfigBlocks = make([]*ConfigBlock, len(rawData.Values))

	i := 0
	for _, rj := range rawData.Values {
		cb := &ConfigBlock{Type: c.Type}
		if err := cb.UnmarshalJSON(rj); err != nil {
			return nil
		}

		c.ConfigBlocks[i] = cb
		i++
	}

	// Sort the result by order
	sort.Slice(c.ConfigBlocks, func(i, j int) bool {
		return c.ConfigBlocks[j].Index > c.ConfigBlocks[i].Index
	})

	return nil
}

// NewConfiguration returns an empty configuration
func NewConfiguration(t ConfigType) *Configuration {
	return &Configuration{
		Type: t,
	}
}

var unmarshalFuncs = map[ConfigType]map[string]func([]byte) (ConfigData, error){
	ConfigNetwork: {
		"interface": unmarshalNetworkInterface,
	},
}

func unmarshalNetworkInterface(input []byte) (ConfigData, error) {
	ni := &NetworkInterface{}
	return ni, json.Unmarshal(input, ni)
}
