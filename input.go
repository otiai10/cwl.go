package cwl

// Input represents an element of "inputs" in "CWL".
type Input struct {
	ID      string
	Types   []Type
	Doc     string
	Label   string
	Binding *Binding
	Default *InputDefault
	Format  string
}

// New constructs "Input" struct from interface{}.
func (_ Input) New(i interface{}) Input {
	dest := Input{}
	switch x := i.(type) {
	case map[string]interface{}:
		for key, v := range x {
			switch key {
			case "id":
				dest.ID = v.(string)
			case "type":
				dest.Types = Type{}.NewList(v)
			case "label":
				dest.Label = v.(string)
			case "doc":
				dest.Doc = v.(string)
			case "inputBinding":
				dest.Binding = Binding{}.New(v)
			case "default":
				dest.Default = InputDefault{}.New(v)
			case "format":
				dest.Format = v.(string)
			}
		}
	case string:
		dest.Types = Type{}.NewList(x)
	case []interface{}:
		for _, v := range x {
			dest.Types = append(dest.Types, Type{}.New(v))
		}
	}
	return dest
}

// Inputs represents "inputs" field in CWL.
type Inputs []Input

// New constructs new "Inputs" struct.
func (_ Inputs) New(i interface{}) Inputs {
	dest := Inputs{}
	switch x := i.(type) {
	case []interface{}:
		for _, v := range x {
			dest = append(dest, Input{}.New(v))
		}
	case map[string]interface{}:
		for key, v := range x {
			input := Input{}.New(v)
			input.ID = key
			dest = append(dest, input)
		}
	}
	return dest
}
