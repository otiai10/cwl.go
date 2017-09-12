package cwl

// Arguments ...
type Arguments []string

// New constructs "Arguments" struct.
func (baseCommands Arguments) New(i interface{}) Arguments {
	dest := Arguments{}
	switch x := i.(type) {
	case string:
		dest = append(dest, x)
	case []interface{}:
		for _, elm := range x {
			str, ok := elm.(string)
			if !ok {
                                return dest
                        }
			dest = append(dest, str)
		}
	}
	return dest
}

