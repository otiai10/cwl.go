package cwl

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	Version      string
	Class        string
	Hints        Hints
	Doc          string
	Graphs       Graphs
	BaseCommands BaseCommands
	Arguments    Arguments
	Namespaces   Namespaces
	Schemas      Schemas
	Stdin        string
	Stdout       string
	Stderr       string
	Inputs       Inputs `json:"inputs"`
	// ProvidedInputs ProvidedInputs `json:"-"`
	Outputs      Outputs
	Requirements Requirements
	Steps        Steps
	ID           string // ID only appears if this Root is a step in "steps"
	Expression   string // appears only if Class is "ExpressionTool"

	// Path
	Path string `json:"-"`
}

// UnmarshalMap decode map[string]interface{} to *Root.
func (root *Root) UnmarshalMap(docs map[string]interface{}) error {
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
		case "$schemas":
			root.Schemas = root.Schemas.New(val)
		case "$graph":
			root.Graphs = root.Graphs.New(val)
		case "stdin":
			root.Stdin = val.(string)
		case "stdout":
			root.Stdout = val.(string)
		case "stderr":
			root.Stderr = val.(string)
		case "inputs":
			root.Inputs = root.Inputs.New(val)
		case "outputs":
			root.Outputs = root.Outputs.New(val)
		case "requirements":
			root.Requirements = root.Requirements.New(val)
		case "steps":
			root.Steps = root.Steps.New(val)
		case "id":
			root.ID = val.(string)
		case "expression":
			root.Expression = val.(string)
		}
	}
	return nil
}

// UnmarshalJSON ...
func (root *Root) UnmarshalJSON(b []byte) error {
	docs := map[string]interface{}{}
	if err := json.Unmarshal(b, &docs); err != nil {
		return err
	}
	return root.UnmarshalMap(docs)
}

// Decode decodes specified file to this root
func (root *Root) Decode(r io.Reader) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("Parse error: %v", e)
		}
	}()
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

// AsStep constructs Root as a step of "steps" from interface.
func (root *Root) AsStep(i interface{}) *Root {
	dest := new(Root)
	switch x := i.(type) {
	case string:
		dest.ID = x
	case map[string]interface{}:
		err := dest.UnmarshalMap(x)
		if err != nil {
			panic(fmt.Sprintf("Failed to parse step as CWL.Root: %v", err))
		}
	}
	return dest
}
