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
	root.BaseCommands = BaseCommands{}
	root.Hints = Hints{}
	root.Inputs = Inputs{}
	// root.ProvidedInputs = ProvidedInputs{}
	return root
}

// Root ...
type Root struct {
	Version        string
	Class          string
	Hints          Hints
	Doc            string
	BaseCommands   BaseCommands
	Arguments      Arguments
	Namespaces     Namespaces
	Stdin          string
	Stdout         string
	Inputs Inputs `json:"inputs"`
	// ProvidedInputs ProvidedInputs `json:"-"`
	Outputs      Outputs
	Requirements Requirements
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
			root.BaseCommands = root.BaseCommands.New(val)
		case "arguments":
			root.Arguments = root.Arguments.New(val)
		case "$namespaces":
			root.Namespaces = root.Namespaces.New(val)
		case "stdin":
			root.Stdin = val.(string)
		case "stdout":
			root.Stdout = val.(string)
		case "inputs":
			root.Inputs = root.Inputs.New(val)
		case "outputs":
			root.Outputs = root.Outputs.New(val)
		case "requirements":
			root.Requirements = root.Requirements.New(val)
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
