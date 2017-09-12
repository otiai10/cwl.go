package cwl

// BaseCommands ...
type BaseCommands []string

// New constructs "BaseCommands" struct.
func (baseCommands BaseCommands) New(i interface{}) BaseCommands {
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
	return BaseCommands{}
}
