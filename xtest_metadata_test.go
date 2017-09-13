package cwl

import (
	"testing"

	. "github.com/otiai10/mint"
)

func TestDecode_metadata(t *testing.T) {
	f := cwl("metadata.cwl")
	root := NewCWL()
	Expect(t, root).TypeOf("*cwl.Root")
	err := root.Decode(f)
	Expect(t, err).ToBe(nil)
	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Class).ToBe("CommandLineTool")
	Expect(t, root.Doc).ToBe("Print the contents of a file to stdout using 'cat' running in a docker container.")
	Expect(t, len(root.Hints)).ToBe(1)
	Expect(t, root.Hints).TypeOf("cwl.Hints")
	Expect(t, root.Hints[0].Class).ToBe("DockerRequirement")
	Expect(t, root.Hints[0].DockerPull).ToBe("debian:wheezy")
	Expect(t, len(root.Inputs)).ToBe(2)
	Expect(t, root.Inputs[0].ID).ToBe("file1")
	Expect(t, root.Inputs[0].Types[0].Type).ToBe("File")
	Expect(t, root.Inputs[0].Binding.Position).ToBe(1)
	Expect(t, len(root.Outputs)).ToBe(0)
	Expect(t, len(root.BaseCommands)).ToBe(1)
	Expect(t, root.BaseCommands[0]).ToBe("cat")
	// $namespaces
	Expect(t, len(root.Namespaces)).ToBe(2)
	Expect(t, root.Namespaces[0]["dct"]).ToBe("http://purl.org/dc/terms/")
	Expect(t, root.Namespaces[1]["foaf"]).ToBe("http://xmlns.com/foaf/0.1/")
	// $namespaces
	Expect(t, len(root.Schemas)).ToBe(2)
	Expect(t, root.Schemas[0]).ToBe("foaf.rdf")
	Expect(t, root.Schemas[1]).ToBe("dcterms.rdf")
}
