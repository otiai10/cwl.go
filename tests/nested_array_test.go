package cwlgotest

import (
        "testing"

        cwl "github.com/otiai10/cwl.go"
        . "github.com/otiai10/mint"
)

func TestDecode_nested_array(t *testing.T) {
        f := load("nested-array.cwl")
        root := cwl.NewCWL()
        err := root.Decode(f)
        Expect(t, err).ToBe(nil)
        Expect(t, root.Version).ToBe("v1.0")
        Expect(t, root.Class).ToBe("CommandLineTool")
        Expect(t, root.BaseCommands[0]).ToBe("echo")
        Expect(t, len(root.Inputs)).ToBe(1)
        Expect(t, root.Inputs[0].ID).ToBe("letters")
        Expect(t, root.Inputs[0].Types[0].Type).ToBe("array")
        Expect(t, root.Inputs[0].Types[0].Items[0].Type).ToBe("array")
        Expect(t, root.Inputs[0].Types[0].Items[0].Items[0].Type).ToBe("string")
        Expect(t, root.Inputs[0].Binding.Position).ToBe(1)
        Expect(t, root.Stdout).ToBe("echo.txt")
        Expect(t, root.Outputs[0].ID).ToBe("echo")
        Expect(t, root.Outputs[0].Types[0].Type).ToBe("stdout")
}
