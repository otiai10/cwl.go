package cwlgotest

import (
	"sort"
	"testing"

	cwl "github.com/otiai10/cwl.go"
	. "github.com/otiai10/mint"
)

func TestDecode_step_valuefrom2_wf(t *testing.T) {
	f := load("step-valuefrom2-wf.cwl")
	root := cwl.NewCWL()
	err := root.Decode(f)
	Expect(t, err).ToBe(nil)
	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Class).ToBe("Workflow")
	Expect(t, root.Requirements[0].Class).ToBe("StepInputExpressionRequirement")
	Expect(t, root.Requirements[1].Class).ToBe("InlineJavascriptRequirement")
	Expect(t, root.Requirements[2].Class).ToBe("MultipleInputFeatureRequirement")
	sort.Sort(root.Inputs)
	Expect(t, root.Inputs[1].ID).ToBe("a")
	Expect(t, root.Inputs[1].Types[0].Type).ToBe("int")
	Expect(t, root.Inputs[0].ID).ToBe("b")
	Expect(t, root.Inputs[0].Types[0].Type).ToBe("int")
	Expect(t, root.Outputs[0].ID).ToBe("val")
	Expect(t, root.Outputs[0].Types[0].Type).ToBe("string")
	Expect(t, root.Outputs[0].Source[0]).ToBe("step1/echo_out")
	Expect(t, root.Steps[0].ID).ToBe("step1")
	// TODO test run: id: echo
	Expect(t, root.Steps[0].Run.Workflow.Class).ToBe("CommandLineTool")
	Expect(t, root.Steps[0].Run.Workflow.Inputs[0].ID).ToBe("c")
	Expect(t, root.Steps[0].Run.Workflow.Inputs[0].Types[0].Type).ToBe("int")
	Expect(t, root.Steps[0].Run.Workflow.Outputs[0].ID).ToBe("echo_out")
	Expect(t, root.Steps[0].Run.Workflow.Outputs[0].Types[0].Type).ToBe("string")
	Expect(t, root.Steps[0].Run.Workflow.Outputs[0].Binding.Glob[0]).ToBe("step1_out")
	Expect(t, root.Steps[0].Run.Workflow.Outputs[0].Binding.LoadContents).ToBe(true)
	Expect(t, root.Steps[0].Run.Workflow.Outputs[0].Binding.Eval.Raw).ToBe("$(self[0].contents)")
	Expect(t, root.Steps[0].Run.Workflow.BaseCommands[0]).ToBe("echo")
	Expect(t, root.Steps[0].Run.Workflow.Stdout).ToBe("step1_out")
	Expect(t, root.Steps[0].In[0].ID).ToBe("c")
	Expect(t, root.Steps[0].In[0].Source[0]).ToBe("a")
	Expect(t, root.Steps[0].In[0].Source[1]).ToBe("b")
	Expect(t, root.Steps[0].In[0].ValueFrom).ToBe("$(self[0] + self[1])")
	Expect(t, root.Steps[0].Out[0].ID).ToBe("echo_out")
}
