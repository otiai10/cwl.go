package cwl

// RequirementEnvDef only appears if requirement class is "InitialWorkDirRequirement"
type EnvDef struct {
        envName  string
        envValue string
}

// NewList constructs a list of RequirementEntry from interface
func (list EnvDef) NewList(i interface{}) []EnvDef {
        dest := []EnvDef{}
        switch x := i.(type) {
        case map[string]interface{}:
                for key, v := range x {
                        r := EnvDef{}
                        r.envName = key
                        r.envValue = v.(string)
                        dest = append(dest, r)
                }
        }
        return dest
}
