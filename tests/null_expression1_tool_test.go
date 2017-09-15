package cwlgotest

import (
	"testing"

	cwl "github.com/otiai10/cwl.go"
	. "github.com/otiai10/mint"
)

func TestDecode_null_expression1_tool(t *testing.T) {
	f := load("null-expression1-tool.cwl")
	root := cwl.NewCWL()
	err := root.Decode(f)
	Expect(t, err).ToBe(nil)
	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Class).ToBe("ExpressionTool")
	Expect(t, len(root.Requirements)).ToBe(1)
	Expect(t, root.Requirements[0].Class).ToBe("InlineJavascriptRequirement")
	Expect(t, len(root.Inputs)).ToBe(1)
	Expect(t, root.Inputs[0].ID).ToBe("i1")
	Expect(t, root.Inputs[0].Types[0].Type).ToBe("Any")
	// TODO input default
	//Expect(t, root.Inputs[0].Default.Class).ToBe("File")
	//fmt.Println(t, root.Inputs[0].Default)
	Expect(t, root.Outputs[0].ID).ToBe("output")
	Expect(t, root.Outputs[0].Types[0].Type).ToBe("int")
	Expect(t, root.Expression).ToBe(`$({'output': (inputs.i1 == 'the-default' ? 1 : 2)})`)
}
