package cwl

// Schemas ...
type Schemas []string

// New constructs "Schemas" struct.
func (_ Schemas) New(i interface{}) Schemas {
	dest := Schemas{}
	switch x := i.(type) {
	case []interface{}:
		for _, v := range x {
			dest = append(dest, v.(string))
		}
	}
	return dest
}
