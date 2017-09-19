package cwlgotest

import (
	"testing"

	cwl "github.com/otiai10/cwl.go"
	. "github.com/otiai10/mint"
)

func TestDecode_env_tool2(t *testing.T) {
	f := load("env-tool2.cwl")
	root := cwl.NewCWL()
	Expect(t, root).TypeOf("*cwl.Root")
	err := root.Decode(f)
	Expect(t, err).ToBe(nil)
	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Class).ToBe("CommandLineTool")
	Expect(t, len(root.Hints)).ToBe(1)
	Expect(t, root.Hints).TypeOf("cwl.Hints")
	Expect(t, root.Hints[0].Class).ToBe("EnvVarRequirement")
	Expect(t, root.Hints[0].Envs[0].Name).ToBe("TEST_ENV")
	Expect(t, root.Hints[0].Envs[0].Value).ToBe("$(inputs.in)")
	Expect(t, len(root.Inputs)).ToBe(1)
	// TODO in: string
	Expect(t, len(root.Outputs)).ToBe(1)
	Expect(t, len(root.BaseCommands)).ToBe(3)
	Expect(t, root.BaseCommands[0]).ToBe("/bin/bash")
	Expect(t, root.BaseCommands[1]).ToBe("-c")
	Expect(t, root.BaseCommands[2]).ToBe("echo $TEST_ENV")
}
