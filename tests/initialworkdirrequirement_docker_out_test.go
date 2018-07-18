package cwlgotest

import (
	"testing"

	cwl "github.com/otiai10/cwl.go"
	. "github.com/otiai10/mint"
)

func TestDecode_initialworkdirrequirement_docker_out(t *testing.T) {
	f := load("initialworkdirrequirement-docker-out.cwl")
	root := cwl.NewCWL()
	Expect(t, root).TypeOf("*cwl.Root")
	err := root.Decode(f)
	Expect(t, err).ToBe(nil)
	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Class).ToBe("CommandLineTool")
	Expect(t, len(root.Inputs)).ToBe(1)
	Expect(t, root.Inputs[0].ID).ToBe("INPUT")
	Expect(t, root.Inputs[0].Types[0].Type).ToBe("File")
	Expect(t, len(root.Outputs)).ToBe(1)
	Expect(t, root.Outputs[0].ID).ToBe("OUTPUT")
	Expect(t, root.Outputs[0].Types[0].Type).ToBe("File")
	Expect(t, root.Outputs[0].Binding.Glob).ToBe([]string{"$(inputs.INPUT.basename)"})
	Expect(t, root.Outputs[0].SecondaryFiles[0].Entry).ToBe(".fai")
	// TODO outputs
	Expect(t, len(root.Requirements)).ToBe(2)
	Expect(t, root.Requirements[0].Class).ToBe("DockerRequirement")
	Expect(t, root.Requirements[0].DockerPull).ToBe("debian:stretch-slim")
	Expect(t, root.Requirements[1].Class).ToBe("InitialWorkDirRequirement")
	Expect(t, root.Requirements[1].Listing[0].Location).ToBe("$(inputs.INPUT)")
	// TODO: fix "Alias.Key()"
	Expect(t, root.Arguments[0].Binding.ValueFrom.Key()).ToBe("inputs.INPUT.basename).fai")
	// TODO test against "position" but currently just put 0 is failed
	Expect(t, len(root.BaseCommands)).ToBe(1)
	Expect(t, root.BaseCommands[0]).ToBe("touch")
}
