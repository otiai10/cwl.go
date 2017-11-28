package cwl

// Argument represents an element of "arguments" of CWL
// @see http://www.commonwl.org/v1.0/CommandLineTool.html#CommandLineTool
type Argument struct {
	Value   string
	Binding *Binding
}

// New constructs an "Argument" struct from any interface.
func (_ Argument) New(i interface{}) Argument {
	dest := Argument{}
	switch x := i.(type) {
	case string:
		dest.Value = x
	case map[string]interface{}:
		dest.Binding = Binding{}.New(x)
	}
	return dest
}

// Flatten ...
func (arg Argument) Flatten() []string {
	flattened := []string{}
	if arg.Value != "" {
		flattened = append(flattened, arg.Value)
	}
	if arg.Binding != nil {
		if arg.Binding.Prefix != "" {
			flattened = append([]string{arg.Binding.Prefix}, flattened...)
		}
	}
	return flattened
}

// Arguments represents a list of "Argument"
type Arguments []Argument

// New constructs "Arguments" struct.
func (_ Arguments) New(i interface{}) Arguments {
	dest := Arguments{}
	switch x := i.(type) {
	case []interface{}:
		for _, v := range x {
			dest = append(dest, Argument{}.New(v))
		}
	default:
		dest = append(dest, Argument{}.New(x))
	}
	return dest
}

// Len for sorting.
func (args Arguments) Len() int {
	return len(args)
}

// Less for sorting.
func (args Arguments) Less(i, j int) bool {
	prev, next := args[i].Binding, args[j].Binding
	switch [2]bool{prev == nil, next == nil} {
	case [2]bool{true, true}:
		return false
	case [2]bool{false, true}:
		return prev.Position < 0
	case [2]bool{true, false}:
		return next.Position > 0
	default:
		return prev.Position <= next.Position
	}
}

// Swap for sorting.
func (args Arguments) Swap(i, j int) {
	args[i], args[j] = args[j], args[i]
}
