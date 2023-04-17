package t38c

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type cmd struct {
	Name string
	Args []string
}

func newCmd(name string, args ...string) cmd {
	return cmd{name, args}
}

func (c cmd) String() string {
	str := c.Name
	if len(c.Args) > 0 {
		str += " " + strings.Join(c.Args, " ")
	}
	return str
}

func fieldValueString(val fieldValue) string {
	var ret string
	switch val.(type) {
	case float64:
		ret = strconv.FormatFloat(val.(float64), 'f', -1, 64)
	case float32:
		ret = strconv.FormatFloat(val.(float64), 'f', -1, 32)
	default:
		bval, err := json.Marshal(val)
		if err != nil {
			return ""
		}
		ret = string(bval)
	}
	return ret
}

func floatString(val float64) string {
	return strconv.FormatFloat(val, 'f', -1, 64)
}

func rawEventHandler(handler EventHandler) func([]byte) error {
	return func(data []byte) error {
		resp := &GeofenceEvent{}
		if err := json.Unmarshal(data, resp); err != nil {
			return fmt.Errorf("json unmarshal geofence response: %v", err)
		}

		return handler.HandleEvent(resp)
	}
}
