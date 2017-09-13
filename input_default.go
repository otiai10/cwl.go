package cwl

// InputDefault represents "default" field in an element of "inputs".
type InputDefault struct {
	Class    string
	Location string
}

// New constructs new "InputDefault".
func (def InputDefault) New(i interface{}) *InputDefault {
	dest := new(InputDefault)
	switch x := i.(type) {
	case map[string]interface{}:
		for key, val := range x {
			switch key {
			case "class":
				dest.Class = val.(string)
			case "location":
				dest.Location = val.(string)
			}
		}
	}
	return dest
}
