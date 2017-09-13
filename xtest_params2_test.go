package cwl

import (
	"testing"

	. "github.com/otiai10/mint"
)

func TestDecode_params2(t *testing.T) {
	f := cwl("params2.cwl")
	root := NewCWL()
	err := root.Decode(f)
	Expect(t, err).ToBe(nil)
	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Class).ToBe("CommandLineTool")
	Expect(t, root.Inputs[0].ID).ToBe("bar")
	Expect(t, root.Inputs[0].Types[0].Type).ToBe("Any")
	// TODO check specification for this test ID and Type
	Expect(t, root.Outputs[0].ID).ToBe("$import")
	Expect(t, root.Outputs[0].Types[0].Type).ToBe("params_inc.yml")
	Expect(t, root.BaseCommands[0]).ToBe("true")
}
