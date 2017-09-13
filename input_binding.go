package cwl

// InputBinding represents "inputBinding" field in an element of "inputs".
type InputBinding struct {
	Position     int
	Prefix       string
	Separator    string
	LoadContents bool
}

// New constructs new "InputBinding".
func (binding InputBinding) New(i interface{}) *InputBinding {
	dest := new(InputBinding)
	switch x := i.(type) {
	case map[string]interface{}:
		for key, val := range x {
			switch key {
			case "position":
				dest.Position = int(val.(float64))
			case "prefix":
				dest.Prefix = val.(string)
			case "itemSeparator":
				dest.Separator = val.(string)
			case "loadContents":
				dest.LoadContents = val.(bool)
			}
		}
	}
	return dest
}
