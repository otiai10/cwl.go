package cwl

// Hints ...
type Hints []Hint

// New constructs "Hints" struct.
func (hints Hints) New(i interface{}) Hints {
	switch x := i.(type) {
	case []interface{}:
		dest := make([]Hint, len(x))
		for i, elm := range x {
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
	case map[string]interface{}:
		dest := []Hint{}
		for key, val := range x {
			hint := Hint{}
			hint["class"] = key
			dict, ok := val.(map[string]interface{})
			if !ok {
				return dest
			}
			for name, v := range dict {
				hint[name] = v
			}
			dest = append(dest, hint)
		}
		return dest
	}
	return Hints{}
}

// Hint ...
type Hint map[string]interface{}
