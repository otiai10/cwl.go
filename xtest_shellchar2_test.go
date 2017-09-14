package cwl

import (
	"testing"

	. "github.com/otiai10/mint"
)

func TestDecode_shellchar2(t *testing.T) {
	f := cwl("shellchar2.cwl")
	root := NewCWL()
	err := root.Decode(f)
	Expect(t, err).ToBe(nil)
	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Class).ToBe("CommandLineTool")
	Expect(t, root.Doc).ToBe("Ensure that `shellQuote: true` is the default behavior when\n" + "ShellCommandRequirement is in effect.\n")
	Expect(t, root.Requirements[0].Class).ToBe("ShellCommandRequirement")
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
	Expect(t, root.Arguments[0].Value).ToBe("foo 1>&2")
}
