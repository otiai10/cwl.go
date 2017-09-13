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
	ID             string
	Types          []OutputType // Possible types
	Binding        *OutputBinding
	SecondaryFiles []SecondaryFile
	Source         string
	Format         string
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
			case "outputBinding":
				dest.Binding = OutputBinding{}.New(val)
			case "secondaryFiles":
				dest.SecondaryFiles = SecondaryFile{}.NewList(val)
			case "outputSource":
				dest.Source = val.(string)
			case "format":
				dest.Format = val.(string)
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
	case string: // If it's simple dictionary, the value represents Type.
		t := OutputType{Type: x}
		dest = append(dest, t)
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
		for _, elm := range x {
			dest = append(dest, SecondaryFile{Entry: elm.(string)})
		}
	}
	return dest
}

// OutputBinding represents "outputBinding" of "outputs" field.
type OutputBinding struct {
	Glob string
}

// New constructs "OutputBinding" from interface{}
func (binding OutputBinding) New(i interface{}) *OutputBinding {
	dest := new(OutputBinding)
	switch x := i.(type) {
	case map[string]interface{}:
		for key, val := range x {
			switch key {
			case "glob":
				dest.Glob = val.(string)
			}
		}
	}
	return dest
}
