package cwlgotest

import (
	"sort"
	"testing"

	cwl "github.com/otiai10/cwl.go"
	. "github.com/otiai10/mint"
)

func TestDecode_count_lines10_wf(t *testing.T) {
	f := load("count-lines10-wf.cwl")
	root := cwl.NewCWL()
	err := root.Decode(f)
	Expect(t, err).ToBe(nil)

	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Class).ToBe("Workflow")

	Expect(t, root.Inputs[0].ID).ToBe("file1")
	Expect(t, root.Inputs[0].Types[0].Type).ToBe("File")
	Expect(t, root.Outputs[0].ID).ToBe("count_output")
	Expect(t, root.Outputs[0].Types[0].Type).ToBe("int")
	Expect(t, root.Outputs[0].Source).ToBe([]string{"step1/count_output"})

	Expect(t, root.Requirements[0].Class).ToBe("SubworkflowFeatureRequirement")

	Expect(t, root.Steps[0].ID).ToBe("step1")
	Expect(t, root.Steps[0].In[0].ID).ToBe("file1")
	Expect(t, root.Steps[0].In[0].Source[0]).ToBe("file1")
	Expect(t, root.Steps[0].Out[0].ID).ToBe("count_output")

	Expect(t, root.Steps[0].Run.Workflow.Class).ToBe("Workflow")
	Expect(t, root.Steps[0].Run.Workflow.Inputs[0].ID).ToBe("file1")
	Expect(t, root.Steps[0].Run.Workflow.Inputs[0].Types[0].Type).ToBe("File")
	Expect(t, root.Steps[0].Run.Workflow.Outputs[0].ID).ToBe("count_output")
	Expect(t, root.Steps[0].Run.Workflow.Outputs[0].Types[0].Type).ToBe("int")
	Expect(t, root.Steps[0].Run.Workflow.Outputs[0].Source).ToBe([]string{"step2/output"})
	// Recursive steps
	sort.Sort(root.Steps[0].Run.Workflow.Steps)
	Expect(t, root.Steps[0].Run.Workflow.Steps[0].ID).ToBe("step1")
	Expect(t, root.Steps[0].Run.Workflow.Steps[0].Run.Value).ToBe("wc-tool.cwl")
	Expect(t, root.Steps[0].Run.Workflow.Steps[0].In[0].ID).ToBe("file1")
	Expect(t, root.Steps[0].Run.Workflow.Steps[0].In[0].Source[0]).ToBe("file1")
	Expect(t, root.Steps[0].Run.Workflow.Steps[0].Out[0].ID).ToBe("output")
	Expect(t, root.Steps[0].Run.Workflow.Steps[1].ID).ToBe("step2")
	Expect(t, root.Steps[0].Run.Workflow.Steps[1].Run.Value).ToBe("parseInt-tool.cwl")
	Expect(t, root.Steps[0].Run.Workflow.Steps[1].In[0].ID).ToBe("file1")
	Expect(t, root.Steps[0].Run.Workflow.Steps[1].In[0].Source[0]).ToBe("step1/output")
	Expect(t, root.Steps[0].Run.Workflow.Steps[1].Out[0].ID).ToBe("output")
}
