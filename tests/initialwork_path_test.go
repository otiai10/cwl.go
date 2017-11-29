package cwlgotest

import (
	"testing"

	cwl "github.com/otiai10/cwl.go"
	. "github.com/otiai10/mint"
)

func TestDecode_initialwork_path(t *testing.T) {
	f := load("initialwork-path.cwl")
	root := cwl.NewCWL()
	Expect(t, root).TypeOf("*cwl.Root")
	err := root.Decode(f)
	Expect(t, err).ToBe(nil)
	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Class).ToBe("CommandLineTool")
	Expect(t, len(root.Inputs)).ToBe(1)
	Expect(t, root.Inputs[0].ID).ToBe("file1")
	Expect(t, root.Inputs[0].Types[0].Type).ToBe("File")
	Expect(t, len(root.Outputs)).ToBe(0)
	Expect(t, root.Requirements[0].Class).ToBe("InitialWorkDirRequirement")
	Expect(t, root.Requirements[0].Listing[0].EntryName).ToBe("bob.txt")
	Expect(t, root.Requirements[0].Listing[0].Entry).ToBe(`$(inputs.file1)`)
	Expect(t, root.Requirements[1].Class).ToBe("ShellCommandRequirement")
	Expect(t, root.Arguments[0].Binding.ShellQuote).ToBe(false)
	Expect(t, root.Arguments[0].Binding.ValueFrom.Key()).ToBe(`test "$(inputs.file1.path)" = "$(runtime.outdir)/bob.txt"
`)
	// TODO write basecommand
}
