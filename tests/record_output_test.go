package cwlgotest

import (
	"testing"

	cwl "github.com/otiai10/cwl.go"
	. "github.com/otiai10/mint"
)

func TestDecode_record_output(t *testing.T) {
	f := load("record-output.cwl")
	root := cwl.NewCWL()
	err := root.Decode(f)
	Expect(t, err).ToBe(nil)
	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Class).ToBe("CommandLineTool")
	Expect(t, root.Inputs[0].ID).ToBe("irec")
	Expect(t, root.Requirements[0].Class).ToBe("ShellCommandRequirement")
	Expect(t, root.Inputs[0].Types[0].Type).ToBe("record")
	// TODO type allows name ?
	Expect(t, root.Inputs[0].Types[0].Fields[0].Name).ToBe("ifoo")
	Expect(t, root.Inputs[0].Types[0].Fields[0].Types[0].Type).ToBe("File")
	Expect(t, root.Inputs[0].Types[0].Fields[0].Binding.Position).ToBe(2)
	Expect(t, root.Inputs[0].Types[0].Fields[1].Name).ToBe("ibar")
	Expect(t, root.Inputs[0].Types[0].Fields[1].Types[0].Type).ToBe("File")
	Expect(t, root.Inputs[0].Types[0].Fields[1].Binding.Position).ToBe(6)
	Expect(t, root.Outputs[0].ID).ToBe("orec")
	Expect(t, root.Outputs[0].Types[0].Type).ToBe("record")
	Expect(t, root.Outputs[0].Types[0].Fields[0].Name).ToBe("ofoo")
	Expect(t, root.Outputs[0].Types[0].Fields[0].Types[0].Type).ToBe("File")
	Expect(t, root.Outputs[0].Types[0].Fields[0].Binding.Glob[0]).ToBe("foo")
	Expect(t, root.Outputs[0].Types[0].Fields[1].Name).ToBe("obar")
	Expect(t, root.Outputs[0].Types[0].Fields[1].Types[0].Type).ToBe("File")
	Expect(t, root.Outputs[0].Types[0].Fields[1].Binding.Glob[0]).ToBe("bar")
	Expect(t, root.Arguments[0].Binding.ValueFrom.Key()).ToBe("cat")
	Expect(t, root.Arguments[0].Binding.Position).ToBe(1)
	Expect(t, root.Arguments[1].Binding.ValueFrom.Key()).ToBe("> foo")
	Expect(t, root.Arguments[1].Binding.Position).ToBe(3)
	Expect(t, root.Arguments[1].Binding.ShellQuote).ToBe(false)
	Expect(t, root.Arguments[2].Binding.ValueFrom.Key()).ToBe("&&")
	Expect(t, root.Arguments[2].Binding.Position).ToBe(4)
	Expect(t, root.Arguments[2].Binding.ShellQuote).ToBe(false)
	Expect(t, root.Arguments[3].Binding.ValueFrom.Key()).ToBe("cat")
	Expect(t, root.Arguments[3].Binding.Position).ToBe(5)
	Expect(t, root.Arguments[4].Binding.ValueFrom.Key()).ToBe("> bar")
	Expect(t, root.Arguments[4].Binding.Position).ToBe(7)
	Expect(t, root.Arguments[4].Binding.ShellQuote).ToBe(false)
}
