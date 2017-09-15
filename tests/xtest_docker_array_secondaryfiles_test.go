package cwlgotest

import (
	"testing"

	cwl "github.com/otiai10/cwl.go"
	. "github.com/otiai10/mint"
)

func TestDecode_docker_array_secondaryfiles(t *testing.T) {
	f := load("docker-array-secondaryfiles.cwl")
	root := cwl.NewCWL()
	err := root.Decode(f)
	Expect(t, err).ToBe(nil)
	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Class).ToBe("CommandLineTool")
	Expect(t, len(root.Requirements)).ToBe(3)
	Expect(t, root.Requirements[0].Class).ToBe("DockerRequirement")
	Expect(t, root.Requirements[0].DockerPull).ToBe("debian:8")
	Expect(t, root.Requirements[1].Class).ToBe("InlineJavascriptRequirement")
	Expect(t, root.Requirements[2].Class).ToBe("ShellCommandRequirement")
	Expect(t, len(root.Inputs)).ToBe(1)
	Expect(t, root.Inputs[0].ID).ToBe("fasta_path")
	Expect(t, root.Inputs[0].Types[0].Type).ToBe("array")
	// TODO items: File
	Expect(t, root.Inputs[0].SecondaryFiles[0].Entry).ToBe(".fai")
	Expect(t, root.Outputs[0].ID).ToBe("bai_list")
	Expect(t, root.Outputs[0].Types[0].Type).ToBe("File")
	Expect(t, root.Outputs[0].Binding.Glob[0]).ToBe("fai.list")
	Expect(t, root.Arguments[0].Binding.ValueFrom).ToBe(`${ var fai_list = ""; for (var i = 0; i < inputs.fasta_path.length; i ++) { fai_list += " cat " + inputs.fasta_path[i].path +".fai" + " >> fai.list && " } return fai_list.slice(0,-3) }`)
	Expect(t, root.Arguments[0].Binding.Position).ToBe(1)
	Expect(t, root.Arguments[0].Binding.ShellQuote).ToBe(false)
	Expect(t, len(root.BaseCommands)).ToBe(0)
}
