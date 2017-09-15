package cwlgotest

import (
	"testing"

	cwl "github.com/otiai10/cwl.go"
	. "github.com/otiai10/mint"
)

func TestDecode_count_lines12_wf(t *testing.T) {
	f := load("count-lines12-wf.cwl")
	root := cwl.NewCWL()
	err := root.Decode(f)
	Expect(t, err).ToBe(nil)

	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Class).ToBe("Workflow")

	Expect(t, root.Inputs[0].ID).ToBe("file1")
	Expect(t, root.Inputs[0].Types[0].Type).ToBe("array")
	Expect(t, root.Inputs[0].Types[0].Items[0].Type).ToBe("File")
	Expect(t, root.Inputs[1].ID).ToBe("file2")
	Expect(t, root.Inputs[1].Types[0].Type).ToBe("array")
	Expect(t, root.Inputs[1].Types[0].Items[0].Type).ToBe("File")
	Expect(t, root.Outputs[0].ID).ToBe("count_output")
	Expect(t, root.Outputs[0].Types[0].Type).ToBe("int")
	Expect(t, root.Outputs[0].Source).ToBe([]string{"step1/output"})

	Expect(t, root.Requirements[0].Class).ToBe("MultipleInputFeatureRequirement")

	Expect(t, root.Steps[0].ID).ToBe("step1")
	Expect(t, root.Steps[0].Run.Value).ToBe("wc3-tool.cwl")
	Expect(t, root.Steps[0].In[0].ID).ToBe("file1")
	Expect(t, root.Steps[0].In[0].Source[0]).ToBe("file1")
	Expect(t, root.Steps[0].In[0].Source[1]).ToBe("file2")
	Expect(t, root.Steps[0].In[0].LinkMerge).ToBe("merge_flattened")
	Expect(t, root.Steps[0].Out[0].ID).ToBe("output")
}
