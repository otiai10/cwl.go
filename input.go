package cwl

import (
	"fmt"
	"strings"
)

// Input represents "CommandInputParameter".
// @see http://www.commonwl.org/v1.0/CommandLineTool.html#CommandInputParameter
type Input struct {
	ID             string          `json:"id"`
	Label          string          `json:"label"`
	Doc            string          `json:"doc"`
	Format         string          `json:"format"`
	Binding        *Binding        `json:"inputBinding"`
	Default        *InputDefault   `json:"default"`
	Types          []Type          `json:"type"`
	SecondaryFiles []SecondaryFile `json:"secondary_files"`
	// Input.Provided is what provided by parameters.(json|yaml)
	Provided interface{} `json:"-"`
}

// New constructs "Input" struct from interface{}.
func (_ Input) New(i interface{}) Input {
	dest := Input{}
	switch x := i.(type) {
	case map[string]interface{}:
		for key, v := range x {
			switch key {
			case "id":
				dest.ID = v.(string)
			case "type":
				dest.Types = Type{}.NewList(v)
			case "label":
				dest.Label = v.(string)
			case "doc":
				dest.Doc = v.(string)
			case "inputBinding":
				dest.Binding = Binding{}.New(v)
			case "default":
				dest.Default = InputDefault{}.New(v)
			case "format":
				dest.Format = v.(string)
			case "secondaryFiles":
				dest.SecondaryFiles = SecondaryFile{}.NewList(v)
			}
		}
	case string:
		dest.Types = Type{}.NewList(x)
	case []interface{}:
		for _, v := range x {
			dest.Types = append(dest.Types, Type{}.New(v))
		}
	}
	return dest
}

// flatten
func (input Input) flatten(typ Type, binding *Binding) []string {
	flattened := []string{}
	switch typ.Type {
	case "int": // Array of Int
		tobejoined := []string{}
		for _, e := range input.Provided.([]interface{}) {
			tobejoined = append(tobejoined, fmt.Sprintf("%v", e))
		}
		flattened = append(flattened, strings.Join(tobejoined, input.Binding.Separator))
	case "File": // Array of Files
		switch arr := input.Provided.(type) {
		case []string:
			// TODO:
		case []interface{}:
			separated := []string{}
			for _, e := range arr {
				switch v := e.(type) {
				case map[interface{}]interface{}:
					if binding != nil && binding.Prefix != "" {
						separated = append(separated, binding.Prefix)
					}
					separated = append(separated, fmt.Sprintf("%v", v["location"]))
				default:
					// TODO:
				}
			}
			// In case it's Array of Files, unlike array of int,
			// it's NOT gonna be joined with .Binding.Separator.
			flattened = append(flattened, separated...)
		}
	default:
		// TODO:
	}
	return flattened
}

// Flatten ...
func (input Input) Flatten() []string {
	if input.Provided == nil {
		// In case "input.Default == nil" should be validated by usage layer.
		return input.Default.Flatten(input.Binding)
	}
	flattened := []string{}
	if repr := input.Types[0]; len(input.Types) == 1 {
		switch repr.Type {
		case "array":
			flattened = append(flattened, input.flatten(repr.Items[0], repr.Binding)...)
		case "int":
			flattened = append(flattened, fmt.Sprintf("%v", input.Provided.(int)))
		case "File":
			switch provided := input.Provided.(type) {
			case map[interface{}]interface{}:
				// TODO: more strict type casting
				flattened = append(flattened, fmt.Sprintf("%v", provided["location"]))
			default:
			}
		default:
			flattened = append(flattened, fmt.Sprintf("%v", input.Provided))
		}
	}
	if input.Binding != nil && input.Binding.Prefix != "" {
		flattened = append([]string{input.Binding.Prefix}, flattened...)
	}

	return flattened
}

// Inputs represents "inputs" field in CWL.
type Inputs []Input

// New constructs new "Inputs" struct.
func (_ Inputs) New(i interface{}) Inputs {
	dest := Inputs{}
	switch x := i.(type) {
	case []interface{}:
		for _, v := range x {
			dest = append(dest, Input{}.New(v))
		}
	case map[string]interface{}:
		for key, v := range x {
			input := Input{}.New(v)
			input.ID = key
			dest = append(dest, input)
		}
	}
	return dest
}

// Len for sorting.
func (ins Inputs) Len() int {
	return len(ins)
}

// Less for sorting.
func (ins Inputs) Less(i, j int) bool {
	prev, next := ins[i].Binding, ins[j].Binding
	switch [2]bool{prev == nil, next == nil} {
	case [2]bool{true, true}:
		return true
	case [2]bool{false, true}:
		return prev.Position < 0
	case [2]bool{true, false}:
		return next.Position > 0
	default:
		return prev.Position <= next.Position
	}
}

// Swap for sorting.
func (ins Inputs) Swap(i, j int) {
	ins[i], ins[j] = ins[j], ins[i]
}
