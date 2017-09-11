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
	// root.RequiredInputs = RequiredInputs{}
	// root.ProvidedInputs = ProvidedInputs{}
	return root
}

// Root ...
type Root struct {
	Version string `json:"cwlVersion"`
	// Class          string         `json:"class"`
	// BaseCommand    string         `json:"baseCommand"`
	// RequiredInputs RequiredInputs `json:"inputs"`
	// ProvidedInputs ProvidedInputs `json:"-"`
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
