package cwl

// Graphs represents "$graph" field in CWL.
type Graphs []*Root

// Graph represents an element of "steps"
type Graph struct {
	Run *Root
}

// New constructs "Graphs" from interface.
func (_ Graphs) New(i interface{}) Graphs {
	dest := Graphs{}
	switch x := i.(type) {
	case []interface{}:
		for _, v := range x {
			g := new(Root)
			g = g.AsStep(v)
			dest = append(dest, g)
		}
	}
	return dest
}
