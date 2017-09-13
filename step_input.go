package cwl

// StepInput represents WorkflowStepInput.
// @see http://www.commonwl.org/v1.0/Workflow.html#WorkflowStepInput
type StepInput struct {
	ID        string
	Source    []string
	LinkMerge string
	Default   *InputDefault
}

// New constructs a StepInput struct from any interface.
func (_ StepInput) New(i interface{}) StepInput {
	dest := StepInput{}
	switch x := i.(type) {
	case map[string]interface{}:
		for key, v := range x {
			if dest.ID == "" {
				dest.ID = key
			}

			if key == "id" {
				dest.ID = v.(string)
			} else {
				switch e := v.(type) {
				case string:
					dest.Source = []string{e}
				case []interface{}:
					for _, s := range e {
						dest.Source = append(dest.Source, s.(string))
					}
				case map[string]interface{}:
					for key, v := range e {
						switch key {
						case "id":
							dest.ID = v.(string)
						case "source":
							if list, ok := v.([]interface{}); ok {
								for _, s := range list {
									dest.Source = append(dest.Source, s.(string))
								}
							}
						case "linkMerge":
							dest.LinkMerge = v.(string)
						case "default":
							dest.Default = InputDefault{}.New(v)
						}
					}
				}
			}
		}
	}
	return dest
}

// NewList constructs a list of StepInput from interface.
func (_ StepInput) NewList(i interface{}) []StepInput {
	dest := []StepInput{}
	switch x := i.(type) {
	case []interface{}:
		for _, v := range x {
			dest = append(dest, StepInput{}.New(v))
		}
	case map[string]interface{}:
		for key, v := range x {
			item := make(map[string]interface{})
			item[key] = v
			dest = append(dest, StepInput{}.New(item))
		}
	default:
		dest = append(dest, StepInput{}.New(x))
	}
	return dest
}
