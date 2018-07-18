package cwlgotest

import (
	"testing"

	cwl "github.com/otiai10/cwl.go"
	. "github.com/otiai10/mint"
)

func TestDecode_env_tool1(t *testing.T) {
	f := load("env-tool1.cwl")
	root := cwl.NewCWL()
	Expect(t, root).TypeOf("*cwl.Root")
	err := root.Decode(f)
	Expect(t, err).ToBe(nil)
	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, len(root.BaseCommands)).ToBe(3)
	Expect(t, root.BaseCommands[0]).ToBe("/bin/sh")
	Expect(t, root.BaseCommands[1]).ToBe("-c")
	Expect(t, root.BaseCommands[2]).ToBe("echo $TEST_ENV")
	Expect(t, len(root.Inputs)).ToBe(1)
	// TODO ignore "in: string'
	Expect(t, len(root.Outputs)).ToBe(1)
	Expect(t, root.Outputs[0].ID).ToBe("out")
	Expect(t, root.Outputs[0].Types[0].Type).ToBe("File")
	Expect(t, root.Outputs[0].Binding.Glob).ToBe([]string{"out"})
}
