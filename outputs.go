package cwl

// Outputs represents "outputs" field in "CWL".
type Outputs []Output

// New constructs "Outputs" struct.
func (outputs Outputs) New(i interface{}) Outputs {
	dest := Outputs{}
	switch x := i.(type) {
	case []interface{}:
		for _, elm := range x {
			dest = append(dest, Output{}.New(elm))
		}
	case map[string]interface{}:
		for key, val := range x {
			output := Output{}.New(val)
			output.ID = key
			dest = append(dest, output)
		}
	}
	return dest
}

// Output represents an element of "outputs".
type Output struct {
	ID    string
	Types []OutputType // Possible types
}

// New constructs "Output" struct from interface.
func (output Output) New(i interface{}) Output {
	dest := Output{}
	switch x := i.(type) {
	case map[string]interface{}:
		for key, val := range x {
			switch key {
			case "id":
				dest.ID = val.(string)
			case "type":
				dest.Types = OutputType{}.NewList(val)
			}
		}
	case string: // If it's simple dictionary, value represents Type.
		dest.Types = []OutputType{{Type: x}}
	}
	return dest
}

// OutputType represents "type" of an element of "outputs".
type OutputType struct {
	Type  string
	Items string // type of array element if "Type" is "array"
}

// NewList constructs list of "OutputType".
func (typ OutputType) NewList(i interface{}) []OutputType {
	dest := []OutputType{}
	switch x := i.(type) {
	case []interface{}:
		for _, elm := range x {
			dest = append(dest, OutputType{Type: elm.(string)})
		}
	case map[string]interface{}:
		t := OutputType{}
		for key, val := range x {
			switch key {
			case "type":
				t.Type = val.(string)
			case "items":
				t.Items = val.(string)
			}
		}
		dest = append(dest, t)
	}
	return dest
}
