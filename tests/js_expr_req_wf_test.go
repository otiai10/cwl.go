package cwlgotest

import (
        "testing"

        cwl "github.com/otiai10/cwl.go"
        . "github.com/otiai10/mint"
)

func TestDecode_js_expr_req_wf(t *testing.T) {
        f := load("js-expr-req-wf.cwl")
        root := cwl.NewCWL()
        err := root.Decode(f)
        Expect(t, err).ToBe(nil)
        Expect(t, root.Version).ToBe("v1.0")
        Expect(t, len(root.Graphs)).ToBe(2)
        // 0
        Expect(t, root.Graphs[0].ID).ToBe("tool")
        Expect(t, root.Graphs[0].Class).ToBe("CommandLineTool")
        Expect(t, root.Graphs[0].Requirements[0].Class).ToBe("InlineJavascriptRequirement")
        Expect(t, root.Graphs[0].Requirements[0].ExpressionLib[0].Value).ToBe("function foo() { return 2; }")
        Expect(t, len(root.Graphs[0].Inputs)).ToBe(0)
        Expect(t, root.Graphs[0].Arguments[0].Value).ToBe("echo")
        Expect(t, root.Graphs[0].Stdout).ToBe("whatever.txt")
        Expect(t, len(root.Graphs[0].Outputs)).ToBe(1)
        Expect(t, root.Graphs[0].Outputs[0].ID).ToBe("out")
        Expect(t, root.Graphs[0].Outputs[0].Types[0].Type).ToBe("stdout")
        // 1
        Expect(t, root.Graphs[1].ID).ToBe("wf")
        Expect(t, root.Graphs[1].Class).ToBe("Workflow")
        Expect(t, root.Graphs[1].Requirements[0].Class).ToBe("InlineJavascriptRequirement")
        Expect(t, root.Graphs[1].Requirements[0].ExpressionLib[0].Value).ToBe("function bar() { return 1; }")
        Expect(t, len(root.Graphs[1].Inputs)).ToBe(0)
        Expect(t, root.Graphs[1].Outputs[0].ID).ToBe("out")
        Expect(t, root.Graphs[1].Outputs[0].Types[0].Type).ToBe("File")
        Expect(t, root.Graphs[1].Outputs[0].Source[0]).ToBe("tool/out")
        Expect(t, root.Graphs[1].Steps[0].ID).ToBe("tool")
        // Expect(t, root.Graphs[1].Steps[0].Run.Workflow.ID).ToBe("#tool")
        Expect(t, len(root.Graphs[1].Steps[0].In)).ToBe(0)
        // TODO check empty In
        Expect(t, len(root.Graphs[1].Steps[0].Out)).ToBe(1)
        Expect(t, root.Graphs[1].Steps[0].Out[0].ID).ToBe("out")
}
