package cwl

import (
	"testing"

	. "github.com/otiai10/mint"
)

func TestDecode_revtool(t *testing.T) {
	f := cwl("revtool.cwl")
	root := NewCWL()
	err := root.Decode(f)
	Expect(t, err).ToBe(nil)
	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Class).ToBe("CommandLineTool")
	Expect(t, root.Doc).ToBe("Reverse each line using the `rev` command")
	Expect(t, root.Inputs[0].ID).ToBe("input")
	Expect(t, root.Inputs[0].Types[0].Type).ToBe("File")
	// TODO check specification for this test ID and Type
	Expect(t, root.Outputs[0].ID).ToBe("output")
	Expect(t, root.Outputs[0].Types[0].Type).ToBe("File")
	Expect(t, root.Outputs[0].Binding.Glob[0]).ToBe("output.txt")
	Expect(t, root.BaseCommands[0]).ToBe("rev")
}
