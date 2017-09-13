package cwl

// Steps represents "steps" field in CWL.
type Steps []Step

// New constructs "Steps" from interface.
func (_ Steps) New(i interface{}) Steps {
	dest := Steps{}
	switch x := i.(type) {
	case []interface{}:
		for _, v := range x {
			s := Step{}.New(v)
			dest = append(dest, s)
		}
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
	Run          Run
	Requirements []Requirement
	Scatter      string
}

// Run `run` accept string | CommandLineTool | ExpressionTool | Workflow
type Run struct {
	Value    string
	Workflow *Root
}

// New constructs "Step" from interface.
func (_ Step) New(i interface{}) Step {
	dest := Step{}
	switch x := i.(type) {
	case map[string]interface{}:
		for key, v := range x {
			switch key {
			case "id":
				dest.ID = v.(string)
			case "run":
				switch x2 := v.(type) {
				case string:
					dest.Run.Value = x2
				case map[string]interface{}:
					dest.Run.Workflow = dest.Run.Workflow.AsStep(v)
				}
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
