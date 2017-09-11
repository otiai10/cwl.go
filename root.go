package cwl

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/otiai10/yaml2json"
)

// NewCWL ...
func NewCWL() *Root {
	root := new(Root)
	root.Hints = Hints{}
	root.RequiredInputs = RequiredInputs{}
	// root.ProvidedInputs = ProvidedInputs{}
	return root
}

// Root ...
type Root struct {
	Version        string
	Class          string
	Hints          Hints
	Doc            string
	BaseCommand    string `json:"baseCommand"`
	Stdout         string
	RequiredInputs RequiredInputs `json:"inputs"`
	// ProvidedInputs ProvidedInputs `json:"-"`
	Outputs Outputs
}

// UnmarshalJSON ...
func (root *Root) UnmarshalJSON(b []byte) error {
	docs := map[string]interface{}{}
	if err := json.Unmarshal(b, &docs); err != nil {
		return err
	}
	for key, val := range docs {
		switch key {
		case "cwlVersion":
			root.Version = val.(string)
		case "class":
			root.Class = val.(string)
		case "hints":
			root.Hints = root.Hints.New(val)
		case "doc":
			root.Doc = val.(string)
		case "baseCommand":
			root.BaseCommand = val.(string)
		case "stdout":
			root.Stdout = val.(string)
		case "inputs":
			root.RequiredInputs = root.RequiredInputs.New(val)
		case "outputs":
			root.Outputs = root.Outputs.New(val)
		}
	}
	return nil
}

// Decode decodes specified io.Reader to this root
func (root *Root) Decode(r io.Reader) error {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	buf, err = yaml2json.Y2J(bytes.NewReader(buf))
	if err != nil {
		return err
	}
	if err = json.Unmarshal(buf, root); err != nil {
		return err
	}
	return nil
}
