package cwl

// Schemas ...
type Schemas []string

// New constructs "Schemas" struct.
func (schemas Schemas) New(i interface{}) Schemas {
	dest := Schemas{}
	switch x := i.(type) {
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
