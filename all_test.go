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
	// TODO inputs and outputs
}
