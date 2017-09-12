package cwl

// Inputs represents "inputs" field in CWL.
type Inputs []RequiredInput

// New constructs new "Inputs" struct.
func (inputs Inputs) New(i interface{}) Inputs {
	dest := Inputs{}
	switch x := i.(type) {
	case []interface{}:
		for _, elm := range x {
			input := RequiredInput{}.New(elm)
			dest = append(dest, input)
		}
	case map[string]interface{}:
		for key, val := range x {
			input := RequiredInput{}.New(val)
			input.ID = key
			dest = append(dest, input)
		}
	}
	return dest
}

// RequiredInput represents an element of "inputs" in "CWL".
type RequiredInput struct {
	ID      string
	Types   []InputType
	Doc     string
	Label   string
	Binding *InputBinding
	Default *InputDefault
	Format  string
}

// New constructs "RequiredInput" struct from interface{}.
func (input RequiredInput) New(i interface{}) RequiredInput {
	dest := RequiredInput{}
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

// NewFromDict constructs "RequiredInput" from dictionary formed map.
func (input RequiredInput) NewFromDict(dict map[string]interface{}) RequiredInput {
	dest := RequiredInput{}
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

// InputBinding represents "inputBinding" field in an element of "inputs".
type InputBinding struct {
	Position     int
	Prefix       string
	Separator    string
	LoadContents bool
}

// New constructs new "InputBinding".
func (binding InputBinding) New(i interface{}) *InputBinding {
	dest := new(InputBinding)
	switch x := i.(type) {
	case map[string]interface{}:
		for key, val := range x {
			switch key {
			case "position":
				dest.Position = int(val.(float64))
			case "prefix":
				dest.Prefix = val.(string)
			case "itemSeparator":
				dest.Separator = val.(string)
			case "loadContents":
				dest.LoadContents = val.(bool)
			}
		}
	}
	return dest
}

// InputDefault represents "default" field in an element of "inputs".
type InputDefault struct {
	Class    string
	Location string
}

// New constructs new "InputDefault".
func (def InputDefault) New(i interface{}) *InputDefault {
	dest := new(InputDefault)
	switch x := i.(type) {
	case map[string]interface{}:
		for key, val := range x {
			switch key {
			case "class":
				dest.Class = val.(string)
			case "location":
				dest.Location = val.(string)
			}
		}
	}
	return dest
}
