package cwl

import (
	"testing"

	. "github.com/otiai10/mint"
)

func TestDecode_count_lines8_wf(t *testing.T) {
	f := cwl("count-lines8-wf.cwl")
	root := NewCWL()
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
	Expect(t, root.Steps[0].Run.ID).ToBe("count-lines1-wf.cwl")
	Expect(t, root.Steps[0].In[0].ID).ToBe("file1")
	Expect(t, root.Steps[0].In[0].Source[0]).ToBe("file1")
	Expect(t, root.Steps[0].Out[0].ID).ToBe("count_output")
}
