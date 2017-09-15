package cwlgotest

import (
	"testing"

	cwl "github.com/otiai10/cwl.go"
	. "github.com/otiai10/mint"
)

func TestDecode_dynresreq(t *testing.T) {
	f := load("dynresreq.cwl")
	root := cwl.NewCWL()
	err := root.Decode(f)
	Expect(t, err).ToBe(nil)
	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Class).ToBe("CommandLineTool")
	Expect(t, root.Requirements[0].Class).ToBe("ResourceRequirement")
	// TODO check CoresMin and CoresMax
	//Expect(t, root.Requirements[0].CoreMin).ToBe("$(inputs.special_file.size)")
	Expect(t, root.Inputs[0].ID).ToBe("special_file")
	Expect(t, root.Inputs[0].Types[0].Type).ToBe("File")
	Expect(t, root.Outputs[0].ID).ToBe("output")
	Expect(t, root.Outputs[0].Types[0].Type).ToBe("stdout")
	Expect(t, root.BaseCommands[0]).ToBe("echo")
	Expect(t, root.Stdout).ToBe("cores.txt")
	Expect(t, root.Arguments[0].Value).ToBe("$(runtime.cores)")
}
