package cwlgotest

import (
	"reflect"
	"sort"
	"testing"

	cwl "github.com/otiai10/cwl.go"
	. "github.com/otiai10/mint"
)

func TestDecode_binding_test(t *testing.T) {
	f := load("binding-test.cwl")
	root := cwl.NewCWL()
	err := root.Decode(f)
	Expect(t, err).ToBe(nil)

	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Class).ToBe("CommandLineTool")

	Expect(t, root.Hints[0].Class).ToBe("DockerRequirement")
	Expect(t, root.Hints[0].DockerPull).ToBe("python:2-slim")

	sort.Sort(root.Inputs)
	Expect(t, root.Inputs[0].ID).ToBe("#args.py")
	Expect(t, root.Inputs[0].Types[0].Type).ToBe("File")
	Expect(t, root.Inputs[0].Default.Kind).ToBe(reflect.Map)
	Expect(t, root.Inputs[0].Binding.Position).ToBe(-1)
	Expect(t, root.Inputs[1].ID).ToBe("reference")
	Expect(t, root.Inputs[1].Types[0].Type).ToBe("File")
	Expect(t, root.Inputs[1].Binding.Position).ToBe(2)
	Expect(t, root.Inputs[2].ID).ToBe("reads")
	Expect(t, root.Inputs[2].Types[0].Type).ToBe("array")
	Expect(t, root.Inputs[2].Types[0].Items[0].Type).ToBe("File")
	Expect(t, root.Inputs[2].Types[0].Binding.Prefix).ToBe("-YYY")
	Expect(t, root.Inputs[2].Binding.Position).ToBe(3)
	Expect(t, root.Inputs[2].Binding.Prefix).ToBe("-XXX")

	Expect(t, root.Outputs[0].ID).ToBe("args")
	Expect(t, root.Outputs[0].Types[0].Type).ToBe("string[]")
}
