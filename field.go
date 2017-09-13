package cwl

// Field represents CommandInputRecordField
// @see http://www.commonwl.org/v1.0/CommandLineTool.html#CommandInputRecordField
type Field struct {
	Name    string
	Doc     string
	Types   []Type
	Binding *Binding
	Label   string
}

// NewList constructs a list of Field from any interface.
func (_ Field) NewList(i interface{}) []Field {
	dest := []Field{}
	switch x := i.(type) {
	case []interface{}:
		for _, v := range x {
			dest = append(dest, Field{}.New(v))
		}
	}
	return dest
}

// New constructs a Field struct from any interface.
func (_ Field) New(i interface{}) Field {
	dest := Field{}
	switch x := i.(type) {
	case map[string]interface{}:
		for key, v := range x {
			switch key {
			case "name":
				dest.Name = v.(string)
			case "type":
				dest.Types = Type{}.NewList(v)
			case "inputBinding":
				dest.Binding = Binding{}.New(v)
			}
		}
	}
	return dest
}
