package cwl

import (
	"testing"

	. "github.com/otiai10/mint"
)

func TestDecode_optional_output(t *testing.T) {
	f := cwl("optional-output.cwl")
	root := NewCWL()
	err := root.Decode(f)
	Expect(t, err).ToBe(nil)
	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Class).ToBe("CommandLineTool")
	Expect(t, root.Doc).ToBe("Print the contents of a file to stdout using 'cat' running in a docker container.")
	Expect(t, root.Hints[0].Class).ToBe("DockerRequirement")
	Expect(t, root.Hints[0].DockerPull).ToBe("debian:wheezy")
	Expect(t, root.Inputs[0].ID).ToBe("file1")
	Expect(t, root.Inputs[0].Types[0].Type).ToBe("File")
	Expect(t, root.Inputs[0].Label).ToBe("Input File")
	Expect(t, root.Inputs[0].Doc).ToBe("The file that will be copied using 'cat'")
	Expect(t, root.Inputs[0].Binding.Position).ToBe(1)
	Expect(t, root.Outputs[0].ID).ToBe("output_file")
	Expect(t, root.Outputs[0].Types[0].Type).ToBe("File")
	Expect(t, root.Outputs[0].Binding.Glob[0]).ToBe("output.txt")
	Expect(t, root.Outputs[0].SecondaryFiles[0].Entry).ToBe(".idx")
	Expect(t, root.Outputs[1].ID).ToBe("optional_file")
	Expect(t, root.Outputs[1].Types[0].Type).ToBe("File?")
	Expect(t, root.Outputs[1].Binding.Glob[0]).ToBe("bumble.txt")
	Expect(t, root.BaseCommands[0]).ToBe("cat")
	Expect(t, root.Stdout).ToBe("output.txt")
}
