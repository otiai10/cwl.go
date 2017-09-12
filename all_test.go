package cwl

import (
	"fmt"
	"os"
	"testing"

	. "github.com/otiai10/mint"
)

const version = "1.0"

// Provides path for testable official .cwl files.
func cwl(name string) string {
	return fmt.Sprintf("./cwl/v%[1]s/v%[1]s/%s", version, name)
}
func TestDecode_bwa_mem_tool(t *testing.T) {
	f, err := os.Open(cwl("bwa-mem-tool.cwl"))
	if err != nil {
		panic(err)
	}
	root := NewCWL()
	Expect(t, root).TypeOf("*cwl.Root")
	err = root.Decode(f)
	Expect(t, err).ToBe(nil)
	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Class).ToBe("CommandLineTool")
	Expect(t, root.Hints).TypeOf("cwl.Hints")
	Expect(t, root.Hints[0]["class"]).ToBe("ResourceRequirement")
	Expect(t, root.Hints[0]["coresMin"]).ToBe(float64(2))

	Expect(t, len(root.RequiredInputs)).ToBe(int(5))
	Expect(t, root.RequiredInputs[0]).TypeOf("cwl.RequiredInput")
	Expect(t, root.RequiredInputs[0].ID).ToBe("reference")
	Expect(t, root.RequiredInputs[0].Type.Type).ToBe("File")
	Expect(t, root.RequiredInputs[0].Binding.Position).ToBe(2)
	Expect(t, root.RequiredInputs[1].ID).ToBe("reads")
	Expect(t, root.RequiredInputs[1].Type.Type).ToBe("array")
	Expect(t, root.RequiredInputs[1].Type.Items).ToBe("File")
	Expect(t, root.RequiredInputs[1].Binding.Position).ToBe(3)
	Expect(t, root.RequiredInputs[2].Binding.Prefix).ToBe("-m")
	Expect(t, root.RequiredInputs[3].Binding.Separator).ToBe(",")
	Expect(t, root.RequiredInputs[4].Default.Class).ToBe("File")
	Expect(t, root.RequiredInputs[4].Default.Location).ToBe("args.py")
	Expect(t, root.Outputs[0].ID).ToBe("sam")
	Expect(t, root.Outputs[0].Types[0].Type).ToBe("null")
	Expect(t, root.Outputs[1].ID).ToBe("args")
	Expect(t, root.Outputs[1].Types[0].Type).ToBe("array")
	Expect(t, root.Outputs[1].Types[0].Items).ToBe("string")
}
func TestDecode_cat3_nodocker(t *testing.T) {
	f, err := os.Open(cwl("cat3-nodocker.cwl"))
	if err != nil {
		panic(err)
	}
	root := NewCWL()
	Expect(t, root).TypeOf("*cwl.Root")
	err = root.Decode(f)
	Expect(t, err).ToBe(nil)
	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Doc).ToBe("Print the contents of a file to stdout using 'cat'.")
	Expect(t, root.Class).ToBe("CommandLineTool")
	Expect(t, root.BaseCommand).ToBe("cat")
	Expect(t, root.Stdout).ToBe("output.txt")
	Expect(t, len(root.RequiredInputs)).ToBe(int(1))
	Expect(t, root.RequiredInputs[0].ID).ToBe("file1")
	Expect(t, root.RequiredInputs[0].Type.Type).ToBe("File")
	Expect(t, root.RequiredInputs[0].Label).ToBe("Input File")
	Expect(t, root.RequiredInputs[0].Doc).ToBe("The file that will be copied using 'cat'")
	Expect(t, root.RequiredInputs[0].Binding.Position).ToBe(1)
}
func TestDecode_cat3_tool_mediumcut(t *testing.T) {
	f, err := os.Open(cwl("cat3-tool-mediumcut.cwl"))
	if err != nil {
		panic(err)
	}
	root := NewCWL()
	Expect(t, root).TypeOf("*cwl.Root")
	err = root.Decode(f)
	Expect(t, err).ToBe(nil)
	Expect(t, root.Version).ToBe("v1.0")
	Expect(t, root.Doc).ToBe("Print the contents of a file to stdout using 'cat' running in a docker container.")
	Expect(t, root.Class).ToBe("CommandLineTool")
	Expect(t, root.BaseCommand).ToBe("cat")
	Expect(t, root.Stdout).ToBe("cat-out")
	Expect(t, root.Hints).TypeOf("cwl.Hints")
	Expect(t, root.Hints[0]["class"]).ToBe("DockerRequirement")
	Expect(t, root.Hints[0]["dockerPull"]).ToBe("debian:wheezy")
	Expect(t, len(root.RequiredInputs)).ToBe(int(1))
	Expect(t, root.RequiredInputs[0].ID).ToBe("file1")
	Expect(t, root.RequiredInputs[0].Type.Type).ToBe("File")
	Expect(t, root.RequiredInputs[0].Label).ToBe("Input File")
	Expect(t, root.RequiredInputs[0].Doc).ToBe("The file that will be copied using 'cat'")
	Expect(t, root.RequiredInputs[0].Binding.Position).ToBe(1)
}
