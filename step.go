package cwl

// Steps represents "steps" field in CWL.
type Steps []Step

// New constructs "Steps" from interface.
func (_ Steps) New(i interface{}) Steps {
	dest := Steps{}
	switch x := i.(type) {
	case []interface{}:
		// TODO;
	case map[string]interface{}:
		for key, v := range x {
			s := Step{}.New(v)
			s.ID = key
			dest = append(dest, s)
		}
	}
	return dest
}

// Step represents WorkflowStep.
// @see http://www.commonwl.org/v1.0/Workflow.html#WorkflowStep
type Step struct {
	ID           string
	In           []StepInput
	Out          []StepOutput
	Run          *Root
	Requirements []Requirement
	Scatter      string
}

// New constructs "Step" from interface.
func (_ Step) New(i interface{}) Step {
	dest := Step{}
	switch x := i.(type) {
	case map[string]interface{}:
		for key, v := range x {
			switch key {
			case "run":
				dest.Run = dest.Run.AsStep(v)
			case "in":
				dest.In = StepInput{}.NewList(v)
			case "out":
				dest.Out = StepOutput{}.NewList(v)
			case "requirements":
				dest.Requirements = Requirements{}.New(v)
			case "scatter":
				dest.Scatter = v.(string)
			}
		}
	}
	return dest
}
