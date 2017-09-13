package cwl

// InputType represents "type" field in an element of "inputs".
type InputType struct {
	Type    string
	Items   string
	Binding *InputBinding
}

// NewList constructs a list of InputType from interface
func (typ InputType) NewList(i interface{}) []InputType {
	dest := []InputType{}
	switch x := i.(type) {
	case string:
		dest = append(dest, InputType{Type: x})
	case map[string]interface{}:
		t := InputType{}
		if val, ok := x["type"]; ok {
			t.Type = val.(string)
		}
		if val, ok := x["items"]; ok {
			t.Items = val.(string)
		}
		if val, ok := x["inputBinding"]; ok {
			t.Binding = InputBinding{}.New(val)
		}
		dest = append(dest, t)
	case []interface{}:
		for _, s := range x {
			dest = append(dest, InputType{Type: s.(string)})
		}
	}
	return dest
}
