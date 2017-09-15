package cwlgotest

import (
	"reflect"
	"testing"

	cwl "github.com/otiai10/cwl.go"
	. "github.com/otiai10/mint"
)

func TestDecode_bwa_mem_tool(t *testing.T) {
	f := load("bwa-mem-tool.cwl")
	root := cwl.NewCWL()
	Expect(t, root).TypeOf("*cwl.Root")
	err := root.Decode(f)
	Expect(t, err).ToBe(nil)
	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Class).ToBe("CommandLineTool")
	Expect(t, root.Hints).TypeOf("cwl.Hints")
	Expect(t, root.Hints[0].Class).ToBe("ResourceRequirement")
	Expect(t, root.Hints[0].CoresMin).ToBe(2)

	Expect(t, len(root.Inputs)).ToBe(5)
	Expect(t, root.Inputs[0]).TypeOf("cwl.Input")
	Expect(t, root.Inputs[0].ID).ToBe("reference")
	Expect(t, root.Inputs[0].Types[0].Type).ToBe("File")
	Expect(t, root.Inputs[0].Binding.Position).ToBe(2)
	Expect(t, root.Inputs[1].ID).ToBe("reads")
	Expect(t, root.Inputs[1].Types[0].Type).ToBe("array")
	Expect(t, root.Inputs[1].Types[0].Items[0].Type).ToBe("File")
	Expect(t, root.Inputs[1].Binding.Position).ToBe(3)
	Expect(t, root.Inputs[2].Binding.Prefix).ToBe("-m")
	Expect(t, root.Inputs[3].Binding.Separator).ToBe(",")
	Expect(t, root.Inputs[4].Default.Kind).ToBe(reflect.Map)
	Expect(t, root.Outputs[0].ID).ToBe("sam")
	Expect(t, root.Outputs[0].Types[0].Type).ToBe("null")
	Expect(t, root.Outputs[0].Types[1].Type).ToBe("File")
	Expect(t, root.Outputs[0].Binding.Glob).ToBe([]string{"output.sam"})
	Expect(t, root.Outputs[1].ID).ToBe("args")
	Expect(t, root.Outputs[1].Types[0].Type).ToBe("array")
	Expect(t, root.Outputs[1].Types[0].Items[0].Type).ToBe("string")
}
