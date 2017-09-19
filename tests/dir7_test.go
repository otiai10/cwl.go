package cwlgotest

import (
	"testing"

	cwl "github.com/otiai10/cwl.go"
	. "github.com/otiai10/mint"
)

func TestDecode_dir7(t *testing.T) {
	f := load("dir7.cwl")
	root := cwl.NewCWL()
	err := root.Decode(f)
	Expect(t, err).ToBe(nil)

	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Class).ToBe("ExpressionTool")

	Expect(t, root.Requirements[0].Class).ToBe("InlineJavascriptRequirement")

	Expect(t, root.Inputs[0].ID).ToBe("files")
	Expect(t, root.Inputs[0].Types[0].Type).ToBe("File[]")
	Expect(t, root.Outputs[0].ID).ToBe("dir")
	Expect(t, root.Outputs[0].Types[0].Type).ToBe("Directory")
	Expect(t, root.Expression).ToBe(`${
return {"dir": {"class": "Directory", "basename": "a_directory", "listing": inputs.files}};
}`)
}
