package cwl

import (
	"fmt"
	"io"
	"io/ioutil"
	"sort"

	yaml "gopkg.in/yaml.v2"
)

// NewCWL ...
func NewCWL() *Root {
	root := new(Root)
	root.RequiredInputs = map[string]RequiredInput{}
	root.ProvidedInputs = ProvidedInputs{}
	return root
}

// Root ...
type Root struct {
	Version        string                   `yaml:"cwlVersion"`
	Path           string                   `yaml:"-"`
	Class          string                   `yaml:"class"`
	BaseCommand    string                   `yaml:"baseCommand"`
	RequiredInputs map[string]RequiredInput `yaml:"inputs"`
	ProvidedInputs ProvidedInputs           `yaml:"-"`
}

// RequiredInput ...
type RequiredInput struct {
	Name    string       `yaml:"-"`
	Type    string       `yaml:"string"`
	Binding InputBinding `yaml:"inputBinding"`
}

// RequiredInputsSortable ...
type RequiredInputsSortable []RequiredInput

// Len ...
func (s RequiredInputsSortable) Len() int { return len(s) }

// Swap ...
func (s RequiredInputsSortable) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// Less ...
func (s RequiredInputsSortable) Less(i, j int) bool {
	return s[i].Binding.Position < s[j].Binding.Position
}

// InputBinding ...
type InputBinding struct {
	Position int `yaml:"position"`
}

// Args TODO: should be Workflow
func (root *Root) Args() ([]string, error) {
	inputs := RequiredInputsSortable{}
	for key, val := range root.RequiredInputs {
		if _, ok := root.ProvidedInputs[key]; !ok {
			return nil, fmt.Errorf("Input `%s` is required but not provided", key)
		}
		val.Name = key
		inputs = append(inputs, val)
	}
	sort.Sort(inputs)

	args := []string{}
	for _, required := range inputs {
		args = append(args, root.ProvidedInputs[required.Name].Arg())
	}
	return args, nil
}

// Decode decodes specified io.Reader to this root
func (root *Root) Decode(r io.Reader) error {
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
