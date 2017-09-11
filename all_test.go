package cwl

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	. "github.com/otiai10/mint"
)

const cwl_version = "1.0"

func cwl(name string) string {
	return fmt.Sprintf("./testdata/cwl/v%[1]s/v%[1]s/%s", cwl_version, name)
}

func TestDecode(t *testing.T) {
	file, err := os.Open("./testdata/001-1st-tool.yaml")
	if err != nil {
		panic(err)
	}
	root := NewCWL()
	Expect(t, root).TypeOf("*cwl.Root")
	err = root.Decode(file)
	Expect(t, err).ToBe(nil)
	Expect(t, root.Version).ToBe("v1.0")
}

func TestDecodeInputs(t *testing.T) {
	file, err := os.Open("./testdata/001-inputs.yaml")
	if err != nil {
		panic(nil)
	}
	inputs := NewInputs()
	err = inputs.Decode(file)
	Expect(t, err).ToBe(nil)
	Expect(t, inputs["message"]).ToBe(ProvidedInput{Self: "Hello world!", Type: reflect.String})
	Expect(t, inputs["inputfile"].Class).ToBe("File")
	Expect(t, inputs["inputfile"].Path).ToBe("hogee.txt")
}

func TestDecode_bwa_mem_tool(t *testing.T) {
	file, err := os.Open(cwl("bwa-mem-tool.cwl"))
	if err != nil {
		panic(err)
	}
	root := NewCWL()
	Expect(t, root).TypeOf("*cwl.Root")
	err = root.Decode(file)
	Expect(t, err).ToBe(nil)
	Expect(t, root.Version).ToBe("v1.0")
}
