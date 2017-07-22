package cwl

import (
	"os"
	"reflect"
	"testing"

	. "github.com/otiai10/mint"
)

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
