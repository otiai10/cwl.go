package cwlgotest

import (
	"sort"
	"testing"

	cwl "github.com/otiai10/cwl.go"
	. "github.com/otiai10/mint"
)

func TestDecode_basename_fields_test(t *testing.T) {
	f := load("basename-fields-test.cwl")
	root := cwl.NewCWL()
	err := root.Decode(f)
	Expect(t, err).ToBe(nil)
	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Class).ToBe("Workflow")
	Expect(t, root.Requirements[0].Class).ToBe("StepInputExpressionRequirement")

	sort.Sort(root.Inputs)
	Expect(t, root.Inputs[0].ID).ToBe("tool")
	Expect(t, root.Inputs[0].Types[0].Type).ToBe("File")
	// TODO check specification for this test ID and Type
	Expect(t, root.Outputs[0].ID).ToBe("extFile")
	Expect(t, root.Outputs[0].Types[0].Type).ToBe("File")
	Expect(t, root.Outputs[0].Source[0]).ToBe("ext/out")
	Expect(t, root.Outputs[1].ID).ToBe("rootFile")
	Expect(t, root.Outputs[1].Types[0].Type).ToBe("File")
	Expect(t, root.Outputs[1].Source[0]).ToBe("root/out")
	count := 0
	for _, st := range root.Steps {
		switch st.ID {
		case "root":
			Expect(t, st.Run.Value).ToBe("echo-file-tool.cwl")
			Expect(t, st.In[0].ValueFrom).ToBe("$(inputs.tool.nameroot)")
			Expect(t, st.Out[0].ID).ToBe("out")
			// TODO tool: tool
			count = count + 1
		case "ext":
			Expect(t, st.Run.Value).ToBe("echo-file-tool.cwl")
			Expect(t, st.In[0].ValueFrom).ToBe("$(inputs.tool.nameext)")
			Expect(t, st.Out[0].ID).ToBe("out")
			// TODO tool: tool
			count = count + 1
		}
	}
	Expect(t, count).ToBe(2)
}
