package cwlgotest

import (
	"reflect"
	"testing"

	cwl "github.com/otiai10/cwl.go"
	. "github.com/otiai10/mint"
)

func TestDecode_inline_js(t *testing.T) {
	f := load("inline-js.cwl")
	root := cwl.NewCWL()
	Expect(t, root).TypeOf("*cwl.Root")
	err := root.Decode(f)
	Expect(t, err).ToBe(nil)
	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Class).ToBe("CommandLineTool")
	// TODO test BaseCommand because this file has two baseCommand fields
	//fmt.Println(root.BaseCommands)
	//Expect(t, len(root.BaseCommands)).ToBe(0)
	//Expect(t, root.BaseCommands[0]).ToBe("touch")
	Expect(t, len(root.Requirements)).ToBe(1)
	Expect(t, root.Requirements[0].Class).ToBe("InlineJavascriptRequirement")
	Expect(t, len(root.Requirements)).ToBe(1)
	Expect(t, root.Hints[0].Class).ToBe("DockerRequirement")
	Expect(t, root.Hints[0].DockerPull).ToBe("python:2-slim")
	Expect(t, len(root.Inputs)).ToBe(1)
	Expect(t, root.Inputs[0].ID).ToBe("args.py")
	Expect(t, root.Inputs[0].Types[0].Type).ToBe("File")
	Expect(t, root.Inputs[0].Default.Kind).ToBe(reflect.Map)
	Expect(t, root.Inputs[0].Binding.Position).ToBe(-1)
	Expect(t, len(root.Outputs)).ToBe(1)
	Expect(t, root.Outputs[0].ID).ToBe("args")
	Expect(t, root.Outputs[0].Types[0].Type).ToBe("array")
	Expect(t, root.Outputs[0].Types[0].Items[0].Type).ToBe("string")
	Expect(t, len(root.Arguments)).ToBe(3)
	Expect(t, root.Arguments[0].Binding.Prefix).ToBe("-A")
	Expect(t, root.Arguments[0].Binding.ValueFrom).ToBe("$(1+1)")
	Expect(t, root.Arguments[1].Binding.Prefix).ToBe("-B")
	Expect(t, root.Arguments[1].Binding.ValueFrom).ToBe(`$("/foo/bar/baz".split('/').slice(-1)[0])`)
	Expect(t, root.Arguments[2].Binding.Prefix).ToBe("-C")
	Expect(t, root.Arguments[2].Binding.ValueFrom).ToBe(`${
  var r = [];
  for (var i = 10; i >= 1; i--) {
    r.push(i);
  }
  return r;
}
`)
}
