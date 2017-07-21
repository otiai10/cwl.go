package cwl

import (
	"io"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// NewCWL ...
func NewCWL() *Root {
	root := new(Root)
	root.RequiredInputs = map[string]RequiredInput{}
	root.ProvidedInputs = map[string]ProvidedInput{}
	return root
}

// Decode ...
func Decode(r io.Reader, root *Root) error {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(buf, root)
	if err != nil {
		return err
	}
	return nil
}
