package cwlgotest

import (
	"testing"

	cwl "github.com/otiai10/cwl.go"
	. "github.com/otiai10/mint"
)

func TestDecode_sorttool(t *testing.T) {
	f := load("sorttool.cwl")
	root := cwl.NewCWL()
	err := root.Decode(f)
	Expect(t, err).ToBe(nil)
	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Class).ToBe("CommandLineTool")
	Expect(t, root.Doc).ToBe("Sort lines using the `sort` command")
	Expect(t, len(root.Inputs)).ToBe(2)
	Expect(t, root.Inputs[0].ID).ToBe("reverse")
	Expect(t, root.Inputs[0].Types[0].Type).ToBe("boolean")
	Expect(t, root.Inputs[0].Binding.Position).ToBe(1)
	Expect(t, root.Inputs[0].Binding.Prefix).ToBe("--reverse")
	Expect(t, root.Inputs[1].ID).ToBe("input")
	Expect(t, root.Inputs[1].Types[0].Type).ToBe("File")
	Expect(t, root.Inputs[1].Binding.Position).ToBe(2)
	Expect(t, len(root.Outputs)).ToBe(1)
	Expect(t, root.Outputs[0].ID).ToBe("output")
	Expect(t, root.Outputs[0].Types[0].Type).ToBe("File")
	Expect(t, root.Outputs[0].Binding.Glob[0]).ToBe("output.txt")
	Expect(t, root.BaseCommands[0]).ToBe("sort")
	Expect(t, root.Stdout).ToBe("output.txt")
}
