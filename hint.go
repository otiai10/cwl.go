package cwl

// Hints ...
type Hints []Hint

// New constructs "Hints" struct.
func (hints Hints) New(i interface{}) Hints {
	list, ok := i.([]interface{})
	if !ok {
		return Hints{}
	}
	dest := make([]Hint, len(list))
	for i, elm := range list {
		m, ok := elm.(map[string]interface{})
		if !ok {
			return dest
		}
		hint := Hint{}
		for key, val := range m {
			hint[key] = val
		}
		dest[i] = hint
	}
	return dest
}

// Hint ...
type Hint map[string]interface{}
