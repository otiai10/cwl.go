package cwlgotest

import (
	"testing"

	cwl "github.com/otiai10/cwl.go"
	. "github.com/otiai10/mint"
)

func TestDecode_dir5(t *testing.T) {
	f := load("dir5.cwl")
	root := cwl.NewCWL()
	err := root.Decode(f)
	Expect(t, err).ToBe(nil)

	Expect(t, root.Requirements[0].Class).ToBe("ShellCommandRequirement")
	Expect(t, root.Requirements[1].Class).ToBe("InitialWorkDirRequirement")
	Expect(t, root.Requirements[1].Listing[0].Location).ToBe("$(inputs.indir.listing)")
	Expect(t, root.Inputs[0].ID).ToBe("indir")
	Expect(t, root.Inputs[0].Types[0].Type).ToBe("Directory")
	Expect(t, root.Outputs[0].ID).ToBe("outlist")
	Expect(t, root.Outputs[0].Types[0].Type).ToBe("File")
	Expect(t, root.Outputs[0].Binding.Glob).ToBe([]string{"output.txt"})
	Expect(t, root.Arguments[0].Value).ToBe("find")
	Expect(t, root.Arguments[1].Value).ToBe("-L")
	Expect(t, root.Arguments[2].Value).ToBe(".")
	Expect(t, root.Arguments[3].Binding.ShellQuote).ToBe(false)
	Expect(t, root.Arguments[3].Binding.ValueFrom.Key()).ToBe("|")
	Expect(t, root.Arguments[4].Value).ToBe("sort")
	Expect(t, root.Stdout).ToBe("output.txt")
}
