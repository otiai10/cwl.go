package cwl

// Arguments ...
type Arguments []string

// New constructs "Arguments" struct.
func (baseCommands Arguments) New(i interface{}) Arguments {
	switch x := i.(type) {

	case string:
		dest := []string{}
		dest = append(dest, x)
		return dest
	case []interface{}:
		dest := make([]string, len(x))
		for i, elm := range x {
			str, ok := elm.(string)
			if !ok {
                                return dest
                        }
			dest[i] = str
		}
		return dest
	}
	return Arguments{}
}

