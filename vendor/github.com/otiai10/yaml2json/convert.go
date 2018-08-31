package yaml2json

import (
	"encoding/json"
	"io"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// Y2J converts yaml to json.
func Y2J(r io.Reader) ([]byte, error) {
	result := []byte{}
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return result, err
	}
	var root interface{}
	if err := yaml.Unmarshal(b, &root); err != nil {
		return result, err
	}
	return json.Marshal(convert(root))
}

// J2Y converts json to yaml.
func J2Y(r io.Reader) ([]byte, error) {
	result := []byte{}
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return result, err
	}
	var root interface{}
	if err := json.Unmarshal(b, &root); err != nil {
		return result, err
	}
	return yaml.Marshal(convert(root))
}

// convert ...
func convert(parent interface{}) interface{} {
	switch entity := parent.(type) {
	case map[interface{}]interface{}:
		node := map[string]interface{}{}
		for key, val := range entity {
			node[key.(string)] = convert(val)
		}
		return node
	case []interface{}:
		for idx, val := range entity {
			entity[idx] = convert(val)
		}
		return entity
	}
	return parent
}
