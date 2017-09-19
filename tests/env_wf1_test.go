package cwlgotest

import (
	"testing"

	cwl "github.com/otiai10/cwl.go"
	. "github.com/otiai10/mint"
)

func TestDecode_env_wf1(t *testing.T) {
	f := load("env-wf1.cwl")
	root := cwl.NewCWL()
	Expect(t, root).TypeOf("*cwl.Root")
	err := root.Decode(f)
	Expect(t, err).ToBe(nil)
	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Class).ToBe("Workflow")
	Expect(t, len(root.Inputs)).ToBe(1)
	// TODO in: string
	Expect(t, len(root.Outputs)).ToBe(1)
	Expect(t, root.Outputs[0].ID).ToBe("out")
	Expect(t, root.Outputs[0].Types[0].Type).ToBe("File")
	Expect(t, root.Outputs[0].Source).ToBe([]string{"step1/out"})
	Expect(t, len(root.Requirements)).ToBe(1)
	Expect(t, root.Requirements[0].Class).ToBe("EnvVarRequirement")
	Expect(t, root.Requirements[0].EnvDef[0].Name).ToBe("TEST_ENV")
	Expect(t, root.Requirements[0].EnvDef[0].Value).ToBe(`override`)
}
