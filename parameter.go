package cwl

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

// Parameter ...
type Parameter interface{}

// Parameters represents specific parameters to run workflow which is described by CWL.
type Parameters map[string]Parameter

// NewParameters ...
func NewParameters() *Parameters {
	return &Parameters{}
}

// Decode ...
func (p *Parameters) Decode(f *os.File) error {
	switch filepath.Ext(f.Name()) {
	case "json":
		return json.NewDecoder(f).Decode(p)
	default:
		b, err := ioutil.ReadAll(f)
		if err != nil {
			return err
		}
		return yaml.Unmarshal(b, p)
	}
}
