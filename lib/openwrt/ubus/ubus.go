package ubus

import (
	"encoding/json"
	"os/exec"
)

// call performs a ubus call to the system and unmarshals the results into
// result
func call(result interface{}, args ...string) error {
	out, err := exec.Command("ubus", append([]string{"call"}, args...)...).Output()
	if err != nil {
		return err
	}

	return json.Unmarshal(out, result)
}

// callWithParams marshals the parameters
func callWithParams(result interface{}, params map[string]string, args ...string) error {
	out, err := json.Marshal(params)
	if err != nil {
		return err
	}

	return call(result, append(args, string(out))...)
}
