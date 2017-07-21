package cwl

import (
	"os"
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
	err = Decode(file, root)
	Expect(t, err).ToBe(nil)
	Expect(t, root.Version).ToBe("v1.0")
}
