package cwl

// Inputs represents "inputs" field in CWL.
type Inputs []Input

// New constructs new "Inputs" struct.
func (inputs Inputs) New(i interface{}) Inputs {
	dest := Inputs{}
	switch x := i.(type) {
	case []interface{}:
		for _, elm := range x {
			input := Input{}.New(elm)
			dest = append(dest, input)
		}
	case map[string]interface{}:
		for key, val := range x {
			input := Input{}.New(val)
			input.ID = key
			dest = append(dest, input)
		}
	}
	return dest
}

// Input represents an element of "inputs" in "CWL".
type Input struct {
	ID      string
	Types   []InputType
	Doc     string
	Label   string
	Binding *InputBinding
	Default *InputDefault
	Format  string
}

// New constructs "Input" struct from interface{}.
func (input Input) New(i interface{}) Input {
	dest := Input{}
	switch x := i.(type) {
	case map[string]interface{}:
		dest = input.NewFromDict(x)
	case string:
		dest.Types = []InputType{{Type: x}}
	case []interface{}: // count-lines12-wf.cwl suggests it can be array with length 1.
		if len(x) == 0 {
			return dest
		}
		if dict, ok := x[0].(map[string]interface{}); ok {
			dest.Types = InputType{}.NewList(dict)
		}
	}
	return dest
}

// NewFromDict constructs "Input" from dictionary formed map.
func (input Input) NewFromDict(dict map[string]interface{}) Input {
	dest := Input{}
	for key, val := range dict {
		switch key {
		case "id":
			dest.ID = val.(string)
		case "type":
			dest.Types = InputType{}.NewList(val)
		case "label":
			dest.Label = val.(string)
		case "doc":
			dest.Doc = val.(string)
		case "inputBinding":
			dest.Binding = InputBinding{}.New(val)
		case "default":
			dest.Default = InputDefault{}.New(val)
		case "format":
			dest.Format = val.(string)
		}
	}
	return dest
}
