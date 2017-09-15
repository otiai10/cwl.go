package cwlgotest

import (
	"testing"

	cwl "github.com/otiai10/cwl.go"
	. "github.com/otiai10/mint"
)

func TestDecode_nameroot(t *testing.T) {
	f := load("nameroot.cwl")
	root := cwl.NewCWL()
	err := root.Decode(f)
	Expect(t, err).ToBe(nil)
	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Class).ToBe("CommandLineTool")
	Expect(t, len(root.Inputs)).ToBe(1)
	Expect(t, root.Inputs[0].ID).ToBe("file1")
	Expect(t, root.Inputs[0].Types[0].Type).ToBe("File")
	Expect(t, len(root.Outputs)).ToBe(1)
	Expect(t, root.Outputs[0].ID).ToBe("b")
	Expect(t, len(root.BaseCommands)).ToBe(0)
	Expect(t, len(root.Arguments)).ToBe(4)
	Expect(t, root.Arguments[0].Value).ToBe("echo")
	Expect(t, root.Arguments[1].Value).ToBe("$(inputs.file1.basename)")
	Expect(t, root.Arguments[2].Value).ToBe("$(inputs.file1.nameroot)")
	Expect(t, root.Arguments[3].Value).ToBe("$(inputs.file1.nameext)")
}
