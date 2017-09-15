package cwlgotest

import (
	"testing"

	cwl "github.com/otiai10/cwl.go"
	. "github.com/otiai10/mint"
)

func TestDecode_glob_expr_list(t *testing.T) {
	f := load("glob-expr-list.cwl")
	root := cwl.NewCWL()
	Expect(t, root).TypeOf("*cwl.Root")
	err := root.Decode(f)
	Expect(t, err).ToBe(nil)
	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Class).ToBe("CommandLineTool")
	Expect(t, len(root.Inputs)).ToBe(1)
	Expect(t, root.Inputs[0].ID).ToBe("ids")
	Expect(t, root.Inputs[0].Types[0].Type).ToBe("string[]")
	Expect(t, root.Inputs[0].Binding.Position).ToBe(1)
	Expect(t, len(root.Outputs)).ToBe(1)
	Expect(t, root.Outputs[0].ID).ToBe("files")
	Expect(t, root.Outputs[0].Types[0].Type).ToBe("File[]")
	Expect(t, root.Outputs[0].Binding.Glob).ToBe([]string{"$(inputs.ids)"})
	Expect(t, len(root.BaseCommands)).ToBe(1)
	Expect(t, root.BaseCommands[0]).ToBe("touch")
}
