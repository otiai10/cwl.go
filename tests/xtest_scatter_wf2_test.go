package cwlgotest

import (
	"testing"

	cwl "github.com/otiai10/cwl.go"
	. "github.com/otiai10/mint"
)

func TestDecode_scatter_wf2(t *testing.T) {
	f := load("scatter-wf2.cwl")
	root := cwl.NewCWL()
	err := root.Decode(f)
	Expect(t, err).ToBe(err)

	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Class).ToBe("Workflow")
	Expect(t, root.Inputs[0].ID).ToBe("inp1")
	Expect(t, root.Inputs[0].Types[0].Type).ToBe("string[]")
	Expect(t, root.Inputs[1].ID).ToBe("inp2")
	Expect(t, root.Inputs[1].Types[0].Type).ToBe("string[]")
	Expect(t, root.Outputs[0].ID).ToBe("out")
	Expect(t, root.Outputs[0].Source[0]).ToBe("step1/echo_out")
	Expect(t, root.Outputs[0].Types[0].Type).ToBe("array")
	Expect(t, root.Outputs[0].Types[0].Items[0].Type).ToBe("array")
	Expect(t, root.Outputs[0].Types[0].Items[0].Items[0].Type).ToBe("string")
	Expect(t, root.Requirements[0].Class).ToBe("ScatterFeatureRequirement")
	Expect(t, root.Steps[0].ID).ToBe("step1")
	Expect(t, root.Steps[0].In[0].ID).ToBe("echo_in1")
	Expect(t, root.Steps[0].In[0].Source[0]).ToBe("inp1")
	Expect(t, root.Steps[0].In[1].ID).ToBe("echo_in2")
	Expect(t, root.Steps[0].In[1].Source[0]).ToBe("inp2")
	Expect(t, root.Steps[0].Out[0].ID).ToBe("echo_out")
	Expect(t, root.Steps[0].Scatter[0]).ToBe("echo_in1")
	Expect(t, root.Steps[0].Scatter[1]).ToBe("echo_in2")
	Expect(t, root.Steps[0].ScatterMethod).ToBe("nested_crossproduct")
	Expect(t, root.Steps[0].Run.Workflow.Class).ToBe("CommandLineTool")
	Expect(t, root.Steps[0].Run.Workflow.ID).ToBe("step1command")
	Expect(t, root.Steps[0].Run.Workflow.Inputs[0].ID).ToBe("echo_in1")
	Expect(t, root.Steps[0].Run.Workflow.Inputs[0].Types[0].Type).ToBe("string")
	Expect(t, root.Steps[0].Run.Workflow.Inputs[1].ID).ToBe("echo_in2")
	Expect(t, root.Steps[0].Run.Workflow.Inputs[1].Types[0].Type).ToBe("string")
	Expect(t, root.Steps[0].Run.Workflow.Outputs[0].ID).ToBe("echo_out")
	Expect(t, root.Steps[0].Run.Workflow.Outputs[0].Types[0].Type).ToBe("string")
	Expect(t, root.Steps[0].Run.Workflow.Outputs[0].Binding.Glob[0]).ToBe("step1_out")
	Expect(t, root.Steps[0].Run.Workflow.Outputs[0].Binding.LoadContents).ToBe(true)
	Expect(t, root.Steps[0].Run.Workflow.Outputs[0].Binding.Eval).ToBe("$(self[0].contents)")
	Expect(t, root.Steps[0].Run.Workflow.BaseCommands[0]).ToBe("echo")
	Expect(t, root.Steps[0].Run.Workflow.Arguments[0].Value).ToBe("-n")
	Expect(t, root.Steps[0].Run.Workflow.Arguments[1].Value).ToBe("foo")
	Expect(t, root.Steps[0].Run.Workflow.Stdout).ToBe("step1_out")
}
