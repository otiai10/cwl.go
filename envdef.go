package cwl

// RequirementEnvDef only appears if requirement class is "InitialWorkDirRequirement"
type RequirementEnvDef struct {
        envName  string
        envValue string
}

// NewList constructs a list of RequirementEntry from interface
func (list RequirementEnvDef) NewList(i interface{}) []RequirementEnvDef {
        dest := []RequirementEnvDef{}
        switch x := i.(type) {
        case map[string]interface{}:
                for key, v := range x {
                        r := RequirementEnvDef{}
                        r.envName = key
                        r.envValue = v.(string)
                        dest = append(dest, r)
                }
        }
        return dest
}
