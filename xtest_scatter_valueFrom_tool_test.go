package cwl

import (
	"testing"

	. "github.com/otiai10/mint"
)

func TestDecode_scatter_valueFrom_tool(t *testing.T) {
	f := cwl("scatter-valueFrom-tool.cwl")
	root := NewCWL()
	err := root.Decode(f)
	Expect(t, err).ToBe(err)

	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Class).ToBe("CommandLineTool")

	Expect(t, root.Inputs[0].ID).ToBe("scattered_message")
	Expect(t, root.Inputs[0].Types[0].Type).ToBe("string")
	Expect(t, root.Inputs[0].Binding.Position).ToBe(2)
	Expect(t, root.Inputs[1].ID).ToBe("message")
	Expect(t, root.Inputs[1].Types[0].Type).ToBe("string")
	Expect(t, root.Inputs[1].Binding.Position).ToBe(1)
	Expect(t, root.Outputs[0].ID).ToBe("out_message")
	Expect(t, root.Outputs[0].Types[0].Type).ToBe("stdout")
	Expect(t, root.BaseCommands[0]).ToBe("echo")
}
