package cwlgotest

import (
	"testing"

	cwl "github.com/otiai10/cwl.go"
	. "github.com/otiai10/mint"
)

func TestDecode_revsort(t *testing.T) {
	f := load("revsort.cwl")
	root := cwl.NewCWL()
	err := root.Decode(f)
	Expect(t, err).ToBe(nil)
	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Class).ToBe("Workflow")
	Expect(t, root.Doc).ToBe("Reverse the lines in a document, then sort those lines.")
	Expect(t, root.Hints[0].Class).ToBe("DockerRequirement")
	Expect(t, root.Hints[0].DockerPull).ToBe("debian:8")
	Expect(t, len(root.Inputs)).ToBe(2)
	Expect(t, root.Inputs[0].ID).ToBe("input")
	Expect(t, root.Inputs[0].Doc).ToBe("The input file to be processed.")
	Expect(t, root.Inputs[0].Types[0].Type).ToBe("File")
	Expect(t, root.Inputs[1].Types[0].Type).ToBe("boolean")
	Expect(t, root.Inputs[1].ID).ToBe("reverse_sort")
	Expect(t, root.Inputs[1].Doc).ToBe("If true, reverse (decending) sort")
	Expect(t, root.Outputs[0].ID).ToBe("output")
	Expect(t, root.Outputs[0].Types[0].Type).ToBe("File")
	Expect(t, root.Outputs[0].Source[0]).ToBe("sorted/output")
	Expect(t, root.Outputs[0].Doc[0]).ToBe("The output with the lines reversed and sorted.")
	Expect(t, len(root.Steps)).ToBe(2)
	Expect(t, root.Steps[0].ID).ToBe("rev")
	Expect(t, root.Steps[0].In[0].ID).ToBe("input")
	Expect(t, root.Steps[0].In[0].Source[0]).ToBe("input")
	Expect(t, root.Steps[0].Out[0].ID).ToBe("output")
	Expect(t, root.Steps[0].Run.Value).ToBe("revtool.cwl")
	Expect(t, root.Steps[1].ID).ToBe("sorted")
	Expect(t, root.Steps[1].In[0].ID).ToBe("input")
	Expect(t, root.Steps[1].In[0].Source[0]).ToBe("rev/output")
	Expect(t, root.Steps[1].In[1].ID).ToBe("reverse")
	Expect(t, root.Steps[1].In[1].Source[0]).ToBe("reverse_sort")
	Expect(t, root.Steps[1].Out[0].ID).ToBe("output")
	Expect(t, root.Steps[1].Run.Value).ToBe("sorttool.cwl")
}
