package cwlgotest

import (
        "testing"

        cwl "github.com/otiai10/cwl.go"
        . "github.com/otiai10/mint"
)

func TestDecode_default_path(t *testing.T) {
        f := load("default_path.cwl")
        root := cwl.NewCWL()
        Expect(t, root).TypeOf("*cwl.Root")
        err := root.Decode(f)
        Expect(t, err).ToBe(nil)
        Expect(t, root.Version).ToBe("v1.0")
        Expect(t, root.Class).ToBe("CommandLineTool")
        Expect(t, len(root.Inputs)).ToBe(1)
        Expect(t, root.Inputs[0].ID).ToBe("file1")
        Expect(t, root.Inputs[0].Types[0].Type).ToBe("File")
        // TODO support default: section
        // TODO support outputs: []
        Expect(t, len(root.Arguments)).ToBe(2)
        Expect(t, root.Arguments[0].Value).ToBe("cat")
        Expect(t, root.Arguments[1].Value).ToBe("$(inputs.file1.path)")
}

