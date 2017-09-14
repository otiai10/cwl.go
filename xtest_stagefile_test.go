package cwl

import (
	"testing"

	. "github.com/otiai10/mint"
)

func TestDecode_stagefile(t *testing.T) {
	f := cwl("stagefile.cwl")
	root := NewCWL()
	err := root.Decode(f)
	Expect(t, err).ToBe(nil)
	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Class).ToBe("CommandLineTool")
	Expect(t, root.Hints[0].Class).ToBe("DockerRequirement")
	Expect(t, root.Hints[0].DockerPull).ToBe("python:2-slim")
	Expect(t, root.Requirements[0].Class).ToBe("InitialWorkDirRequirement")
	Expect(t, root.Requirements[0].Listing[0].Entry).ToBe("$(inputs.infile)")
	Expect(t, root.Requirements[0].Listing[0].EntryName).ToBe("bob.txt")
	Expect(t, root.Requirements[0].Listing[0].Writable).ToBe(true)
	Expect(t, len(root.Inputs)).ToBe(1)
	Expect(t, root.Inputs[0].ID).ToBe("infile")
	Expect(t, root.Inputs[0].Types[0].Type).ToBe("File")
	Expect(t, len(root.Outputs)).ToBe(1)
	Expect(t, root.Outputs[0].ID).ToBe("outfile")
	Expect(t, root.Outputs[0].Types[0].Type).ToBe("File")
	Expect(t, root.Outputs[0].Binding.Glob[0]).ToBe("bob.txt")
	Expect(t, root.BaseCommands[0]).ToBe("python2")
	Expect(t, root.Arguments[0].Value).ToBe("-c")
	Expect(t, root.Arguments[1].Value).ToBe(`f = open("bob.txt", "r+")
f.seek(8)
f.write("Bob.    ")
f.close()
`)
}
