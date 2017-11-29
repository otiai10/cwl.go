package cwlgotest

import (
	"reflect"
	"sort"
	"testing"

	cwl "github.com/otiai10/cwl.go"
	. "github.com/otiai10/mint"
)

func TestDecode_cat1_testcli(t *testing.T) {
	f := load("cat1-testcli.cwl")
	root := cwl.NewCWL()
	err := root.Decode(f)
	Expect(t, err).ToBe(nil)

	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Class).ToBe("CommandLineTool")
	Expect(t, root.Doc).ToBe("Print the contents of a file to stdout using 'cat' running in a docker container.")

	Expect(t, root.Hints[0].Class).ToBe("DockerRequirement")
	Expect(t, root.Hints[0].DockerPull).ToBe("python:2-slim")

	sort.Sort(root.Inputs)

	Expect(t, root.Inputs[0].ID).ToBe("args.py")
	Expect(t, root.Inputs[0].Types[0].Type).ToBe("File")
	Expect(t, root.Inputs[0].Default.Kind).ToBe(reflect.Map)
	Expect(t, root.Inputs[0].Binding.Position).ToBe(-1)
	Expect(t, root.Inputs[1].ID).ToBe("numbering")
	Expect(t, root.Inputs[1].Types[0].Type).ToBe("null")
	Expect(t, root.Inputs[1].Types[1].Type).ToBe("boolean")
	Expect(t, root.Inputs[2].ID).ToBe("file1")
	Expect(t, root.Inputs[2].Types[0].Type).ToBe("File")
	Expect(t, root.Inputs[2].Binding.Position).ToBe(1)

	Expect(t, root.Outputs[0].ID).ToBe("args")
	Expect(t, root.Outputs[0].Types[0].Type).ToBe("string[]")

	Expect(t, root.BaseCommands[0]).ToBe("python")
	Expect(t, root.Arguments[0].Value).ToBe("cat")
}
