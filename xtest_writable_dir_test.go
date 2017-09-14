package cwl

import (
	"testing"

	. "github.com/otiai10/mint"
)

func TestDecode_writable_dir(t *testing.T) {
	f := cwl("writable-dir.cwl")
	root := NewCWL()
	err := root.Decode(f)
	Expect(t, err).ToBe(nil)
	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Class).ToBe("CommandLineTool")
	Expect(t, root.Requirements[1].Class).ToBe("InlineJavascriptRequirement")
	Expect(t, root.Requirements[0].Class).ToBe("InitialWorkDirRequirement")
	Expect(t, root.Requirements[0].Listing[0].EntryName).ToBe("emptyWritableDir")
	Expect(t, root.Requirements[0].Listing[0].Writable).ToBe(true)
	Expect(t, root.Requirements[0].Listing[0].Entry).ToBe("$({class: 'Directory', listing: []})")
	Expect(t, len(root.Inputs)).ToBe(0)
	// TODO check specification for this test ID and Type
	Expect(t, root.Outputs[0].ID).ToBe("out")
	Expect(t, root.Outputs[0].Types[0].Type).ToBe("Directory")
	Expect(t, root.Outputs[0].Binding.Glob[0]).ToBe("emptyWritableDir")
	Expect(t, len(root.Arguments)).ToBe(2)
	Expect(t, root.Arguments[0].Value).ToBe("touch")
	Expect(t, root.Arguments[1].Value).ToBe("emptyWritableDir/blurg")
}
