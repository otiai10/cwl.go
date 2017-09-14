package cwl

// Input represents "CommandInputParameter".
// @see http://www.commonwl.org/v1.0/CommandLineTool.html#CommandInputParameter
type Input struct {
	ID             string        `json:"id"`
	Label          string        `json:"label"`
	Doc            string        `json:"doc"`
	Format         string        `json:"format"`
	Binding        *Binding      `json:"inputBinding"`
	Default        *InputDefault `json:"default"`
	Types          []Type        `json:"type"`
	SecondaryFiles []SecondaryFile
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
			case "secondaryFiles":
				dest.SecondaryFiles = SecondaryFile{}.NewList(v)
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
