package yaml2json

import (
	"os"
	"testing"

	. "github.com/otiai10/mint"
)

func TestY2J(t *testing.T) {
	src, err := os.Open("./testdata/001.yaml")
	if err != nil {
		panic(err)
	}
	b, err := Y2J(src)
	Expect(t, err).ToBe(nil)
	Expect(t, string(b)).ToBe(`{"baz":[100,200],"foo":"bar","spam":{"ham":[2,{"age":17,"name":"otiai10"},"5"]}}`)
}

func TestJ2Y(t *testing.T) {
	src, err := os.Open("./testdata/001.json")
	if err != nil {
		panic(err)
	}
	b, err := J2Y(src)
	Expect(t, err).ToBe(nil)
	Expect(t, string(b)).ToBe(`baz:
- 100
- 200
foo: bar
spam:
  ham:
  - 2
  - age: 17
    name: otiai10
  - "5"
`)
}
