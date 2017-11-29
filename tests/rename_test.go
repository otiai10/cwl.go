package cwlgotest

import (
	"sort"
	"testing"

	cwl "github.com/otiai10/cwl.go"
	. "github.com/otiai10/mint"
)

func TestDecode_rename(t *testing.T) {
	f := load("rename.cwl")
	root := cwl.NewCWL()
	err := root.Decode(f)
	Expect(t, err).ToBe(nil)
	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Class).ToBe("CommandLineTool")
	Expect(t, root.BaseCommands[0]).ToBe("true")
	Expect(t, len(root.Requirements)).ToBe(1)
	Expect(t, root.Requirements[0].Class).ToBe("InitialWorkDirRequirement")
	Expect(t, root.Requirements[0].Listing[0].EntryName).ToBe("$(inputs.newname)")
	Expect(t, root.Requirements[0].Listing[0].Entry).ToBe(`$(inputs.srcfile)`)
	Expect(t, len(root.Inputs)).ToBe(2)
	sort.Sort(root.Inputs)
	Expect(t, root.Inputs[0].ID).ToBe("srcfile")
	Expect(t, root.Inputs[0].Types[0].Type).ToBe("File")
	Expect(t, root.Inputs[1].ID).ToBe("newname")
	Expect(t, root.Inputs[1].Types[0].Type).ToBe("string")
	Expect(t, len(root.Outputs)).ToBe(1)
	Expect(t, root.Outputs[0].ID).ToBe("outfile")
	Expect(t, root.Outputs[0].Types[0].Type).ToBe("File")
	Expect(t, root.Outputs[0].Binding.Glob[0]).ToBe("$(inputs.newname)")
}
