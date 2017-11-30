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
			case "outputBinding":
				dest.Binding = Binding{}.New(v)
			}
		}
	}
	return dest
}

// Inputs represents "inputs" field in CWL.
type Fields []Field

// New constructs new "Inputs" struct.
func (_ Fields) New(i interface{}) Fields {
	dest := Fields{}
	switch x := i.(type) {
	case []interface{}:
		for _, v := range x {
			dest = append(dest, Field{}.New(v))
		}
	case map[string]interface{}:
		//for key, v := range x {
		for _, v := range x {
			field := Field{}.New(v)
			//field.Name = key
			dest = append(dest, field)
		}
	}
	return dest
}

// Len for sorting.
func (ins Fields) Len() int {
	return len(ins)
}

// Less for sorting.
func (ins Fields) Less(i, j int) bool {
	prev, next := ins[i].Binding, ins[j].Binding
	switch [2]bool{prev == nil, next == nil} {
	case [2]bool{true, true}:
		return true
	case [2]bool{false, true}:
		return prev.Position < 0
	case [2]bool{true, false}:
		return next.Position > 0
	default:
		if prev.Position != next.Position {
			return prev.Position <= next.Position
		}
		// sort by parameter name
		return ins[i].Name <= ins[j].Name
	}
}

// Swap for sorting.
func (ins Fields) Swap(i, j int) {
	ins[i], ins[j] = ins[j], ins[i]
}
