package cwlgotest

import (
	"testing"

	cwl "github.com/otiai10/cwl.go"
	. "github.com/otiai10/mint"
)

func TestDecode_shelltest(t *testing.T) {
	f := load("shelltest.cwl")
	root := cwl.NewCWL()
	err := root.Decode(f)
	Expect(t, err).ToBe(nil)
	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Class).ToBe("CommandLineTool")
	Expect(t, root.Doc).ToBe("Reverse each line using the `rev` command then sort.")
	Expect(t, root.Requirements[0].Class).ToBe("ShellCommandRequirement")
	Expect(t, len(root.Inputs)).ToBe(1)
	Expect(t, root.Inputs[0].ID).ToBe("input")
	Expect(t, root.Inputs[0].Types[0].Type).ToBe("File")
	Expect(t, len(root.Outputs)).ToBe(1)
	Expect(t, root.Outputs[0].ID).ToBe("output")
	Expect(t, root.Outputs[0].Types[0].Type).ToBe("File")
	Expect(t, root.Outputs[0].Binding.Glob[0]).ToBe("output.txt")
	Expect(t, root.Arguments[0].Value).ToBe("rev")
	Expect(t, root.Arguments[1].Binding.ValueFrom.Key()).ToBe("inputs.input")
	Expect(t, root.Arguments[2].Binding.ValueFrom.Key()).ToBe(" | ")
	Expect(t, root.Arguments[2].Binding.ShellQuote).ToBe(false)
	Expect(t, root.Arguments[3].Value).ToBe("sort")
	Expect(t, root.Arguments[4].Binding.ValueFrom.Key()).ToBe("> output.txt")
	Expect(t, root.Arguments[4].Binding.ShellQuote).ToBe(false)
}
