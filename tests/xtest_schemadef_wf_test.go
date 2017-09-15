package cwlgotest

import (
	"testing"

	cwl "github.com/otiai10/cwl.go"
	. "github.com/otiai10/mint"
)

func TestDecode_schemadef_wf(t *testing.T) {
	f := load("schemadef-wf.cwl")
	root := cwl.NewCWL()
	err := root.Decode(f)
	Expect(t, err).ToBe(err)

	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Class).ToBe("Workflow")
	Expect(t, root.Requirements[0].Import).ToBe("schemadef-type.yml")
	Expect(t, root.Inputs[0].ID).ToBe("hello")
	Expect(t, root.Inputs[0].Types[0].Type).ToBe("schemadef-type.yml#HelloType")
	Expect(t, root.Outputs[0].ID).ToBe("output")
	Expect(t, root.Outputs[0].Types[0].Type).ToBe("File")
	Expect(t, root.Outputs[0].Source[0]).ToBe("step1/output")
	Expect(t, root.Steps[0].ID).ToBe("step1")
	Expect(t, root.Steps[0].In[0].ID).ToBe("hello")
	Expect(t, root.Steps[0].In[0].Source[0]).ToBe("hello")
	Expect(t, root.Steps[0].Out[0].ID).ToBe("output")
	Expect(t, root.Steps[0].Run.Value).ToBe("schemadef-tool.cwl")
}
