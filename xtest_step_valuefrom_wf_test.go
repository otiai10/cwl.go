package cwl

import (
	"testing"

	. "github.com/otiai10/mint"
)

func TestDecode_step_valuefrom_wf(t *testing.T) {
	f := cwl("step-valuefrom-wf.cwl")
	root := NewCWL()
	err := root.Decode(f)
	Expect(t, err).ToBe(nil)
	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Class).ToBe("Workflow")
	Expect(t, root.Inputs[0].ID).ToBe("in")
	Expect(t, root.Inputs[0].Types[0].Type).ToBe("record")
	Expect(t, root.Inputs[0].Types[0].Fields[0].Name).ToBe("file1")
	Expect(t, root.Inputs[0].Types[0].Fields[0].Types[0].Type).ToBe("File")
	// TODO check specification for this test ID and Type
	Expect(t, root.Outputs[0].ID).ToBe("count_output")
	Expect(t, root.Outputs[0].Types[0].Type).ToBe("int")
	Expect(t, root.Outputs[0].Source[0]).ToBe("step2/output")
	Expect(t, root.Steps[0].Run.Value).ToBe("wc-tool.cwl")
	Expect(t, root.Steps[0].In[0].ID).ToBe("file1")
	Expect(t, root.Steps[0].In[0].Source[0]).ToBe("in")
	Expect(t, root.Steps[0].In[0].ValueFrom).ToBe("$(self.file1)")
	Expect(t, root.Steps[0].Out[0].ID).ToBe("output")
	Expect(t, root.Steps[1].Run.Value).ToBe("parseInt-tool.cwl")
	Expect(t, root.Steps[1].In[0].ID).ToBe("file1")
	Expect(t, root.Steps[1].In[0].Source[0]).ToBe("step1/output")
	Expect(t, root.Steps[1].Out[0].ID).ToBe("output")
}
