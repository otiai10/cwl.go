package cwlgotest

import (
	"sort"
	"testing"

	cwl "github.com/otiai10/cwl.go"
	. "github.com/otiai10/mint"
)

func TestDecode_scatter_valuefrom_wf6(t *testing.T) {
	f := load("scatter-valuefrom-wf6.cwl")
	root := cwl.NewCWL()
	err := root.Decode(f)
	Expect(t, err).ToBe(err)

	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Class).ToBe("Workflow")
	Expect(t, root.Requirements[0].Class).ToBe("ScatterFeatureRequirement")
	Expect(t, root.Requirements[1].Class).ToBe("StepInputExpressionRequirement")
	Expect(t, root.Inputs[0].ID).ToBe("scattered_messages")
	Expect(t, root.Inputs[0].Types[0].Type).ToBe("string[]")
	Expect(t, root.Outputs[0].ID).ToBe("out_message")
	Expect(t, root.Outputs[0].Types[0].Type).ToBe("File[]")
	Expect(t, root.Outputs[0].Source[0]).ToBe("step1/out_message")
	Expect(t, root.Steps[0].ID).ToBe("step1")
	Expect(t, root.Steps[0].Run.Value).ToBe("scatter-valueFrom-tool.cwl")
	Expect(t, root.Steps[0].Scatter[0]).ToBe("scattered_message")
	Expect(t, root.Steps[0].ScatterMethod).ToBe("dotproduct")
	sort.Sort(root.Steps[0].In)
	Expect(t, root.Steps[0].In[1].ID).ToBe("scattered_message")
	Expect(t, root.Steps[0].In[1].Source[0]).ToBe("scattered_messages")
	Expect(t, root.Steps[0].In[0].ID).ToBe("message")
	Expect(t, root.Steps[0].In[0].ValueFrom).ToBe("Hello")
	Expect(t, root.Steps[0].Out[0].ID).ToBe("out_message")
}
