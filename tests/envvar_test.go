package cwlgotest

import (
	"testing"

	cwl "github.com/otiai10/cwl.go"
	. "github.com/otiai10/mint"
)

func TestDecode_envvar(t *testing.T) {
	f := load("envvar.cwl")
	root := cwl.NewCWL()
	Expect(t, root).TypeOf("*cwl.Root")
	err := root.Decode(f)
	Expect(t, err).ToBe(nil)
	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Class).ToBe("CommandLineTool")
	Expect(t, len(root.Inputs)).ToBe(0)
	Expect(t, len(root.Outputs)).ToBe(0)
	Expect(t, len(root.Requirements)).ToBe(1)
	Expect(t, root.Requirements[0].Class).ToBe("ShellCommandRequirement")
	Expect(t, len(root.Arguments)).ToBe(12)
	Expect(t, root.Arguments[0].Value).ToBe("echo")
	Expect(t, root.Arguments[1].Binding.ValueFrom.Key()).ToBe("\"HOME=$HOME\"")
	Expect(t, root.Arguments[1].Binding.ShellQuote).ToBe(false)
	Expect(t, root.Arguments[2].Binding.ValueFrom.Key()).ToBe("\"TMPDIR=$TMPDIR\"")
	Expect(t, root.Arguments[2].Binding.ShellQuote).ToBe(false)
	Expect(t, root.Arguments[3].Binding.ValueFrom.Key()).ToBe("&&")
	Expect(t, root.Arguments[3].Binding.ShellQuote).ToBe(false)
	Expect(t, root.Arguments[4].Value).ToBe("test")
	Expect(t, root.Arguments[5].Binding.ValueFrom.Key()).ToBe("\"$HOME\"")
	Expect(t, root.Arguments[5].Binding.ShellQuote).ToBe(false)
	Expect(t, root.Arguments[6].Value).ToBe("=")
	Expect(t, root.Arguments[7].Value).ToBe("$(runtime.outdir)")
	Expect(t, root.Arguments[8].Value).ToBe("-a")
	Expect(t, root.Arguments[9].Binding.ValueFrom.Key()).ToBe("\"$TMPDIR\"")
	Expect(t, root.Arguments[9].Binding.ShellQuote).ToBe(false)
	Expect(t, root.Arguments[10].Value).ToBe("=")
	Expect(t, root.Arguments[11].Value).ToBe("$(runtime.tmpdir)")
}
