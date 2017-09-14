package cwl

import (
	"testing"

	. "github.com/otiai10/mint"
)

func TestDecode_shellchar(t *testing.T) {
	f := cwl("shellchar.cwl")
	root := NewCWL()
	err := root.Decode(f)
	Expect(t, err).ToBe(nil)
	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Class).ToBe("CommandLineTool")
	Expect(t, root.Doc).ToBe("Ensure that arguments containing shell directives are not interpreted and\n" + "that `shellQuote: false` has no effect when ShellCommandRequirement is not in\n" + "effect.\n")
	Expect(t, len(root.Inputs)).ToBe(0)
	Expect(t, len(root.Outputs)).ToBe(2)
	Expect(t, root.Outputs[0].ID).ToBe("stdout_file")
	Expect(t, root.Outputs[0].Types[0].Type).ToBe("stdout")
	Expect(t, root.Outputs[1].ID).ToBe("stderr_file")
	Expect(t, root.Outputs[1].Types[0].Type).ToBe("stderr")
	Expect(t, root.BaseCommands[0]).ToBe("echo")
	Expect(t, root.Arguments[0].Binding.ValueFrom).ToBe("foo 1>&2")
	Expect(t, root.Arguments[0].Binding.ShellQuote).ToBe(false)
}
