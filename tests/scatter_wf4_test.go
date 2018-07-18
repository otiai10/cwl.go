package cwlgotest

import (
	"testing"

	cwl "github.com/otiai10/cwl.go"
	. "github.com/otiai10/mint"
)

func TestDecode_scatter_wf4(t *testing.T) {
	f := load("scatter-wf4.cwl")
	root := cwl.NewCWL()
	err := root.Decode(f)
	Expect(t, err).ToBe(err)

	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Graphs[0].ID).ToBe("echo")
	Expect(t, root.Graphs[0].Class).ToBe("CommandLineTool")
	Expect(t, root.Graphs[0].Inputs[0].ID).ToBe("echo_in1")
	Expect(t, root.Graphs[0].Inputs[0].Types[0].Type).ToBe("string")
	Expect(t, root.Graphs[0].Inputs[1].ID).ToBe("echo_in2")
	Expect(t, root.Graphs[0].Inputs[1].Types[0].Type).ToBe("string")
	Expect(t, root.Graphs[0].Outputs[0].ID).ToBe("echo_out")
	Expect(t, root.Graphs[0].Outputs[0].Types[0].Type).ToBe("string")
	Expect(t, root.Graphs[0].Outputs[0].Binding.Glob[0]).ToBe("step1_out")
	Expect(t, root.Graphs[0].Outputs[0].Binding.LoadContents).ToBe(true)
	Expect(t, root.Graphs[0].Outputs[0].Binding.Eval.Raw).ToBe("$(self[0].contents)")
	Expect(t, root.Graphs[0].BaseCommands[0]).ToBe("echo")
	Expect(t, root.Graphs[0].Arguments[0].Value).ToBe("-n")
	Expect(t, root.Graphs[0].Arguments[1].Value).ToBe("foo")
	Expect(t, root.Graphs[0].Stdout).ToBe("step1_out")

	Expect(t, root.Graphs[1].ID).ToBe("main")
	Expect(t, root.Graphs[1].Class).ToBe("Workflow")
	Expect(t, root.Graphs[1].Inputs[0].ID).ToBe("inp1")
	Expect(t, root.Graphs[1].Inputs[0].Types[0].Type).ToBe("string[]")
	Expect(t, root.Graphs[1].Inputs[1].ID).ToBe("inp2")
	Expect(t, root.Graphs[1].Inputs[1].Types[0].Type).ToBe("string[]")
	Expect(t, root.Graphs[1].Requirements[0].Class).ToBe("ScatterFeatureRequirement")
	Expect(t, root.Graphs[1].Steps[0].Scatter[0]).ToBe("echo_in1")
	Expect(t, root.Graphs[1].Steps[0].Scatter[1]).ToBe("echo_in2")
	Expect(t, root.Graphs[1].Steps[0].ScatterMethod).ToBe("dotproduct")
	Expect(t, root.Graphs[1].Steps[0].ID).ToBe("step1")
	Expect(t, root.Graphs[1].Steps[0].In[0].ID).ToBe("echo_in1")
	Expect(t, root.Graphs[1].Steps[0].In[0].Source[0]).ToBe("inp1")
	Expect(t, root.Graphs[1].Steps[0].In[1].ID).ToBe("echo_in2")
	Expect(t, root.Graphs[1].Steps[0].In[1].Source[0]).ToBe("inp2")
	Expect(t, root.Graphs[1].Steps[0].Out[0].ID).ToBe("echo_out")
	Expect(t, root.Graphs[1].Steps[0].Run.Value).ToBe("#echo")

	Expect(t, root.Graphs[1].Outputs[0].ID).ToBe("out")
	Expect(t, root.Graphs[1].Outputs[0].Source[0]).ToBe("step1/echo_out")
	Expect(t, root.Graphs[1].Outputs[0].Types[0].Type).ToBe("array")
	Expect(t, root.Graphs[1].Outputs[0].Types[0].Items[0].Type).ToBe("string")

}
