package cwlgotest

import (
	"testing"

	cwl "github.com/otiai10/cwl.go"
	. "github.com/otiai10/mint"
)

func TestDecode_file_literal_ex(t *testing.T) {
	f := load("file-literal-ex.cwl")
	root := cwl.NewCWL()
	err := root.Decode(f)
	Expect(t, err).ToBe(nil)
	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Class).ToBe("ExpressionTool")
	Expect(t, len(root.Inputs)).ToBe(0)
	// TODO check specification for this test ID and Type
	Expect(t, root.Outputs[0].ID).ToBe("lit")
	Expect(t, root.Outputs[0].Types[0].Type).ToBe("File")
	Expect(t, root.Expression).ToBe(`${
return {"lit": {"class": "File", "basename": "a_file", "contents": "Hello file literal."}};
}`)
}
