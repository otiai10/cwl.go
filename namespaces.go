package cwl

// Namespaces ...
type Namespaces []Namespace

// New constructs "Namespaces" struct.
func (namespaces Namespaces) New(i interface{}) Namespaces {
	dest := []Namespace{}
	switch x := i.(type) {
	case map[string]interface{}:
		for key, val := range x {
			namespace := Namespace{}
			namespace[key] = val
			dest = append(dest, namespace)
		}
	}
	return dest
}

// Namespace ...
type Namespace map[string]interface{}
