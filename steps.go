package cwl

// Steps represents "steps" field in CWL.
type Steps []Step

// New constructs "Steps" from interface.
func (steps Steps) New(i interface{}) Steps {
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

// Step represents an element of "steps"
type Step struct {
	ID      string
	Run     *Root
	In      []StepInput
	Out     []StepOutput
	Scatter string
}

// New constructs "Step" from interface.
func (step Step) New(i interface{}) Step {
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
			case "scatter":
				dest.Scatter = v.(string)
			}
		}
	}
	return dest
}

// StepInput ...
type StepInput struct {
	Name     string
	Location string
}

// NewList constructs a list of StepInput from interface.
func (si StepInput) NewList(i interface{}) []StepInput {
	dest := []StepInput{}
	switch x := i.(type) {
	case []interface{}:
		// TODO:
	case map[string]interface{}:
		input := StepInput{}
		for key, v := range x {
			input.Name = key
			input.Location = v.(string)
		}
		dest = append(dest, input)
	}
	return dest
}

// StepOutput ...
type StepOutput struct {
	Name     string
	Location string
}

// NewList constructs a list of StepOutput from interface.
func (so StepOutput) NewList(i interface{}) []StepOutput {
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
func (so StepOutput) New(i interface{}) StepOutput {
	dest := StepOutput{}
	switch x := i.(type) {
	case string:
		dest.Name = x
		dest.Location = x
	}
	return dest
}
