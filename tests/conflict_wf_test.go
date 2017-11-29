package cwlgotest

import (
	"sort"
	"testing"

	cwl "github.com/otiai10/cwl.go"
	. "github.com/otiai10/mint"
)

func TestDecode_conflict_wf(t *testing.T) {
	f := load("conflict-wf.cwl")
	root := cwl.NewCWL()
	err := root.Decode(f)
	Expect(t, err).ToBe(err)

	Expect(t, root.Version).ToBe("v1.0")
	sort.Sort(root.Graphs)

	Expect(t, root.Graphs[0].ID).ToBe("cat")
	Expect(t, root.Graphs[0].Class).ToBe("CommandLineTool")
	sort.Sort(root.Graphs[0].Inputs)
	Expect(t, root.Graphs[0].Inputs[0].ID).ToBe("file1")
	Expect(t, root.Graphs[0].Inputs[0].Types[0].Type).ToBe("File")
	Expect(t, root.Graphs[0].Inputs[0].Binding.Position).ToBe(1)
	Expect(t, root.Graphs[0].Inputs[1].ID).ToBe("file2")
	Expect(t, root.Graphs[0].Inputs[1].Types[0].Type).ToBe("File")
	Expect(t, root.Graphs[0].Inputs[1].Binding.Position).ToBe(2)
	Expect(t, root.Graphs[0].Outputs[0].ID).ToBe("fileout")
	Expect(t, root.Graphs[0].Outputs[0].Types[0].Type).ToBe("File")
	Expect(t, root.Graphs[0].Outputs[0].Binding.Glob[0]).ToBe("out.txt")
	Expect(t, root.Graphs[0].BaseCommands[0]).ToBe("cat")
	Expect(t, root.Graphs[0].Stdout).ToBe("out.txt")

	Expect(t, root.Graphs[1].ID).ToBe("collision")
	Expect(t, root.Graphs[1].Class).ToBe("Workflow")
	Expect(t, root.Graphs[1].Inputs[0].ID).ToBe("input_1")
	Expect(t, root.Graphs[1].Inputs[0].Types[0].Type).ToBe("string")
	Expect(t, root.Graphs[1].Inputs[1].ID).ToBe("input_2")
	Expect(t, root.Graphs[1].Inputs[1].Types[0].Type).ToBe("string")
	Expect(t, root.Graphs[1].Outputs[0].ID).ToBe("fileout")
	Expect(t, root.Graphs[1].Outputs[0].Types[0].Type).ToBe("File")
	Expect(t, root.Graphs[1].Outputs[0].Source[0]).ToBe("cat_step/fileout")
	sort.Sort(root.Graphs[1].Steps)
	Expect(t, root.Graphs[1].Steps[0].ID).ToBe("cat_step")
	Expect(t, root.Graphs[1].Steps[0].Run.Value).ToBe("#cat")
	Expect(t, root.Graphs[1].Steps[0].In[0].ID).ToBe("file1")
	Expect(t, root.Graphs[1].Steps[0].In[0].Source[0]).ToBe("echo_1/fileout")
	Expect(t, root.Graphs[1].Steps[0].In[1].ID).ToBe("file2")
	Expect(t, root.Graphs[1].Steps[0].In[1].Source[0]).ToBe("echo_2/fileout")
	Expect(t, root.Graphs[1].Steps[0].Out[0].ID).ToBe("fileout")
	Expect(t, root.Graphs[1].Steps[1].ID).ToBe("echo_1")
	Expect(t, root.Graphs[1].Steps[1].Run.Value).ToBe("#echo")
	Expect(t, root.Graphs[1].Steps[1].In[0].ID).ToBe("text")
	Expect(t, root.Graphs[1].Steps[1].In[0].Source[0]).ToBe("input_1")
	Expect(t, root.Graphs[1].Steps[1].Out[0].ID).ToBe("fileout")
	Expect(t, root.Graphs[1].Steps[2].ID).ToBe("echo_2")
	Expect(t, root.Graphs[1].Steps[2].Run.Value).ToBe("#echo")
	Expect(t, root.Graphs[1].Steps[2].In[0].ID).ToBe("text")
	Expect(t, root.Graphs[1].Steps[2].In[0].Source[0]).ToBe("input_2")
	Expect(t, root.Graphs[1].Steps[2].Out[0].ID).ToBe("fileout")

	Expect(t, root.Graphs[2].ID).ToBe("echo")
	Expect(t, root.Graphs[2].Class).ToBe("CommandLineTool")
	Expect(t, root.Graphs[2].Inputs[0].ID).ToBe("text")
	Expect(t, root.Graphs[2].Inputs[0].Types[0].Type).ToBe("string")
	Expect(t, root.Graphs[2].Outputs[0].ID).ToBe("fileout")
	Expect(t, root.Graphs[2].Outputs[0].Types[0].Type).ToBe("File")
	Expect(t, root.Graphs[2].Outputs[0].Binding.Glob[0]).ToBe("out.txt")
	Expect(t, root.Graphs[2].BaseCommands[0]).ToBe("echo")
	Expect(t, root.Graphs[2].Stdout).ToBe("out.txt")
}
