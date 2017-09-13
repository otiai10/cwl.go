package cwl

// StepOutput represents StepWorkflowOutput.
// @see http://www.commonwl.org/v1.0/Workflow.html#WorkflowStepOutput
type StepOutput struct {
	ID string
}

// NewList constructs a list of StepOutput from interface.
func (_ StepOutput) NewList(i interface{}) []StepOutput {
	dest := []StepOutput{}
	switch x := i.(type) {
	case []interface{}:
		for _, v := range x {
			dest = append(dest, StepOutput{}.New(v))
		}
	}
	return dest
}

// New constructs a StepOutput from interface.
func (_ StepOutput) New(i interface{}) StepOutput {
	dest := StepOutput{}
	switch x := i.(type) {
	case string:
		dest.ID = x
	}
	return dest
}
