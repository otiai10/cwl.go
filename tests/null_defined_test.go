package cwlgotest

import (
        "testing"

        cwl "github.com/otiai10/cwl.go"
        . "github.com/otiai10/mint"
)

func TestDecode_null_defined(t *testing.T) {
        f := load("null-defined.cwl")
        root := cwl.NewCWL()
        err := root.Decode(f)
        Expect(t, err).ToBe(nil)
        Expect(t, root.Version).ToBe("v1.0")
        Expect(t, root.Class).ToBe("CommandLineTool")
        Expect(t, len(root.Requirements)).ToBe(1)
        Expect(t, root.Requirements[0].Class).ToBe("InlineJavascriptRequirement")
        Expect(t, len(root.Inputs)).ToBe(1)
        Expect(t, root.Inputs[0].ID).ToBe("file1")
        Expect(t, root.Inputs[0].Types[0].Type).ToBe("File?")
        Expect(t, root.Outputs[0].ID).ToBe("out")
        Expect(t, root.Outputs[0].Types[0].Type).ToBe("string")
        Expect(t, root.Outputs[0].Binding.Glob[0]).ToBe("out.txt")
        Expect(t, root.Outputs[0].Binding.Contents).ToBe(false)
        Expect(t, root.Outputs[0].Binding.Eval).ToBe("$(self[0].contents)")
        Expect(t, root.Stdout).ToBe("out.txt")
        Expect(t, len(root.Arguments)).ToBe(2)
        Expect(t, root.Arguments[0].Value).ToBe("echo")
        Expect(t, root.Arguments[1].Value).ToBe(`$(inputs.file1 === null ? "t" : "f")`)
}
