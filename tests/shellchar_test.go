package cwlgotest

import (
	"testing"

	cwl "github.com/otiai10/cwl.go"
	. "github.com/otiai10/mint"
)

func TestDecode_shellchar(t *testing.T) {
	f := load("shellchar.cwl")
	root := cwl.NewCWL()
	err := root.Decode(f)
	Expect(t, err).ToBe(nil)
	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Class).ToBe("CommandLineTool")
	Expect(t, root.Doc).ToBe("Ensure that arguments containing shell directives are not interpreted and\n" + "that `shellQuote: false` has no effect when ShellCommandRequirement is not in\n" + "effect.\n")
	Expect(t, len(root.Inputs)).ToBe(0)
	Expect(t, len(root.Outputs)).ToBe(2)
	count := 0
	for _, out := range root.Outputs {
		switch out.ID {
		case "stdout_file":
			Expect(t, out.Types[0].Type).ToBe("stdout")
			count = count + 1
		case "stderr_file":
			Expect(t, out.Types[0].Type).ToBe("stderr")
			count = count + 1
		}
	}
	Expect(t, count).ToBe(2)
	Expect(t, root.BaseCommands[0]).ToBe("echo")
	Expect(t, root.Arguments[0].Binding.ValueFrom.Key()).ToBe("foo 1>&2")
	Expect(t, root.Arguments[0].Binding.ShellQuote).ToBe(false)
}
