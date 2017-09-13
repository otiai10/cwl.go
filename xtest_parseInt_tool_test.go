package cwl

import (
	"testing"

	. "github.com/otiai10/mint"
)

func TestDecode_parseInt_tool(t *testing.T) {
	f := cwl("parseInt-tool.cwl")
	root := NewCWL()
	Expect(t, root).TypeOf("*cwl.Root")
	err := root.Decode(f)
	Expect(t, err).ToBe(nil)
	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Class).ToBe("ExpressionTool")
	Expect(t, len(root.Requirements)).ToBe(1)
	Expect(t, root.Requirements).TypeOf("cwl.Requirements")
	Expect(t, root.Requirements[0].Class).ToBe("InlineJavascriptRequirement")
	Expect(t, len(root.Inputs)).ToBe(1)
	Expect(t, root.Inputs[0].ID).ToBe("file1")
	Expect(t, root.Inputs[0].Types[0].Type).ToBe("File")
	Expect(t, root.Inputs[0].Binding.LoadContents).ToBe(true)
	Expect(t, len(root.Outputs)).ToBe(1)
	Expect(t, root.Outputs[0].ID).ToBe("output")
	Expect(t, root.Outputs[0].Types[0].Type).ToBe("int")
	Expect(t, root.Expression).ToBe("$({'output': parseInt(inputs.file1.contents)})")
}
