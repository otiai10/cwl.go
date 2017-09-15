package cwlgotest

import (
	"testing"

	cwl "github.com/otiai10/cwl.go"
	. "github.com/otiai10/mint"
)

func TestDecode_echo_file_tool(t *testing.T) {
	f := load("echo-file-tool.cwl")
	root := cwl.NewCWL()
	err := root.Decode(f)
	Expect(t, err).ToBe(nil)
	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Class).ToBe("CommandLineTool")
	Expect(t, root.BaseCommands[0]).ToBe("echo")
	Expect(t, root.Inputs[0].ID).ToBe("in")
	Expect(t, root.Inputs[0].Types[0].Type).ToBe("string")
	Expect(t, root.Inputs[0].Binding.Position).ToBe(1)
	// TODO check specification for this test ID and Type
	Expect(t, root.Outputs[0].ID).ToBe("out")
	Expect(t, root.Outputs[0].Types[0].Type).ToBe("stdout")
}
