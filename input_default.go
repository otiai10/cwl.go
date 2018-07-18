package cwl

import (
	"fmt"
	"io/ioutil"
	"reflect"
)

// InputDefault represents "default" field in an element of "inputs".
type InputDefault struct {
	ID string

	Self  interface{}
	Kind  reflect.Kind
	Entry *Entry
	Error error

	// TODO: Refactor
	Int int
}

// New constructs new "InputDefault".
func (d InputDefault) New(i interface{}) *InputDefault {
	dest := &InputDefault{Self: i, Kind: reflect.TypeOf(i).Kind()}
	switch v := i.(type) {
	case nil:
		return dest // do nothing
	case int:
		dest.Int = v
	case map[string]interface{}: // It's "File" in most cases
		dest.Entry, dest.Error = dest.EntryFromDictionary(v)
	}
	return dest
}

// Flatten ...
func (d *InputDefault) Flatten(binding *Binding) []string {
	flattened := []string{}
	if d.Entry != nil {
		flattened = append(flattened, d.Entry.Location)
	}
	if binding != nil && binding.Prefix != "" {
		flattened = append([]string{binding.Prefix}, flattened...)
	}
	return flattened
}

// EntryFromDictionary ...
func (d *InputDefault) EntryFromDictionary(dict map[string]interface{}) (*Entry, error) {
	if dict == nil {
		return nil, nil
	}
	if dict["class"] == nil {
		return nil, nil
	}
	class := dict["class"].(string)
	location := dict["location"]
	contents := dict["contents"]
	if class == "" && location == nil && contents == nil {
		return nil, nil
	}
	switch class {
	case "File":
		// Use location if specified
		if location != nil {
			return &Entry{
				Class:    class,
				Location: fmt.Sprintf("%v", location),
				File:     File{},
			}, nil
		}
		// Use contents if specified
		if contentsstring, ok := contents.(string); ok {
			tmpfile, err := ioutil.TempFile("/tmp", d.ID)
			if err != nil {
				return nil, err
			}
			defer tmpfile.Close()
			if _, err := tmpfile.WriteString(contentsstring); err != nil {
				return nil, err
			}
			return &Entry{
				Class:    class,
				Location: tmpfile.Name(),
				File:     File{},
			}, nil
		}
	}
	return nil, nil
}
