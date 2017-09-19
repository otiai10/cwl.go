package cwlgotest

import (
	"testing"

	cwl "github.com/otiai10/cwl.go"
	. "github.com/otiai10/mint"
)

func TestDecode_dir3(t *testing.T) {
	f := load("dir3.cwl")
	root := cwl.NewCWL()
	err := root.Decode(f)
	Expect(t, err).ToBe(nil)

	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Class).ToBe("CommandLineTool")
	Expect(t, root.BaseCommands[0]).ToBe("tar")
	Expect(t, root.BaseCommands[1]).ToBe("xvf")
	Expect(t, root.Inputs[0].ID).ToBe("inf")
	Expect(t, root.Inputs[0].Types[0].Type).ToBe("File")
	Expect(t, root.Inputs[0].Binding.Position).ToBe(1)
	Expect(t, root.Outputs[0].ID).ToBe("outdir")
	Expect(t, root.Outputs[0].Types[0].Type).ToBe("Directory")
	Expect(t, root.Outputs[0].Binding.Glob).ToBe([]string{"."})
}
