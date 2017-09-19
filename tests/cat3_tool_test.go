package cwlgotest

import (
	"testing"

	cwl "github.com/otiai10/cwl.go"
	. "github.com/otiai10/mint"
)

func TestDecode_cat3_tool(t *testing.T) {
	f := load("cat3-tool.cwl")
	root := cwl.NewCWL()
	Expect(t, root).TypeOf("*cwl.Root")
	err := root.Decode(f)
	Expect(t, err).ToBe(nil)
	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Doc).ToBe("Print the contents of a file to stdout using 'cat' running in a docker container.")
	Expect(t, root.Class).ToBe("CommandLineTool")
	Expect(t, len(root.BaseCommands)).ToBe(1)
	Expect(t, root.BaseCommands[0]).ToBe("cat")
	Expect(t, root.Stdout).ToBe("output.txt")
	Expect(t, root.Hints).TypeOf("cwl.Hints")
	Expect(t, root.Hints[0].Class).ToBe("DockerRequirement")
	Expect(t, root.Hints[0].DockerPull).ToBe("debian:wheezy")
	Expect(t, len(root.Inputs)).ToBe(1)
	Expect(t, root.Inputs[0].ID).ToBe("file1")
	Expect(t, root.Inputs[0].Types[0].Type).ToBe("File")
	Expect(t, root.Inputs[0].Label).ToBe("Input File")
	Expect(t, root.Inputs[0].Doc).ToBe("The file that will be copied using 'cat'")
	Expect(t, root.Inputs[0].Binding.Position).ToBe(1)
}
