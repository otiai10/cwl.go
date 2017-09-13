package cwl

// Graphs represents "$graph" field in CWL.
type Graphs []Graph

// Graph represents an element of "steps"
type Graph struct {
	Run *Root
}

// New constructs "Graphs" from interface.
func (steps Graphs) New(i interface{}) Graphs {
	dest := Graphs{}
	switch x := i.(type) {
	case []interface{}:
		// TODO;
		for _, v := range x {
			g := Graph{}
			g.Run = g.Run.AsStep(v)
			dest = append(dest, g)
		}
	}
	return dest
}

// New constructs "Step" from interface.
func (graph Graph) New(i interface{}) Graph {
	dest := Graph{}
	return dest
}
