package cwl

// EnvDef represents "EnvironmentDef",
// only appears if requirement class is "EnvVarRequirement"
// @see http://www.commonwl.org/v1.0/CommandLineTool.html#EnvironmentDef
type EnvDef struct {
	Name  string
	Value string
}

// NewList constructs a list of EnvDef from interface
func (_ EnvDef) NewList(i interface{}) []EnvDef {
	dest := []EnvDef{}
	switch x := i.(type) {
	case map[string]interface{}:
		for key, v := range x {
			dest = append(dest, EnvDef{Name: key, Value: v.(string)})
		}
	}
	return dest
}
