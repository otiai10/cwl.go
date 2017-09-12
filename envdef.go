package cwl

// EnvDef only appears if requirement class is "EnvVarRequirement"
type EnvDef struct {
	Name  string
	Value string
}

// NewList constructs a list of EnvDef from interface
func (list EnvDef) NewList(i interface{}) []EnvDef {
	dest := []EnvDef{}
	switch x := i.(type) {
	case map[string]interface{}:
		for key, v := range x {
			r := EnvDef{}
			r.Name = key
			r.Value = v.(string)
			dest = append(dest, r)
		}
	}
	return dest
}
