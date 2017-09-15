package cwlgotest

import (
	"testing"

	cwl "github.com/otiai10/cwl.go"
	. "github.com/otiai10/mint"
)

func TestDecode_formattest3(t *testing.T) {
	f := load("formattest3.cwl")
	root := cwl.NewCWL()
	Expect(t, root).TypeOf("*cwl.Root")
	err := root.Decode(f)
	Expect(t, err).ToBe(nil)
	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Class).ToBe("CommandLineTool")
	// $namespaces
	Expect(t, len(root.Namespaces)).ToBe(2)
	Expect(t, root.Namespaces[0]["edam"]).ToBe("http://edamontology.org/")
	Expect(t, root.Namespaces[1]["gx"]).ToBe("http://galaxyproject.org/formats/")
	// $namespaces
	Expect(t, len(root.Schemas)).ToBe(2)
	Expect(t, root.Schemas[0]).ToBe("EDAM.owl")
	Expect(t, root.Schemas[1]).ToBe("gx_edam.ttl")
	Expect(t, root.Doc).ToBe("Reverse each line using the `rev` command")
	Expect(t, root.Hints[0].Class).ToBe("DockerRequirement")
	Expect(t, root.Hints[0].DockerPull).ToBe("debian:wheezy")
	Expect(t, len(root.Inputs)).ToBe(1)
	Expect(t, root.Inputs[0].ID).ToBe("input")
	Expect(t, root.Inputs[0].Types[0].Type).ToBe("File")
	Expect(t, root.Inputs[0].Format).ToBe("gx:fasta")
	Expect(t, root.Inputs[0].Binding.Position).ToBe(0)
	Expect(t, len(root.Outputs)).ToBe(1)
	Expect(t, root.Outputs[0].ID).ToBe("output")
	Expect(t, root.Outputs[0].Types[0].Type).ToBe("File")
	Expect(t, root.Outputs[0].Binding.Glob).ToBe([]string{"output.txt"})
	Expect(t, root.Outputs[0].Format).ToBe("$(inputs.input.format)")
	Expect(t, len(root.BaseCommands)).ToBe(1)
	Expect(t, root.BaseCommands[0]).ToBe("rev")
	Expect(t, root.Stdout).ToBe("output.txt")
}
