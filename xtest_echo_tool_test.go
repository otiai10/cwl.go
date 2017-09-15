package cwl

import (
	"testing"

	. "github.com/otiai10/mint"
)

func TestDecode_echo_tool(t *testing.T) {
	f := cwl("echo-tool.cwl")
	root := NewCWL()
	err := root.Decode(f)
	Expect(t, err).ToBe(nil)
	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Class).ToBe("CommandLineTool")
	Expect(t, root.Inputs[0].ID).ToBe("in")
	Expect(t, root.Inputs[0].Types[0].Type).ToBe("Any")
	// TODO check specification for this test ID and Type
	Expect(t, root.Outputs[0].ID).ToBe("out")
	Expect(t, root.Outputs[0].Types[0].Type).ToBe("string")
	Expect(t, root.Outputs[0].Binding.Glob[0]).ToBe("out.txt")
	Expect(t, root.Outputs[0].Binding.LoadContents).ToBe(true)
	Expect(t, root.BaseCommands[0]).ToBe("echo")
	Expect(t, root.Stdout).ToBe("out.txt")
}
