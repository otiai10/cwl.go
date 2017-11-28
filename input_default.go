package cwl

import (
	"fmt"
	"reflect"
)

// InputDefault represents "default" field in an element of "inputs".
type InputDefault struct {
	Self interface{}
	Kind reflect.Kind
}

// New constructs new "InputDefault".
func (_ InputDefault) New(i interface{}) *InputDefault {
	dest := &InputDefault{Self: i, Kind: reflect.TypeOf(i).Kind()}
	return dest
}

// Flatten ...
func (d *InputDefault) Flatten(binding *Binding) []string {
	flattened := []string{}
	switch v := d.Self.(type) {
	case map[string]interface{}:
		// TODO: more strict type casting ;(
		class, ok := v["class"]
		if ok && class == "File" {
			flattened = append(flattened, fmt.Sprintf("%v", v["location"]))
		}
	}
	if binding != nil && binding.Prefix != "" {
		flattened = append([]string{binding.Prefix}, flattened...)
	}
	return flattened
}
