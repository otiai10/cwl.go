package cwl

import (
	"fmt"
	"sort"
)

// Root ...
type Root struct {
	Version        string                   `yaml:"cwlVersion"`
	Path           string                   `yaml:"-"`
	Class          string                   `yaml:"class"`
	BaseCommand    string                   `yaml:"baseCommand"`
	RequiredInputs map[string]RequiredInput `yaml:"inputs"`
	ProvidedInputs map[string]ProvidedInput `yaml:"-"`
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
