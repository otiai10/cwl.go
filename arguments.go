package cwl

// Arguments ...
type Arguments []Argument

// New constructs "Arguments" struct.
func (baseCommands Arguments) New(i interface{}) Arguments {
	dest := Arguments{}
	switch x := i.(type) {
	case string:
		argument := Argument{}
		argument.String = x
		dest = append(dest, argument)
	case []interface{}:
		for _, elm := range x {
			argument := Argument{}
			switch val := elm.(type) {
			case string:
				argument.String = val
			case map[string]interface{}:
				argument.CommandLineBinding = val
			}
			dest = append(dest, argument)
		}
	}
	return dest
}

// Argument
type Argument struct {
        String         string
        CommandLineBinding         map[string]interface{}
        // TODO support Expression
}
