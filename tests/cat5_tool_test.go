package cwlgotest

import (
	"testing"

	cwl "github.com/otiai10/cwl.go"
	. "github.com/otiai10/mint"
)

func TestDecode_cat5_tool(t *testing.T) {
	f := load("cat5-tool.cwl")
	root := cwl.NewCWL()
	Expect(t, root).TypeOf("*cwl.Root")
	err := root.Decode(f)
	Expect(t, err).ToBe(nil)
	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Class).ToBe("CommandLineTool")
	Expect(t, root.Doc).ToBe("Print the contents of a file to stdout using 'cat' running in a docker container.")
	Expect(t, len(root.Hints)).ToBe(2)
	Expect(t, root.Hints).TypeOf("cwl.Hints")
	Expect(t, root.Hints[0].Class).ToBe("DockerRequirement")
	Expect(t, root.Hints[0].DockerPull).ToBe("debian:wheezy")
	Expect(t, root.Hints[1].Class).ToBe("ex:BlibberBlubberFakeRequirement")
	Expect(t, root.Hints[1].FakeField).ToBe("fraggleFroogle")
	Expect(t, len(root.Inputs)).ToBe(1)
	Expect(t, root.Inputs[0].ID).ToBe("file1")
	Expect(t, root.Inputs[0].Types[0].Type).ToBe("File")
	Expect(t, root.Inputs[0].Label).ToBe("Input File")
	Expect(t, root.Inputs[0].Doc).ToBe("The file that will be copied using 'cat'")
	Expect(t, root.Inputs[0].Binding.Position).ToBe(1)
	Expect(t, len(root.Outputs)).ToBe(1)
	Expect(t, root.Outputs[0].ID).ToBe("output_file")
	Expect(t, root.Outputs[0].Types[0].Type).ToBe("File")
	Expect(t, root.Outputs[0].Binding.Glob).ToBe([]string{"output.txt"})
	Expect(t, len(root.BaseCommands)).ToBe(1)
	Expect(t, root.BaseCommands[0]).ToBe("cat")
	Expect(t, root.Stdout).ToBe("output.txt")
	// $namespaces
	Expect(t, len(root.Namespaces)).ToBe(1)
	Expect(t, root.Namespaces[0]["ex"]).ToBe("http://example.com/")
}
