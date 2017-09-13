package cwl

// Type represents CWL Typeable objects.
// - http://www.commonwl.org/v1.0/CommandLineTool.html#CWLType
// - http://www.commonwl.org/v1.0/CommandLineTool.html#CommandInputRecordSchema
// - http://www.commonwl.org/v1.0/CommandLineTool.html#CommandInputEnumSchema
// - http://www.commonwl.org/v1.0/CommandLineTool.html#CommandInputArraySchema
type Type struct {
	Type    string
	Label   string
	Binding *Binding
	Fields  []Field  // from CommandInputRecordSchema
	Symbols []string // from CommandInputEnumSchema
	Items   []Type   // from CommandInputArraySchema
}

// NewList constructs a list of Type from any interface.
// It only handles []interface{}
func (_ Type) NewList(i interface{}) []Type {
	dest := []Type{}
	switch x := i.(type) {
	case []interface{}:
		for _, s := range x {
			dest = append(dest, Type{}.New(s))
		}
	default:
		dest = append(dest, Type{}.New(x))
	}
	return dest
}

// New constructs single Type struct from any interface.
func (_ Type) New(i interface{}) Type {
	dest := Type{}
	switch x := i.(type) {
	case string:
		dest.Type = x
	case map[string]interface{}:
		if val, ok := x["type"]; ok {
			dest.Type = val.(string)
		}
		if val, ok := x["items"]; ok {
			dest.Items = Type{}.NewList(val)
		}
		if val, ok := x["inputBinding"]; ok {
			dest.Binding = Binding{}.New(val)
		}
	}
	return dest
}
