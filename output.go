package cwl

// Output represents and conbines "CommandOutputParameter" and "WorkflowOutputParameter"
// @see
// - http://www.commonwl.org/v1.0/CommandLineTool.html#CommandOutputParameter
// - http://www.commonwl.org/v1.0/Workflow.html#WorkflowOutputParameter
type Output struct {
	ID             string   `json:"id"`
	Label          string   `json:"label"`
	Doc            []string `json:"doc"`
	Format         string   `json:"format"`
	Binding        *Binding `json:"outputBinding"`
	Source         []string `json:"outputSource"`
	Types          []Type   `json:"type"`
	SecondaryFiles []SecondaryFile
}

// New constructs "Output" struct from interface.
func (_ Output) New(i interface{}) Output {
	dest := Output{}
	switch x := i.(type) {
	case map[string]interface{}:
		for key, v := range x {
			switch key {
			case "id":
				dest.ID = v.(string)
			case "type":
				dest.Types = Type{}.NewList(v)
			case "outputBinding":
				dest.Binding = Binding{}.New(v)
			case "outputSource":
				dest.Source = StringArrayable(v)
			case "format":
				dest.Format = v.(string)
			case "secondaryFiles":
				dest.SecondaryFiles = SecondaryFile{}.NewList(v)
			}
		}
	case string:
		dest.Types = Type{}.NewList(x)
	}
	return dest
}

// Outputs represents "outputs" field in "CWL".
type Outputs []Output

// New constructs "Outputs" struct.
func (_ Outputs) New(i interface{}) Outputs {
	dest := Outputs{}
	switch x := i.(type) {
	case []interface{}:
		for _, v := range x {
			dest = append(dest, Output{}.New(v))
		}
	case map[string]interface{}:
		for key, v := range x {
			output := Output{}.New(v)
			output.ID = key
			dest = append(dest, output)
		}
	}
	return dest
}

// SecondaryFile represents an element of "secondaryFiles".
type SecondaryFile struct {
	Entry string
}

// NewList constructs list of "SecondaryFile".
func (typ SecondaryFile) NewList(i interface{}) []SecondaryFile {
	dest := []SecondaryFile{}
	switch x := i.(type) {
	case []interface{}:
		for _, v := range x {
			dest = append(dest, SecondaryFile{Entry: v.(string)})
		}
	}
	return dest
}
