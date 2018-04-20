package uci

import (
	"encoding/json"
	"fmt"
)

// List represents a UCI list, it can be an array of string or a single string
type List []string

// UnmarshalJSON implements the Unmarshaler interface
func (l *List) UnmarshalJSON(in []byte) error {
	// Try to unmarshal the result as an array
	r := []string{}
	if err := json.Unmarshal(in, &r); err == nil {
		*l = r
		return nil
	}

	s := ""
	if err := json.Unmarshal(in, &s); err == nil {
		*l = []string{s}
		return nil
	}

	return fmt.Errorf("failed to unmarshal input: %q", string(in))
}
